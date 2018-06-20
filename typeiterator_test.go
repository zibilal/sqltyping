package sqltyping

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestTypeIterator(t *testing.T) {
	t.Log("Testing type iterator")
	{
		ival1 := 20
		oval1 := 0

		err := TypeIterator(ival1, &oval1)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if ival1 != oval1 {
			t.Fatalf("%s expected oval == ival = %v, got %v", failed, ival1, oval1)
		} else {
			t.Logf("%s expected oval == ival = %v, got %v", success, ival1, oval1)
		}

		ival2 := 20.50
		oval2 := 0.0

		err = TypeIterator(ival2, &oval2)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if ival2 != oval2 {
			t.Fatalf("%s expected oval == ival = %v, got %v", failed, ival2, oval2)
		} else {
			t.Logf("%s expected oval == ival = %v, got %v", success, ival2, oval2)
		}

		ival3 := "dua puluh lima"
		oval3 := ""

		err = TypeIterator(ival3, &oval3)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		t.Logf("ival %v : oval %v", ival3, oval3)
	}
}

func TestTypeIteratorWithStructInputOutput(t *testing.T) {

	t.Log("testing input and output a simple struct")
	{
		ival := struct {
			Name  string  `json:"name"`
			Point float64 `json:"point"`
		}{
			"Test1", 12.2,
		}

		oval := struct {
			FullName string  `val:"name"`
			Point    float64 `val:"point"`
		}{}

		err := TypeIterator(ival, &oval)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if oval.FullName == ival.Name {
			t.Logf("%s expected oval.FullName = ival.Name", success)
		} else {
			t.Fatalf("%s expected oval.FullName = ival.Name, got 1: %s, 2: %s", failed, oval.FullName, ival.Name)
		}

		if oval.Point == ival.Point {
			t.Logf("%s expected oval.Point = ival.Point", success)
		} else {
			t.Fatalf("%s expected oval.Point = ival.Point, got 1: %f, 2: %f", failed, oval.Point, ival.Point)
		}
	}

	t.Log("testing input and output nested struct")
	{
		ival := struct {
			Name    string
			Point   float64
			Profile interface{}
		}{
			Name:  "Example 1",
			Point: 56.89,
			Profile: struct {
				DeviceId  string
				ApiSecret string
				Document  string
			}{
				"1231341234", "api-bbfdsf32324df343", "http://example.com/doc.pdf",
			},
		}

		oval := struct {
			Name    string
			Point   float64
			Profile interface{}
		}{}

		err := TypeIterator(ival, &oval)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(oval) {
			t.Fatalf("%s expected oval not empty", failed)
		} else {
			b, err := json.MarshalIndent(oval, "", "\t")
			if err != nil {
				t.Fatalf("%s expected error nil, got %s", failed, err.Error())
			} else {
				t.Logf("%s expected rror nil", success)
			}

			t.Logf("Output %s", string(b))
		}
	}

	t.Log("input with nested struct")
	{
		userEx := UserExample{
			FirstName: "firstex",
			LastName:  "lastex",
			Email:     "firstduo@example.com",
			Title:     "title",
			Authentication: Authentication{
				Username:  "xiexample",
				APISecret: "apisecret",
				APIToken:  "apitoken",
				ServiceDetail: ServiceDetail{
					Service: "aservicenm",
					Cost:    15000.00,
				},
			},
		}

		profEx := ProfileExample{}

		err := TypeIterator(userEx, &profEx)
		if err != nil {
			t.Fatalf("%s expected err not nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error not nil", success)
		}

		if IsEmpty(profEx) {
			t.Fatalf("%s expected profEx not empty", failed)
		} else {
			b, err := json.MarshalIndent(profEx, "", "\t")
			if err != nil {
				t.Fatalf("%s expected error nil, got %s", failed, err.Error())
			} else {
				t.Logf("%s expected error nil", success)
			}

			t.Logf("%s", string(b))
		}
	}

	t.Log("input with nested struct, bytes.Buffer output")
	{
		userEx := UserExample{
			FirstName: "firstex",
			LastName:  "lastex",
			Email:     "firstduo@example.com",
			Title:     "title",
			Authentication: Authentication{
				Username:  "xiexample",
				APISecret: "apisecret",
				APIToken:  "apitoken",
				ServiceDetail: ServiceDetail{
					Service: "aservicenm",
					Cost:    15000.00,
				},
			},
		}

		buff := bytes.NewBufferString("")

		err := TypeIterator(userEx, buff)
		if err != nil {
			t.Fatalf("%s expected err not nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error not nil", success)
		}

		if buff.String() == "" {
			t.Fatalf("%s expected buff not empty, got empty", failed)
		} else {
			t.Logf("%s expected buff not empty got %s", success, buff.String())
		}
	}

	t.Log("input with empty nested struct, bytes.Buffer output")
	{
		userEx := UserExample{}

		buff := bytes.NewBufferString("")

		err := TypeIterator(userEx, buff)
		if err != nil {
			t.Fatalf("%s expected err not nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error not nil", success)
		}

		if buff.String() == "" {
			t.Fatalf("%s expected buff not empty, got empty", failed)
		} else {
			t.Logf("%s expected buff not empty got %s", success, buff.String())
		}
	}

}

