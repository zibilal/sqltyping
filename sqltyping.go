package sqltyping

import (
	"fmt"
	"reflect"
	"bytes"
	"strings"
	"text/scanner"
	"strconv"
)

type SqlTyping struct {
	typeQuery     string
}

func NewSqlTyping(typeQuery string) *SqlTyping {
	s := new(SqlTyping)
	s.typeQuery = typeQuery
	return s
}

func (t *SqlTyping) Typing(input interface{}) ([]string, error) {
	ival := reflect.ValueOf(input)

	switch ival.Kind() {
	case reflect.Struct:
		buff := bytes.NewBufferString("")
		err := TypeIterator(input, buff)

		if err != nil {
			return []string{}, err
		}

		return []string{}, nil
	default:
		return []string{}, fmt.Errorf("not supported type %T", input)
	}

}

func processBytes(buff *bytes.Buffer) ([]string, error) {

	var s scanner.Scanner
	s.Init(buff)

	openOffsets := []int{}
	closeOffsets := []int{}
	components := []string{}
	pairs := []string{}

	bytesData := buff.Bytes()

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		if s.TokenText() == "{" {
			openOffsets = append(openOffsets, s.Position.Column)
		}
		if s.TokenText() == "}" {
			closeOffsets = append(closeOffsets, s.Position.Column)
		}
	}


	for i := len(openOffsets) - 1; i >= 0; i-- {
		pairs = append(pairs, fmt.Sprintf("%d,%d", openOffsets[i], closeOffsets[len(closeOffsets)-1-i]))
	}

	fmt.Println("Pairs", pairs)

	for idx := 0; idx < len(pairs); idx++ {
		pair := pairs[idx]
		bytesInside := []byte{}
		split := strings.Split(pair, ",")
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])
		bytesInside = append(bytesInside, bytesData[num1:num2-2]...)
		bytesData = append(bytesData[:num1], bytesData[num2:]...)
		if idx + 1 < len(pairs) {
			upSplit := strings.Split(pairs[idx+1], ",")
			upNum1, _ := strconv.Atoi(upSplit[0])
			upNum2, _ := strconv.Atoi(upSplit[1])
			upNum2 = upNum2 - 2

			if upNum1 - len(bytesInside) >= 1 {
				upNum1 = upNum1 - len(bytesInside)
			}

			if upNum2 - len(bytesInside) >= 1 {
				upNum2 = upNum2 - len(bytesInside)
			}

			pairs[idx+1] = fmt.Sprintf("%d,%d", upNum1, upNum2)
		}

		components = append(components, string(bytesInside))
	}

	return components, nil
}

func (t *SqlTyping) processInput(input string, name string) string {

	switch t.typeQuery {
	case "SELECT":
		pairs := strings.Split(input, " ")
		columns := []string {}
		wheres := []string {}
		for _, pair := range pairs {
			if pair != "" {
				values := strings.Split(pair, "|")
				columns = append(columns, values[0])
				if len(values) == 2 && values[1] != "" {
					wheres = append(wheres, fmt.Sprintf("%s = '%s'", values[0], values[1]) )
				}
			}
		}

		baseQuery := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns,", "), name)
		if len(wheres) > 0 {
			baseQuery = fmt.Sprintf("%s WHERE %s", baseQuery, strings.Join(wheres, " AND "))
		}

		return baseQuery
	case "INSERT":
	case "UPDATE":
	}

	return ""
}
