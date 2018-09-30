package sqltyping

import (
	"reflect"
	"testing"
)

type ApplyCampaignParam struct {
	Header  ApplyCampaignParamHeader `json:"header"`
	Body    ApplyCampaignParamBody   `json:"body"`
	AgentID string                   `json:"agent_id"`
}

type ApplyCampaignParamHeader struct {
	Authorization string `json:"authorization"`
	Channel       string `json:"channel"`
	ClientVersion string `json:"client_version"`
	DeviceID      string `json:"device_id"`
}
type ApplyCampaignParamBody struct {
	OrderSplitID string `json:"order_split_id"`
	CampaignCode string `json:"campaign_code"`
}

type Customer struct {
	FullName string
	Address  string
	Rate     float64
}

type OrderNew struct {
	OrderObjectId uint64
	Status        string
	OrderItems    []Item
	Client        Customer
}

func TestSearchValue(t *testing.T) {
	t.Log("Testing CopyValue function")
	{
		requestData := ApplyCampaignParam{}
		requestHeader := ApplyCampaignParamHeader{
			"fdadfsfsfsdsdfsfsasda", "MOBILE", "7.0.8", "313414313414",
		}
		requestBody := ApplyCampaignParamBody{
			"11123431111", "V0001",
		}
		requestData.Header = requestHeader
		requestData.Body = requestBody

		result := searchValue(requestData, "OrderSplitID", "", nil)
		if result == nil {
			t.Fatalf("%s expected result not nil", failed)
		} else {
			if reflect.DeepEqual(result, "11123431111") {
				t.Logf("%s expected result = \"11123431111\"", success)
			} else {
				t.Fatalf("%s expected result = \"11123431111\"", failed)
			}
		}

		result = searchValue(requestData, "ClientVersion", "", nil)
		if result == nil {
			t.Fatalf("%s expected result not nil", failed)
		} else {
			if reflect.DeepEqual(result, "7.0.8") {
				t.Logf("%s expected result = \"7.0.8\"", success)
			} else {
				t.Fatalf("%s expected result = \"7.0.8\"", failed)
			}
		}

		result = searchValue(requestData, "ClientVersions", "", nil)
		if result == nil {
			t.Logf("%s expected result == nil", success)
		} else {
			t.Fatalf("%s expected result == nil, got %+v", failed, result)
		}

	}
}

func TestCopyValue(t *testing.T) {
	t.Log("Testing copy value first with literal struct value")
	{
		requestData := ApplyCampaignParam{}
		requestHeader := ApplyCampaignParamHeader{
			"fdadfsfsfsdsdfsfsasda", "MOBILE", "7.0.8", "313414313414",
		}
		requestBody := ApplyCampaignParamBody{
			"11123431111", "V0001",
		}
		requestData.Header = requestHeader
		requestData.Body = requestBody

		output := struct {
			OrderSplitID string
			CampaignCode string
		}{}

		err := CopyValue(requestData, &output)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(output) {
			t.Fatalf("%s expected output is not empty", failed)
		} else {
			t.Logf("%s expected output is not empty, got %+v", success, output)
		}
	}
	t.Log("Testing copy value first with literal struct value")
	{
		requestData := ApplyCampaignParam{}
		requestHeader := ApplyCampaignParamHeader{
			"fdadfsfsfsdsdfsfsasda", "MOBILE", "7.0.8", "313414313414",
		}
		requestBody := ApplyCampaignParamBody{
			"11123431111", "V0001",
		}
		requestData.Header = requestHeader
		requestData.Body = requestBody

		output := struct {
			Authorization string
			CampaignCode  string
		}{}

		err := CopyValue(requestData, &output)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(output) {
			t.Fatalf("%s expected output is not empty", failed)
		} else {
			t.Logf("%s expected output is not empty, got %+v", success, output)
		}
	}
}

func TestCopyValueTheSecond(t *testing.T) {
	t.Log("Testing CopyValue with other type struct")
	{
		data := OrderNew{}

		items := []Item{
			{
				Id: "001",
				ItemName: "Pulsa1",
				Price: 12000,
			},
			{
				Id: "002",
				ItemName: "Pulsa2",
				Price: 36000,
			},
			{
				Id: "003",
				ItemName: "Pulsa3",
				Price: 24000,
			},
		}

		customer := Customer{}
		customer.FullName = "Customer One"
		customer.Address = "Customer Address"
		customer.Rate = 4.84

		data.OrderItems = items
		data.OrderObjectId = 120024
		data.Status = "Order Success"
		data.Client = customer

		output := struct {
			FullName string
			Rate float64
			OrderItems []Item
		}{}

		err := CopyValue(data, &output)
		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error nil", success)
		}

		if IsEmpty(output) {
			t.Fatalf("%s expected output not empty", failed)
		} else {
			t.Logf("%s expected output not empty, got %+v", success, output)
		}

	}
}
