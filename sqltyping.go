package sqltyping

import (
	"fmt"
	"reflect"
	"bytes"
	"strings"
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

		result := []string{}
		splits := strings.Split(buff.String(), "\n")
		fmt.Println("SPLITS", splits)
		for _, s := range splits {
			if s != "" {
				fmt.Println("S", s)
				dsplit := strings.Split(s, "||")
				fmt.Println("dsplit", dsplit)
				name := strings.ToLower(dsplit[0] + "s")
				result = append(result, t.processInput(dsplit[1], name) )
			}
		}

		return result, nil
	default:
		return []string{}, fmt.Errorf("not supported type %T", input)
	}

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