func TestTypeIteratorWithMapInput(t *testing.T) {
	t.Log("Testing TypeIterator with map input, output bytes.Buffer")
	{
		map1 := make(map[string]ServiceDetail)
		map1["1"] = ServiceDetail{}
		map1["2"] = ServiceDetail{}
		map1["3"] = ServiceDetail{
			"Premium MS Spring", 30000,
		}

		map2 := bytes.NewBufferString("")
		err := TypeIterator(map1, map2)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(map2) {
			t.Fatalf("%s expected map2 not empty", failed)
		} else {
			if err != nil {
				t.Fatalf("%s expected error empty, got %s", success, err.Error())
			}
			t.Logf("%s expected map2 not empty, result: %s", success, map2.String())
		}
	}

	t.Log("Testing TypeIterator with map input, output map of struct")
	{
		map1 := make(map[string]ServiceDetail)
		map1["1"] = ServiceDetail{
			"Premium MS Order", 15000,
		}
		map1["2"] = ServiceDetail{
			"Premium MS Deposit", 25000,
		}
		map1["3"] = ServiceDetail{
			"Premium MS Spring", 30000,
		}

		map2 := make(map[string]ServiceCost)

		err := TypeIterator(map1, &map2)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		//if IsEmpty(map2) || len(map2) == 0 {
		if IsEmpty(map2) {
			t.Fatalf("%s expected map2 not empty", failed)
		} else {
			b, err := json.MarshalIndent(map2, "", "\t")
			if err != nil {
				t.Fatalf("%s expected error empty, got %s", success, err.Error())
			}
			t.Logf("%s expected map2 not empty, result: %s", success, string(b))
		}
	}
}

func TestTypeIteratorStructBuffer(t *testing.T) {
	t.Log("Testing TypeIterator input struct output buffers")
	{
		order := OrderExample{
			Id: 12, Created: time.Now(), Updated: time.Now(), Status: "created",
		}
		buff := bytes.NewBufferString("")
		err := TypeIterator(order, buff)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if buff.String() == "" {
			t.Fatalf("%s expected buff string not empty", failed)
		} else {
			t.Logf("%s expected buff string not empty, result: \n%s", success, buff.String())
		}
	}

	t.Log("Testing TypeIterator input struct output buffers, user struct")
	{
		user := User{
			Id: "bhf1234584",
			SecretDetail: SecretDetailEx{
				Id: "11244",
			},
		}

		buff := bytes.NewBufferString("")
		err := TypeIterator(user, buff)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if buff.String() == "" {
			t.Fatalf("%s expected buff string not empty", failed)
		} else {
			t.Logf("%s expected buff string not empty, result: \n%s", success, buff.String())
		}
	}
}

