package sqltyping

import (
	"testing"
	"encoding/json"
)

// Types related to this service
type StoreOrder struct {
	CartID           string            `json:"cart_id" valid:"funcVal:Required"`
	CategoryType     int               `json:"category_type"`
	Phone            string            `json:"phone" valid:"funcVal:Required"`
	Items            interface{}       `json:"items" valid:"funcVal:Required"`
	OrderSplitID     string            `json:"order_split_id" transform:"order_split_id"`
	DetailType       int               `json:"detail_type"`
	QueueNo          string            `json:"queue_no"`
	Reference        string            `json:"reference"`
	KudoboxID        string            `json:"kudobox_id"`
	SpotID           string            `json:"spot_id"`
	TotalPaid        string            `json:"total_paid"`
	KudoFee          string            `json:"kudo_fee"`
	ShippingCharges  string            `json:"shipping_charges"`
	ExpiryDate       string            `json:"expiry_date"`
	CreatedAt        string            `json:"created_at"`
	EventLaunchingID string            `json:"event_launching_id"`
	StatusTrx        string            `json:"status_trx"`
	OrderPaymentType int               `json:"order_payment_type"`
	Vhide            bool              `json:"vhide"`
	LastStatus       string            `json:"last_status"`
	Email            string            `json:"email,omitempty"`
	Shipping         interface{}       `json:"shipping,omitempty"`
	AgentID          string            `json:"-" transform:"agent_id"`
	CustomerID       uint64            `json:"-" transform:"customer_id"`
	AgentInfo        UserInfo          `json:"-"`
	Headers          map[string]string `json:"headers"`
	DeviceID         string            `json:"-"`
	Channel          string            `json:"-"`
}

type StoreOrderItem struct {
	Attributes            string  `json:"attributes,omitempty" transform:"attributes"`
	ItemCommission        float64 `json:"item_komisi" transform:"commission"`
	IsBukalapakItem       bool    `json:"is_bukalapak_item,omitempty"`
	ItemID                uint64  `json:"item_id,omitempty" transform:"item_id"`
	ImageURL              string  `json:"item_image,omitempty" transform:"item_image"`
	ItemName              string  `json:"item_name,omitempty" transform:"item_name"`
	Price                 float64 `json:"price,omitempty" transform:"price"`
	ItemReferenceID       string  `json:"item_reference_id,omitempty" transform:"item_reference_id"`
	MaxSku                int     `json:"max_sku,omitempty"`
	OriginalPrice         int     `json:"original_price,omitempty"`
	Pack                  int     `json:"pack"`
	Quantity              int     `json:"quantity,omitempty" transform:"quantity"`
	RequireAddress        bool    `json:"required_address,omitempty"`
	RequireNote           bool    `json:"required_note,omitempty"`
	VendorID              uint64  `json:"vendor_id,omitempty" transform:"vendor_id"`
	Wholesale             int     `json:"whole_sale" transform:"whole_sale"`
	WholesaleState        bool    `json:"whole_state,omitempty" transform:"whole_state"`
	RetailPrice           float64 `json:"retail_price" transform:"retail_price"`
	Note                  string  `json:"note"`
	Status                string  `json:"status"`
	RefundShippingCharges string  `json:"refund_shipping_charges,default:0"`

	Description string `json:"-" transform:"description"`

	ResellerPrice float64 `json:"-"`
}

