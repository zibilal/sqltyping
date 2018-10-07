package sqltyping

import (
	"database/sql"
	"testing"
	"encoding/json"
)

func TestTypeIterator_AccepSqlNullInt64(t *testing.T) {
	t.Log("Testing Accept SqlNilInt")
	{
		input := struct {
			Name  sql.NullString
			Rate  sql.NullFloat64
			Level sql.NullInt64
		}{
			Name:  sql.NullString{String: "Test Name"},
			Rate:  sql.NullFloat64{Float64: 12.56},
			Level: sql.NullInt64{Int64: 15},
		}

		output := struct {
			Name  string
			Rate  float64
			Level int64
		}{}

		err := TypeIterator(input, &output)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(output) {
			t.Fatalf("%s expected output not empty", failed)
		} else {
			b, _ := json.MarshalIndent(output, "", "\t")
			t.Logf("%s expected output not empty, got %s", success, string(b))
		}
	}
}