func TestTypeIteratorStructBufferSecond(t *testing.T) {
	t.Log("Testing type iterator input struct with a slice of struct, output bytes buffer")
	{
		order := OrderEx{
			Id:      "o123",
			Created: time.Now(),
			Updated: time.Now(),
			Status:  "OrderCreated",
			Items: []OrderItem{
				{
					Id:       "itm123",
					ItemName: "XL 2 Giga",
					Price:    150000,
				}, {
					Id:       "itm124",
					ItemName: "XL 5 Giga",
					Price:    300000,
				},
			},
		}

		buff := bytes.NewBufferString("")
		err := TypeIterator(order, buff)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if buff.String() == "" {
			t.Fatalf("%s expected not empty buff result, got empty buff", failed)
		} else {
			t.Logf("%s expected not empty buff result, buff result = %s", success, buff.String())
		}
	}
}

func TestTypeIteratorStructWithSliceToStructWithSlice(t *testing.T) {
	t.Log("Struct with interface")
	{
		order := OrderEx{
			Id:      "o123",
			Created: time.Now(),
			Updated: time.Now(),
			Status:  "OrderCreated",
			Items: []OrderItem{
				{
					Id:       "itm123",
					ItemName: "XL 2 Giga",
					Price:    150000,
				}, {
					Id:       "itm124",
					ItemName: "XL 5 Giga",
					Price:    300000,
				},
			},
		}

		orderOutput := struct {
			Id      string
			Created time.Time
			Updated time.Time
			Status  string
			Items   interface{}
		}{}

		err := TypeIterator(order, &orderOutput)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(orderOutput) {
			t.Fatalf("%s expected orderOutput not empty", failed)
		} else {
			b, err := json.MarshalIndent(orderOutput, "", "\t")
			if err != nil {
				t.Fatalf("%s expected error nil, got %s", failed, err.Error())
			}
			t.Logf("%s expected orderOutput not emtpy, got %s", success, string(b))
		}
	}

	t.Log("Struct with slice of struct")
	{
		order := OrderEx{
			Id:      "o123",
			Created: time.Now(),
			Updated: time.Now(),
			Status:  "OrderCreated",
			Items: []OrderItem{
				{
					Id:       "itm123",
					ItemName: "XL 2 Giga",
					Price:    150000,
				}, {
					Id:       "itm124",
					ItemName: "XL 5 Giga",
					Price:    300000,
				},
			},
		}

		orderOutput := struct {
			Id      string
			Created time.Time
			Updated time.Time
			Status  string
			Items   []Item
		}{}

		err := TypeIterator(order, &orderOutput)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(orderOutput) {
			t.Fatalf("%s expected orderOutput not empty", failed)
		} else {
			b, err := json.MarshalIndent(orderOutput, "", "\t")
			if err != nil {
				t.Fatalf("%s expected error nil, got %s", failed, err.Error())
			}
			t.Logf("%s expected orderOutput not emtpy, got %s", success, string(b))
		}
	}
}

type OrderEx struct {
	Id      string
	Updated time.Time
	Created time.Time
	Status  string
	Items   []OrderItem
}

type OrderItem struct {
	Id       string
	ItemName string
	Price    float64
}

type Item struct {
	Id       string
	ItemName string
	Price    float64
}

type User struct {
	Id           string         `sqltype:"id"`
	Username     string         `sqltype:"username"`
	FirstName    string         `sqltype:"first_name"`
	LastName     string         `sqltype:"last_name"`
	Email        string         `sqltype:"email"`
	SecretDetail SecretDetailEx `sqltype:"secret_detail"`
}

type SecretDetailEx struct {
	Id        string `json:"id"`
	APISecret string `json:"api_secret"`
	APIToken  string `json:"api_token"`
}

type OrderExample struct {
	Id      int64
	Updated time.Time
	Created time.Time
	Status  string
}

type ProfileExample struct {
	FirstName      string
	LastName       string
	Email          string
	Title          string
	Authentication SecretDetail
}

type SecretDetail struct {
	APISecret string `json:"api_secret"`
	APIToken  string `json:"api_token"`
}

type UserExample struct {
	FirstName      string
	LastName       string
	Email          string
	Title          string
	Authentication Authentication
}

type Authentication struct {
	Username      string
	APISecret     string
	APIToken      string
	ServiceDetail ServiceDetail
}

type ServiceDetail struct {
	Service string  `ksmg:"service_name"`
	Cost    float64 `kmsg:"service_cost"`
}

type ServiceCost struct {
	Service string  `json:"service_name"`
	Cost    float64 `json:"service_cost"`
}
