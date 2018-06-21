package sqltyping

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
)

type offsetstack struct {
	offsets []int
	count   int
}

func (s *offsetstack) push(v int) {
	s.offsets = append(s.offsets[:s.count], v)
	s.count++
}

func (s *offsetstack) pop() int {
	if s.count == 0 {
		return 0
	}
	s.count--
	return s.offsets[s.count]
}

const (
	SelectQuery = "SELECT"
	InsertQuery = "INSERT"
	UpdateQuery = "UPDATE"
)

type SqlTyping struct {
	typeQuery    string
	updateKey string
}

func NewSqlTyping(typeQuery string) *SqlTyping {
	s := new(SqlTyping)
	s.typeQuery = typeQuery
	return s
}

func (t *SqlTyping) SetUpdateKey(updateKey string) {
	t.updateKey = updateKey
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
		components, err := processBytes(buff.Bytes(), []string{})

		results := []string{}

		for _, component := range components {
			result := t.processInput(component)
			results = append(results, result)
		}

		return results, nil
	default:
		return []string{}, fmt.Errorf("not supported type %T", input)
	}

}

func processBytes(bytesData []byte, components []string) ([]string, error) {
	if len(bytesData) == 0 {
		return components, nil
	} else {
		buff := bytes.NewBuffer(bytesData)
		var s scanner.Scanner
		s.Init(buff)
		pairs := []string{}

		offStack := new(offsetstack)

		var tok rune
		for ; tok != scanner.EOF; tok = s.Scan() {
			if s.TokenText() == "{" {
				offStack.push(s.Position.Column)
			}
			if s.TokenText() == "}" {
				pairs = append(pairs,
					fmt.Sprintf("%d,%d", offStack.pop(), s.Position.Column))
			}
		}

		for idx := 0; idx < len(pairs); idx++ {
			pair := pairs[idx]
			bytesInside := []byte{}
			split := strings.Split(pair, ",")
			num1, _ := strconv.Atoi(split[0])
			num2, _ := strconv.Atoi(split[1])

			bytesInside = append(bytesInside, bytesData[num1-1:num2]...)
			bytesData = append(bytesData[:num1-1], bytesData[num2:]...)

			components = append(components, string(bytesInside[1:len(bytesInside)-1]))
			return processBytes(bytesData, components)
		}
	}

	return []string{}, nil
}

func (t *SqlTyping) processInput(input string) string {

	switch t.typeQuery {
	case SelectQuery:
		return t.processSelect(input)
	case InsertQuery:
		return t.processInsert(input)
	case UpdateQuery:
		return t.processUpdate(input)
	}

	return ""
}

func (t *SqlTyping) processSelect(input string) string {
	splitComma := strings.Split(input, ",")
	fromClause := ""
	columns := []string{}
	wheres := []string{}

	for _, byComma := range splitComma {
		pair := strings.Split(byComma, ":")
		switch pair[0] {
		case "table_name":
			fromClause = convertCamelCaseToSnakeCase(pair[1])
		case "column_name":
			splitValue := strings.Split(pair[1], "|")
			if len(splitValue) == 2 {
				columns = append(columns, splitValue[0])
				if splitValue[1] != "" {
					wheres = append(wheres, fmt.Sprintf("%s='%s'", splitValue[0], splitValue[1]))
				}
			}
		}
	}

	base := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ","), fromClause)
	buff := bytes.NewBufferString(base)
	if len(wheres) > 0 {
		buff.WriteString(" WHERE " + strings.Join(wheres, " AND "))
	}

	return buff.String()
}

func (t *SqlTyping) processInsert(input string) string {
	intos := []string{}
	tableName := ""
	values := []string{}

	splitComma := strings.Split(input, ",")
	for _, byComma := range splitComma {
		pair := strings.Split(byComma, ":")
		switch pair[0] {
		case "table_name":
			tableName = convertCamelCaseToSnakeCase(pair[1])
		case "column_name":
			splitValue := strings.Split(pair[1], "|")
			if len(splitValue) == 2 && splitValue[1] != "" {
				intos = append(intos, splitValue[0])
				values = append(values, "'"+splitValue[1]+"'")
			}
		}
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(intos, ","), strings.Join(values, ","))
}

func (t *SqlTyping) processUpdate(input string) string {
	setColumns := []string{}
	where := ""

	tableName := ""

	splitComma := strings.Split(input, ",")
	for _, byComma := range splitComma {
		pair := strings.Split(byComma, ":")
		switch pair[0] {
		case "table_name":
			tableName = convertCamelCaseToSnakeCase(pair[1])
		case "column_name":
			splitValue := strings.Split(pair[1], "|")
			if len(splitValue) == 2 && splitValue[1] != "" {
				if strings.ToLower(splitValue[0]) == "id" || splitValue[0] == t.updateKey {
					if where == "" {
						where = fmt.Sprintf(" WHERE %s='%s'", splitValue[0], splitValue[1])
					} else {
						where += where + fmt.Sprintf(" AND %s='%s'", splitValue[0], splitValue[1])
					}
				} else {
					setColumns = append(setColumns, fmt.Sprintf("%s='%s'", splitValue[0], splitValue[1]))
				}
			}
		}
	}

	return fmt.Sprintf("UPDATE %s SET %s%s", tableName, strings.Join(setColumns, ","), where)
}

func convertCamelCaseToSnakeCase(input string) string {

	re := regexp.MustCompile("[A-Za-z0-9][a-z0-9]+")
	strs := re.FindAllString(input, -1)
	buff := bytes.NewBufferString("")
	for idx, str := range strs {
		if idx == 0 {
			buff.WriteString(strings.ToLower(str))
		} else {
			buff.WriteString("_" + strings.ToLower(str))
		}
	}

	return buff.String()
}
