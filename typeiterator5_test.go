package sqltyping

import (
	"bitbucket.org/kudoindonesia/microservice_order/helpers"
	"database/sql"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestTypeIterator_AccepSqlNulls(t *testing.T) {
	t.Log("Testing Accept SqlNilInt")
	{
		input := struct {
			Name      sql.NullString
			Rate      sql.NullFloat64
			Level     sql.NullInt64
			CreatedAt mysql.NullTime
			Status    sql.NullInt64
		}{
			Name:      sql.NullString{String: "Test Name"},
			Rate:      sql.NullFloat64{Float64: 12.56},
			Level:     sql.NullInt64{Int64: 15},
			CreatedAt: mysql.NullTime{Time: time.Now()},
			Status:    sql.NullInt64{Int64: 5},
		}

		output := struct {
			Name      string
			Rate      float64
			Level     int64
			CreatedAt time.Time
			Status    uint8
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

func TestTypeIterator_OutputSqlNulls(t *testing.T) {
	t.Log("Testing for giving data to SqlNull type variants")
	{
		input := struct {
			ID        string
			Name      string
			Rate      float64
			Level     int64
			Status    uint8
			CreatedAt time.Time
		}{
			ID:        "00dasfa0001",
			Name:      "Test Name",
			Rate:      12.56,
			Level:     12,
			Status:    2,
			CreatedAt: time.Now(),
		}

		output := struct {
			ID        sql.NullInt64
			Name      sql.NullString
			Rate      sql.NullFloat64
			Level     sql.NullInt64
			Status    sql.NullInt64
			CreatedAt mysql.NullTime
		}{}

		err := TypeIterator(input, &output)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(output) {
			t.Fatalf("%s expected output is not empty", failed)
		} else {
			b, _ := json.MarshalIndent(output, "", "\t")
			t.Logf("%s expected output is not empty, got %s", success, string(b))
		}
	}
}

func TestTypeIteratorUpdateStructFromStruct(t *testing.T) {
	t.Log("Testing for update struct from a struct")
	{
		ex := ExampleType{
			Id: "111223",
			Type: "simple",
			Reference: "2018112211212200123455678",
		}

		aref := struct {
			Reference string
		}{
			"20181122112122001...",
		}

		err := TypeIterator(aref, &ex)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if ex.Reference == aref.Reference {
			t.Logf("%s expected reference %s", success, aref.Reference)
		} else {
			t.Fatalf("%s expected reference %s, got %s", failed, aref.Reference, ex.Reference)
		}
	}
}

type ExampleType struct {
	Id string
	Type string
	Reference string
}

type ExampleEvent struct {
	Title string
	Text string
	Timestamp time.Time
	Hostname string
	AggregationKey string
	SourceTypeName string
	Tags []string
}

func TestEventDataType(t *testing.T) {
	t.Log("Testing event data type")
	{
		evt := ExampleEvent{}

		evtMap := map[string]interface{} {
			"Title": "event title",
			"Text": "Event Description of the Event to describe it",
			"Timestamp": time.Now(),
			"AggregationKey": "event category",
			"SourceTypeName": "source type name",
			"Tags": []string{"tags.01", "tags.02", "tags.03"},
		}

		err := TypeIterator(evtMap, &evt)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if helpers.IsEmpty(evt) {
			t.Fatalf("%s expected evt is not empty", failed)
		} else {
			b, _ := json.MarshalIndent(evt, "", "\t")
			t.Logf("%s expected evt is not empty, value:%s", success, string(b))
		}
	}
}