package sqltyping

import (
	"bytes"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

/*func TestSqlTyping(t *testing.T) {
	t.Log("Testing sql typing")
	{
		expectedSql := `SELECT id, username, first_name, last_name, email FROM users WHERE id = 'bhf1234584'`

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

		if expectedSql == results[0] {
			t.Logf("%s expected query = %s", success, expectedSql)
		} else {
			t.Errorf("%s expected query = %s, got %s", failed, expectedSql, results[0])
		}
	}
}*/

func TestTescan(t *testing.T) {

	t.Log("Testing text scan")
	{
		var data = `{table_name:User,column_name:id|bhf1234584,column_name:username|,column_name:first_name|,column_name:last_name|,column_name:email|,column_name:secret_detail{table_name:SecretDetailEx,column_name:id|11244,column_name:api_secret|,column_name:api_token|}}`

		bytesData := []byte(data)
		buff := bytes.NewBuffer(bytesData)
		t.Log("LEN: Buff", len(buff.Bytes()) )
		components, err := processBytes(buff)

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

		expectedFirst := "table_name:SecretDetailEx,column_name:id|11244,column_name:api_secret|,column_name:api_token"
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
		buff := bytes.NewBuffer(bytesData)
		components, err := processBytes(buff)

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

		expectedFirst := "table_name:SecretDetailEx,column_name:id|11244,column_name:api_secret|,column_name:api_token"
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
