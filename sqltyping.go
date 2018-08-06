package sqltyping

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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
	typeQuery string
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
		if err != nil {
			return []string{}, err
		}

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

func (t *SqlTyping) TypingUpdateWithWhereClause(input interface{}, query interface{}) (string, error) {
	ival := reflect.ValueOf(input)
	ique := reflect.ValueOf(query)

	if ival.Kind() != reflect.Struct || ique.Kind() != reflect.Struct {
		return "", errors.New("please provide both input and query as variable of type struct")
	}

	var err error
	buff1 := bytes.NewBufferString("")
	err = TypeIterator(input, buff1)
	if err != nil {
		return "", err
	}
	buff2 := bytes.NewBufferString("")
	err = TypeIterator(query, buff2)
	if err != nil {
		return "", err
	}

	data1, err := singleProcessBytes(buff1.Bytes())
	if err != nil {
		return "", err
	}

	data2, err := singleProcessBytes(buff2.Bytes())
	if err != nil {
		return "", err
	}

	qstr := t.processUpdateWithWhere(data1, data2)

	return qstr, nil
}

func singleProcessBytes(bytesData []byte) (string, error) {
	if len(bytesData) == 0 {
		return "", errors.New("[singleProcessBytes]Empty bytes data")
	} else {
		buff := bytes.NewBuffer(bytesData)
		dt := buff.String()
		var tmp string
		var pairs []string
		offStack := new(offsetstack)
		var h, l int
		for idx := 1; idx < len(dt); idx += 2 {

			if len(dt)%2 != 0 && idx+2 > len(dt)-1 {
				h = idx + 2
				l = idx + 2
				tmp = dt[idx : idx+2]
			} else {
				h = idx + 1
				l = idx + 1
				tmp = dt[idx-1 : idx+1]
			}
			if tmp == "((" {
				offStack.push(h)
			}
			if tmp == "))" {
				pairs = append(pairs,
					fmt.Sprintf("%d,%d", offStack.pop(), l-2))
			}
		}

		bytesInside := []byte{}
		for idx := 0; idx < len(pairs); idx++ {
			pair := pairs[idx]
			split := strings.Split(pair, ",")
			num1, _ := strconv.Atoi(split[0])
			num2, _ := strconv.Atoi(split[1])

			bytesInside = append(bytesInside, bytesData[num1:num2]...)
		}
		return string(bytesInside), nil
	}
}

func processBytes(bytesData []byte, components []string) ([]string, error) {
	if len(bytesData) == 0 {
		return components, nil
	} else {
		buff := bytes.NewBuffer(bytesData)
		dt := buff.String()

		var (
			itm        uint8
			countOpen  uint8
			countClose uint8
			pairs      []string
		)
		offStack := new(offsetstack)
		for idx := 0; idx < len(dt); idx++ {
			itm = dt[idx]

			if itm == '(' {
				countOpen++
				if countOpen == 2 {
					offStack.push(idx + 1)
					countOpen = 0
				}
			}

			if itm == ')' {
				countClose++
				if countClose == 2 {
					pairs = append(pairs, fmt.Sprintf("%d,%d", offStack.pop(), idx-1))
					countClose = 0
				}
			}
		}

		for idx := 0; idx < len(pairs); idx++ {
			pair := pairs[idx]
			bytesInside := []byte{}
			split := strings.Split(pair, ",")
			num1, _ := strconv.Atoi(split[0])
			num2, _ := strconv.Atoi(split[1])

			if len(bytesData)-num2 == 3 {
				num2 += 1
			}

			bytesInside = append(bytesInside, bytesData[num1:num2]...)
			bytesData = append(bytesData[:num1-2], bytesData[num2+2:]...)

			if len(bytesData) < 2 {
				bytesData = []byte{}
			}

			components = append(components, string(bytesInside))
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
	splitComma := strings.Split(input, ";")
	fromClause := ""
	columns := []string{}
	wheres := []string{}

	for _, byComma := range splitComma {
		pair := strings.Split(byComma, "=")
		switch pair[0] {
		case "table_name":
			fromClause = convertCamelCaseToSnakeCase(pair[1])
		case "column_name":
			splitValue := strings.Split(pair[1], "|")
			if len(splitValue) == 2 {
				columns = append(columns, splitValue[0])
				if splitValue[1] != "" && splitValue[1] != "0" {
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

	splitComma := strings.Split(input, ";")
	for _, byComma := range splitComma {
		pair := strings.Split(byComma, "=")
		switch pair[0] {
		case "table_name":
			tableName = convertCamelCaseToSnakeCase(pair[1])
		case "column_name":
			splitValue := strings.Split(pair[1], "|")
			if len(splitValue) == 2 && strings.TrimSpace(splitValue[1]) != "" {
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

	splitComma := strings.Split(input, ";")
	for _, byComma := range splitComma {
		pair := strings.Split(byComma, "=")
		switch pair[0] {
		case "table_name":
			tableName = convertCamelCaseToSnakeCase(pair[1])
		case "column_name":
			splitValue := strings.Split(pair[1], "|")
			if len(splitValue) == 2 && strings.TrimSpace(splitValue[1]) != "" && strings.TrimSpace(splitValue[1]) != "0" {
				if strings.ToLower(splitValue[0]) == "id" || splitValue[0] == t.updateKey {
					if where == "" {
						where = fmt.Sprintf(" %s='%s'", splitValue[0], splitValue[1])
					} else {
						where += fmt.Sprintf(" AND %s='%s'", splitValue[0], splitValue[1])
					}
				} else {
					setColumns = append(setColumns, fmt.Sprintf("%s='%s'", splitValue[0], splitValue[1]))
				}
			}
		}
	}

	if where != "" {
		where = " WHERE" + where
	}

	return fmt.Sprintf("UPDATE %s SET %s%s", tableName, strings.Join(setColumns, ","), where)
}

func (t *SqlTyping) processUpdateWithWhere(input string, where string) string {
	setColumns := []string{}
	tableName := ""

	splitCommaSet := strings.Split(input, ";")
	for _, byComma := range splitCommaSet {
		pair := strings.Split(byComma, "=")
		switch pair[0] {
		case "table_name":
			tableName = convertCamelCaseToSnakeCase(pair[1])
		case "column_name":
			splitValue := strings.Split(pair[1], "|")
			if len(splitValue) == 2 {
				setColumns = append(setColumns, fmt.Sprintf("%s='%s'", splitValue[0], splitValue[1]))
			}
		}
	}

	splitWhereCommaSet := strings.Split(where, ";")
	theWhere := []string{}
	for _, byComma := range splitWhereCommaSet {
		pair := strings.Split(byComma, "=")
		switch pair[0] {
		case "column_name":
			splitValue := strings.Split(pair[1], "|")
			if len(splitValue) == 2 && strings.TrimSpace(splitValue[1]) != "" {
				theWhere = append(theWhere, fmt.Sprintf("%s='%s'", splitValue[0], splitValue[1]))
			}
		}
	}

	whereClause := ""
	if len(theWhere) > 0 {
		whereClause = " WHERE " + strings.Join(theWhere, " AND ")
	}

	return fmt.Sprintf("UPDATE %s SET %s%s", tableName, strings.Join(setColumns, ","), whereClause)
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
