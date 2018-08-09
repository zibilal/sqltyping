package sqltyping

import (
	"bytes"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestTypingSelect(t *testing.T) {
	eventLaunching := struct {
		ID           uint64 `json:"id"`
		OrderSplitID string `json:"order_split_id"`
		UseVoucher   int    `json:"use_voucher"`
		Status       int    `json:"status"`
	}{
		ID: 151175,
	}

	sqlTyping := NewSqlTyping("SELECT")
	results, err := sqlTyping.Typing(eventLaunching)

	if err != nil {
		t.Fatalf("%s expected error nil, got %s", failed, err.Error())
	} else {
		t.Logf("%s expected error nil", success)
	}

	if len(results) != 1 {
		t.Fatalf("%s expected result have length 1, got %d", failed, len(results))
	} else {
		t.Logf("%s expected result have length 1", success)
	}

	expectedQuery := "SELECT id,order_split_id,use_voucher,status FROM  WHERE id='151175'"
	if expectedQuery == results[0] {
		t.Logf("%s expected query == %s", success, expectedQuery)
	} else {
		t.Fatalf("%s expected query == %s , got %s", failed, expectedQuery, results[0])
	}
}

func TestSqlTyping(t *testing.T) {
	t.Log("Testing sql typing")
	{
		expectedSql := `SELECT id,username,first_name,last_name,email FROM user WHERE id='bhf1234584'`

		user := User{
			Id: "bhf1234584",
		}

		sqlTyping := NewSqlTyping("SELECT")
		results, err := sqlTyping.Typing(user)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if len(results) > 0 {
			t.Logf("%s expected results length bigger than zero = %d", success, len(results))
		} else {
			t.Fatalf("%s expected results length bigger than zero = %d", failed, len(results))
		}

		if len(results) > 1 && expectedSql == results[1] {
			t.Logf("%s expected query = %s", success, expectedSql)
		} else {
			t.Errorf("%s expected query = %s, got %s", failed, expectedSql, results[0])
		}
	}
}

func TestSqlTypingProcessSelect(t *testing.T) {
	t.Log("Testing SqlTyping.processSelect")
	{
		dataInput := "table_name=User;column_name=id|bhf1234584;column_name=username|;column_name=first_name|;column_name=last_name|;column_name=email|;column_name=secret_detail"
		expectedQuery := `SELECT id,username,first_name,last_name,email FROM user WHERE id='bhf1234584'`
		typing := NewSqlTyping("SELECT")
		result := typing.processSelect(dataInput)

		if result == "" {
			t.Fatalf("%s expected result is not empty", failed)
		} else {

			if result == expectedQuery {
				t.Logf("%s expected result = %s", success, expectedQuery)
			} else {
				t.Fatalf("%s expected result = %s, got %s", failed, expectedQuery, result)
			}
		}
	}
}

func TestSqlTypingProcessInsert(t *testing.T) {
	t.Log("Testing SqlTyping.processInsert")
	{
		dataInput := "table_name=User;column_name=id|bhf1234584;column_name=username|example;column_name=first_name|first;column_name=last_name|last;column_name=email|first.last@example.com;column_name=secret_detail"
		expectedQuery := `INSERT INTO user (id,username,first_name,last_name,email) VALUES ('bhf1234584','example','first','last','first.last@example.com')`
		typing := NewSqlTyping("INSERT")
		result := typing.processInsert(dataInput)

		if result == "" {
			t.Fatalf("%s expected result is not empty", failed)
		} else {

			if result == expectedQuery {
				t.Logf("%s expected result = %s", success, expectedQuery)
			} else {
				t.Fatalf("%s expected result = %s, got %s", failed, expectedQuery, result)
			}
		}
	}
	t.Log("Testing SqlTyping.processInsert with empty id")
	{
		dataInput := "table_name=User;column_name=id|;column_name=username|example;column_name=first_name|first;column_name=last_name|last;column_name=email|first.last@example.com;column_name=secret_detail"
		expectedQuery := `INSERT INTO user (username,first_name,last_name,email) VALUES ('example','first','last','first.last@example.com')`
		typing := NewSqlTyping("INSERT")
		result := typing.processInsert(dataInput)

		if result == "" {
			t.Fatalf("%s expected result is not empty", failed)
		} else {

			if result == expectedQuery {
				t.Logf("%s expected result = %s", success, expectedQuery)
			} else {
				t.Fatalf("%s expected result = %s, got %s", failed, expectedQuery, result)
			}
		}
	}
}

func TestSqlTypingProcessUpdate(t *testing.T) {
	t.Log("Testing SqlTyping.processUpdate")
	{
		dataInput := "table_name=User;column_name=id|bhf1234584;column_name=username|example;column_name=first_name|first;column_name=last_name|last;column_name=email|first.last@example.com;column_name=secret_detail"
		expectedQuery := `UPDATE user SET first_name='first',last_name='last',email='first.last@example.com' WHERE id='bhf1234584' AND username='example'`
		typing := NewSqlTyping("UPDATE")
		typing.SetUpdateKey("username")
		result := typing.processUpdate(dataInput)

		if result == "" {
			t.Fatalf("%s expected result is not empty", failed)
		} else {

			if result == expectedQuery {
				t.Logf("%s expected result = %s", success, expectedQuery)
			} else {
				t.Fatalf("%s expected result = %s, got %s", failed, expectedQuery, result)
			}
		}
	}
	t.Log("Testing SqlTyping.processUpdate, with some empty data")
	{
		dataInput := "table_name=User;column_name=id|;column_name=username|example;column_name=first_name|;column_name=last_name|last;column_name=email|first.last@example.com;column_name=secret_detail"
		expectedQuery := `UPDATE user SET username='example',last_name='last',email='first.last@example.com'`
		typing := NewSqlTyping("UPDATE")
		result := typing.processUpdate(dataInput)

		if result == "" {
			t.Fatalf("%s expected result is not empty", failed)
		} else {

			if result == expectedQuery {
				t.Logf("%s expected result = %s", success, expectedQuery)
			} else {
				t.Fatalf("%s expected result = %s, got %s", failed, expectedQuery, result)
			}
		}
	}

	t.Log("Testing SqlTyping.processUpdate, with decided where column name")
	{
		dataInput := "table_name=User;column_name=id|;column_name=username|example;column_name=first_name|;column_name=last_name|last;column_name=email|first.last@example.com;column_name=secret_detail"
		expectedQuery := `UPDATE user SET last_name='last',email='first.last@example.com' WHERE username='example'`
		typing := NewSqlTyping("UPDATE")
		typing.SetUpdateKey("username")
		result := typing.processUpdate(dataInput)

		if result == "" {
			t.Fatalf("%s expected result is not empty", failed)
		} else {

			if result == expectedQuery {
				t.Logf("%s expected result = %s", success, expectedQuery)
			} else {
				t.Fatalf("%s expected result = %s, got %s", failed, expectedQuery, result)
			}
		}
	}
}

func TestTescan(t *testing.T) {

	t.Log("Testing text scan")
	{
		var data = `((table_name=User;column_name=id|bhf1234584;column_name=username|;column_name=first_name|;column_name=last_name|;column_name=email|;column_name=secret_detail((table_name=SecretDetailEx;column_name=id|11244;column_name=api_secret|;column_name=api_token|))))`

		bytesData := []byte(data)
		buff := bytes.NewBuffer(bytesData)
		t.Log("LEN: Buff", len(buff.Bytes()))
		components, err := processBytes(bytesData, []string{})

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if len(components) != 2 {
			t.Fatalf("%s expected length of components == 2, got %d", failed, len(components))
		} else {
			t.Logf("%s expected length of components == 2", success)
		}

		expectedFirst := "table_name=SecretDetailEx;column_name=id|11244;column_name=api_secret|;column_name=api_token|"
		expectedSecond := "table_name=User;column_name=id|bhf1234584;column_name=username|;column_name=first_name|;column_name=last_name|;column_name=email|;column_name=secret_detail"

		if expectedFirst == components[0] {
			t.Logf("%s expected first component == %s", success, expectedFirst)
		} else {
			t.Fatalf("%s expected first component == %s, got %s", failed, expectedFirst, components[0])
		}

		if expectedSecond == components[1] {
			t.Logf("%s expected second component == %s", success, expectedSecond)
		} else {
			t.Fatalf("%s expected second component == %s, got %s", failed, expectedSecond, components[1])
		}
	}

}

func TestScanWithStructOfSlice(t *testing.T) {
	t.Log("Testing text scan")
	{
		var data = `((table_name=OrderEx;column_name=Id|o123;column_name=Updated|2018-06-17 03:43:33;column_name=Created|2018-06-17 03:43:33;column_name=Status|OrderCreated;column_name=Items((table_name=OrderItem;column_name=Id|itm123;column_name=ItemName|XL 2 Giga;column_name=Price|150000))((table_name=OrderItem;column_name=Id|itm124;column_name=ItemName|XL 5 Giga;column_name=Price|300000))))`
		bytesData := []byte(data)
		components, err := processBytes(bytesData, []string{})

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if len(components) != 3 {
			t.Errorf("%s expected length of components == 2, got %d", failed, len(components))
		} else {
			t.Logf("%s expected length of components == 2", success)

			for _, component := range components {
				t.Logf("%s %s", success, component)
			}

		}
	}
}

func TestConvertCamelCaseToSnakeCase(t *testing.T) {

	t.Log("1 conversion")
	{
		input1 := "TestConvertCamelCaseToSnakeCase"
		expected1 := "test_convert_camel_case_to_snake_case"

		str := convertCamelCaseToSnakeCase(input1)

		if str == expected1 {
			t.Logf("%s expected result = expected1 : %s", success, expected1)
		} else {
			t.Errorf("%s expected result = expected1 : got %s", failed, str)
		}
	}

	t.Log("2 conversion")
	{
		input1 := "testConvertCamelCaseToSnakeCase"
		expected1 := "test_convert_camel_case_to_snake_case"

		str := convertCamelCaseToSnakeCase(input1)

		if str == expected1 {
			t.Logf("%s expected result = expected1 : %s", success, expected1)
		} else {
			t.Errorf("%s expected result = expected1 : got %s", failed, str)
		}
	}

	t.Log("3 conversion")
	{
		input1 := "test&*4ConvertCamelCaseToSnakeCasess34ssd"
		expected1 := "test_convert_camel_case_to_snake_casess34ssd"

		str := convertCamelCaseToSnakeCase(input1)

		if str == expected1 {
			t.Logf("%s expected result = expected1 : %s", success, expected1)
		} else {
			t.Errorf("%s expected result = expected1 : got %s", failed, str)
		}
	}
}

func TestSimpleInsertQuery(t *testing.T) {
	t.Log("Test simple insert query")
	{
		data := struct {
			Username   string
			Session    string
			ExpiryDate time.Time
		}{
			Username: "example1",
			Session:  "123123412431243",
		}

		typing := NewSqlTyping(InsertQuery)
		queries, _ := typing.Typing(data)

		t.Log("Generated query", queries)

	}
}

type TheStruct struct {
	Username string
	Session  string
	Echo     string
}

func TestSimpleUpdateQuery(t *testing.T) {
	t.Log("Test simple update query")
	{
		data := TheStruct{
			Username: "example1",
			Session:  "456789898989",
			Echo:     "The echo2",
		}

		data2 := TheStruct{
			Echo: "the echo",
		}

		typing := NewSqlTyping(UpdateQuery)
		query, err := typing.TypingUpdateWithWhereClause(data, data2)

		if err != nil {
			t.Fatalf("%s Expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s Expected error nil", success)
		}

		expectedQuery := `UPDATE the_struct SET Username='example1',Session='456789898989',Echo='The echo2' WHERE Echo='the echo'`

		if expectedQuery == query {
			t.Logf("%s expected query == %s", success, expectedQuery)
		} else {
			t.Fatalf("%s expected query == %s, got %s", failed, expectedQuery, query)
		}
	}
}

func TestNullSql(t *testing.T) {
	t.Log("Testing select null sql")
	{
		data := struct {
			Name    sql.NullString  `db:"name"`
			EmailAD sql.NullString  `db:"email_ad"`
			Salary  sql.NullFloat64 `db:"salary"`
			DateNow mysql.NullTime  `db:"date_now"`
		}{
			Name: sql.NullString{
				String: "Testing",
			},
		}

		typing := NewSqlTyping(SelectQuery)
		query, err := typing.Typing(data)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if len(query) != 1 {
			t.Fatalf("%s expected generate one query, got %d", failed, len(query))
		} else {
			t.Logf("%s expected generate one query", success)
		}

		expectedQuery := "SELECT name,email_ad,salary,date_now FROM  WHERE name='Testing'"

		if query[0] == expectedQuery {
			t.Logf("%s expected query %s", success, expectedQuery)
		} else {
			t.Fatalf("%s expected query %s, got %s", failed, expectedQuery, query[0])
		}
	}

	t.Log("Testing update null sql")
	{
		data := struct {
			Name    sql.NullString  `db:"name"`
			EmailAD sql.NullString  `db:"email_ad"`
			Salary  sql.NullFloat64 `db:"salary"`
			DateNow mysql.NullTime  `db:"date_now"`
		}{}

		typing := NewSqlTyping(UpdateQuery)
		query, err := typing.TypingUpdateWithWhereClause(data, struct {
			Name string
		}{
			"Test Name",
		})

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		expectedQuery := `UPDATE  SET name='',email_ad='',salary='0',date_now='' WHERE Name='Test Name'`
		if expectedQuery == query {
			t.Logf("%s expected query %s", success, expectedQuery)
		} else {
			t.Logf("%s expected query %s, got %s", success, expectedQuery, query)
		}
	}
}
