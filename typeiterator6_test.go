package sqltyping

import (
	"testing"
)

type APIResponse struct {
	HTTPCode int         `json:"-"`
	Code     int         `json:"code"`
	Message  interface{} `json:"message"`
	Data     interface{} `json:"data,omitempty"`
}

type makePaymentResponse struct {
	Customer        customerMakePaymentResponse `json:"customer"`
	Items           []itemMakePaymentResponse   `json:"order_items"`
	OrderDate       string                      `json:"order_date"`
	PaymentInfo     interface{}                 `json:"payment_info"`
	Reference       string                      `json:"reference"`
	ShippingCost    string                      `json:"shipping_cost"`
	TotalCommission int                         `json:"total_commission"`
	TotalPaid       int                         `json:"total_paid"`
}

type itemMakePaymentResponse struct {
	Commission string `json:"item_komisi"`
	ItemImage  string `json:"item_image"`
	Name       string `json:"item_name"`
	Price      string `json:"price"`
	Quantity   string `json:"quantity"`
}

type customerMakePaymentResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Phonenumber string `json:"phonenumber"`
}

func TestTypeIteratorWitNestedInterfaceStruct(t *testing.T) {
	t.Log("Test TypeIterator with nested interface struct")
	{
		sRef := struct {
			Message struct {
				Reference string `json:"reference"`
			} `json:"message"`
		}{struct {
			Reference string `json:"reference"`
		}{
			"2018112307070200253...",
		},

		}

		resp := APIResponse{
			Message: makePaymentResponse{
				Reference: "2018112307070200253",
			},
		}

		TypeIterator(sRef, &resp)

		t.Log(resp)
	}
}