// UserInfo struct
type UserInfo struct {
	ActiveStatus int    `json:"active_status"`
	AgentID      string `json:"cashier_id"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	IsAgent      bool   `json:"is_agent"`
	LastName     string `json:"last_name"`
	Nis          string `json:"nis"`
	Phonenumber  string `json:"phonenumber"`
}

type Shipping struct {
	Courier       string  `json:"courier,omitempty"`
	Save          bool    `json:"save,omitempty"`
	ShippingCost  float64 `json:"shipping_cost,omitempty"`
	ShippingType  uint    `json:"shipping_type,omitempty"`
	Weight        float64 `json:"weight,omitempty"`
	Address       string  `json:"address,omitempty"`
	City          string  `json:"city,omitempty"`
	RecipientName string  `json:"recipient_name,omitempty"`
	District      string  `json:"kecamatan,omitempty"`
	Province      string  `json:"province,omitempty"`
	Email         string  `json:"email,omitempty"`
	Village       string  `json:"kelurahan,omitempty"`
	Postcode      string  `json:"postcode,omitempty"`
	PhoneNumber   string  `json:"phone_number,omitempty"`
	VillageID     uint    `json:"kelurahan_id,omitempty"`
	DistrictID    uint    `json:"kecamatan_id,omitempty"`
	ProvinceID    uint    `json:"province_id,omitempty"`
	CityID        uint    `json:"city_id,omitempty"`
}


type StoreOrderRequest struct {
	CartID       string                  `json:"cart_id" valid:"funcVal:Required;funcVal:Match,format:^[a-zA-Z0-9]+$"`
	CategoryType int                     `json:"category_type"`
	Phone        string                  `json:"phone" valid:"funcVal:Required;funcVal:Match,format:^[0-9]+$"`
	Items        []StoreOrderItemRequest `json:"items" valid:"funcVal:Required"`
	Email        string                  `json:"email" valid:"funcVal:Email"`
	Shipping     Shipping                `json:"shipping,omitempty"`
	Headers      map[string]string       `json:"headers"`
	DeviceID     string                  `json:"-"`
	Channel      string                  `json:"-"`
}

type StoreOrderItemRequest struct {
	Attributes      string  `json:"attributes" transform:"attributes"`
	ItemCommission  float32 `json:"item_komisi" transform:"commission"`
	IsBukalapakItem bool    `json:"mIsBukalapakItem"`
	ItemID          uint64  `json:"item_id" valid:"funcVal:Required" transform:"item_id"`
	ImageURL        string  `json:"image_url" transform:"item_image"`
	ItemName        string  `json:"item_name" transform:"item_name"`
	Price           float64 `json:"price" transform:"price"`
	ItemReferenceID string  `json:"item_reference_id" transform:"item_reference_id"`
	MaxSku          int     `json:"mMaxSku" transform:"max_sku"`
	OriginalPrice   int     `json:"mOriginalPrice" transform:"original_price"`
	Pack            int     `json:"pack" transform:"pack"`
	Quantity        int     `json:"quantity" valid:"funcVal:Required" transform:"quantity"`
	RequireAddress  bool    `json:"mRequireAddress" transform:"require_address"`
	RequireNote     bool    `json:"mRequireNote" transform:"require_note"`
	VendorID        uint64  `json:"vendor_id" transform:"vendor_id" valid:"funcVal:Required"`
	Wholesale       int     `json:"wholesale" transform:"whole_sale"`
	WholesaleState  bool    `json:"mWholesaleState" transform:"whole_sale_state"`
}

type ShippingRequest struct {
	Courier       string  `json:"courier,omitempty"`
	Save          bool    `json:"save,omitempty"`
	ShippingCost  float64 `json:"shipping_cost,omitempty"`
	ShippingType  uint    `json:"shipping_type,omitempty"`
	Weight        float64 `json:"weight,omitempty"`
	Address       string  `json:"address,omitempty"`
	City          string  `json:"city,omitempty"`
	RecipientName string  `json:"recipient_name,omitempty"`
	District      string  `json:"kecamatan,omitempty"`
	Province      string  `json:"province,omitempty"`
	Email         string  `json:"email,omitempty"`
	Village       string  `json:"kelurahan,omitempty"`
	Postcode      string  `json:"postcode,omitempty"`
	PhoneNumber   string  `json:"phone_number,omitempty"`
	VillageID     uint    `json:"kelurahan_id,omitempty"`
	DistrictID    uint    `json:"kecamatan_id,omitempty"`
	ProvinceID    uint    `json:"province_id,omitempty"`
	CityID        uint    `json:"city_id,omitempty"`
}

var jsonMsg = `
{
    "cart_id": "00004",
    "items": [{
        "attributes": "{\"customer\":{\"cellphone_number\":\"081220172375\"},\"purchase_referral\":{\"action\":\"pulsa\"}}",
        "item_komisi": 0.0,
        "mIsBukalapakItem": false,
        "item_id": 26175079,
        "image_url": "https://dev-static.kudo.co.id/api/images/category/ic_pulsa_telkomsel.png",
        "item_name": "Voucher Rp25.000 (081220172375)",
        "price": 26000,
        "item_reference_id": "NARINDO_AIRTIME_OTS25",
        "mMaxQty": 1,
        "mMaxSku": 1,
        "mMinQty": 1,
        "quantity": 2,
        "mRequireAddress": false,
        "mRequireNote": false,
        "vendor_id": 83
    }],
    "phone": "081220172375"
}
`
func TestTypeIterator4(t *testing.T) {
	t.Log("Test type iterator 4")
	{
		soRequest := StoreOrderRequest{}
		err := json.Unmarshal([]byte(jsonMsg), &soRequest)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error() )
		} else {
			t.Logf("%s expected error nil", success)
		}
		soRequest.Headers = map[string]string {
			"data1": "value1",
		}

		storeOrder := StoreOrder{
		}
		storeOrder.Headers = make(map[string]string)

		err = TypeIterator(soRequest, &storeOrder)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error() )
		} else {
			t.Logf("%s expected error nil", success)
		}

		b, _ := json.MarshalIndent(storeOrder, "", "\t")
		b2, _ := json.MarshalIndent(soRequest, "", "\t")
		t.Logf("%s RESULT: %s", success, string(b))
		t.Logf("%s RESULT: %s", success, string(b2))
		t.Logf("headers %v", soRequest.Headers)
	}
}


