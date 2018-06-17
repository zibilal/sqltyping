package sqltyping

import (
	"bytes"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestSqlTyping(t *testing.T) {
	t.Log("Testing sql typing")
	{
		expectedSql := `SELECT id,username,first_name,last_name,email FROM users WHERE id='bhf1234584'`

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

		if expectedSql == results[1] {
			t.Logf("%s expected query = %s", success, expectedSql)
		} else {
			t.Errorf("%s expected query = %s, got %s", failed, expectedSql, results[0])
		}
	}
}

func TestSqlTypingProcessSelect(t *testing.T) {
	t.Log("Testing SqlTyping.processSelect")
	{
		dataInput := "table_name:User,column_name:id|bhf1234584,column_name:username|,column_name:first_name|,column_name:last_name|,column_name:email|,column_name:secret_detail"
		expectedQuery := `SELECT id,username,first_name,last_name,email FROM users WHERE id='bhf1234584'`
		result := processSelect(dataInput)

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
		dataInput := "table_name:User,column_name:id|bhf1234584,column_name:username|example,column_name:first_name|first,column_name:last_name|last,column_name:email|first.last@example.com,column_name:secret_detail"
		expectedQuery := `INSERT INTO users (id,username,first_name,last_name,email) VALUES ('bhf1234584','example','first','last','first.last@example.com')`
		result := processInsert(dataInput)

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
		dataInput := "table_name:User,column_name:id|bhf1234584,column_name:username|example,column_name:first_name|first,column_name:last_name|last,column_name:email|first.last@example.com,column_name:secret_detail"
		expectedQuery := `UPDATE users SET username='example',first_name='first',last_name='last',email='first.last@example.com' WHERE id='bhf1234584'`
		result := processUpdate(dataInput)

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
		var data = `{table_name:User,column_name:id|bhf1234584,column_name:username|,column_name:first_name|,column_name:last_name|,column_name:email|,column_name:secret_detail{table_name:SecretDetailEx,column_name:id|11244,column_name:api_secret|,column_name:api_token|}}`

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
			t.Errorf("%s expected length of components == 2, got %d", failed, len(components))
		} else {
			t.Logf("%s expected length of components == 2", success)
		}

		expectedFirst := "table_name:SecretDetailEx,column_name:id|11244,column_name:api_secret|,column_name:api_token|"
		expectedSecond := "table_name:User,column_name:id|bhf1234584,column_name:username|,column_name:first_name|,column_name:last_name|,column_name:email|,column_name:secret_detail"

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
		var data = `{table_name:OrderEx,column_name:Id|o123,column_name:Updated|2018-06-17 03:43:33,column_name:Created|2018-06-17 03:43:33,column_name:Status|OrderCreated,column_name:Items{table_name:OrderItem,column_name:Id|itm123,column_name:ItemName|XL 2 Giga,column_name:Price|150000}{table_name:OrderItem,column_name:Id|itm124,column_name:ItemName|XL 5 Giga,column_name:Price|300000}}`
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
