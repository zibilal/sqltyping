package sqltyping

import (
	"bytes"
	"strings"
	"testing"
	"text/scanner"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestSqlTyping(t *testing.T) {
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
}

func TestTescan(t *testing.T) {

	t.Log("Testing text scan")
	{
		var data = `{ User|| id|bhf1234584  username|  first_name|  last_name|  email|  secret_detail { SecretDetailEx|| id|11244  api_secret|  api_token|  } }`
		var s scanner.Scanner
		s.Init(strings.NewReader(data))

		b := make([]byte, 0)
		buff := bytes.NewBuffer(b)

		fromCount := 0
		//columnCount := 0
		whereOpenCount := 0

		var tok rune
		for tok != scanner.EOF {
			tok = s.Scan()
			switch tok {
			case scanner.Ident:
				if fromCount == 1 && bytes.Contains(buff.Bytes(), []byte("FROM")) {
					buff.WriteString( " " + strings.ToLower(s.TokenText()) + "s ")
					fromCount--
				} else {
					t.Logf("Ident = %d: %s", s.Position.Column, s.TokenText())
				}
			default:
				if fromCount == 0 && s.TokenText() == "{" {
					buff.WriteString("FROM")
					fromCount++
				} else if s.TokenText() == "|" {
					whereOpenCount++
				} else {
					t.Logf("Default = %d: %s", s.Position.Column, s.TokenText())
				}
			}
		}

		t.Logf("Result %s", buff.String())
	}

}
