package sqltyping

import (
	"testing"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
	"bytes"
)

type CoreOrderItems4 struct {
	AddPartnerPrice       sql.NullInt64   `db:"add_partner_price"`
	AgentDistributorID    sql.NullInt64   `db:"agent_distributor_id"`
	Attributes            sql.NullString  `db:"attributes"`
	CashierID             sql.NullInt64   `db:"cashier_id"`
	CollectionID          sql.NullInt64   `db:"collection_id"`
	CreatedAt             mysql.NullTime  `db:"created_at"`
	Description           sql.NullString  `db:"description"`
	DeviceID              sql.NullString  `db:"device_id"`
	ID                    sql.NullInt64   `db:"id"`
	ItemCardCharging      sql.NullFloat64 `db:"item_card_charging"`
	ItemDeleteStatus      sql.NullString  `db:"item_delete_status"`
	ItemDiscount          sql.NullFloat64 `db:"item_discount"`
	ItemDiscountShipping  sql.NullFloat64 `db:"item_discount_shipping"`
	ItemImage             sql.NullString  `db:"item_image"`
	ItemKomisi            sql.NullFloat64 `db:"item_komisi"`
	ItemKudoFee           sql.NullFloat64 `db:"item_kudo_fee"`
	ItemName              sql.NullString  `db:"item_name" transform:"item_name"`
	ItemReferenceID       sql.NullString  `db:"item_reference_id"`
	ItemShipping          sql.NullFloat64 `db:"item_shipping"`
	KudoboxID             sql.NullInt64   `db:"kudobox_id"`
	MerchantTrxID         sql.NullString  `db:"merchant_trx_id"`
	OrderID               sql.NullInt64   `db:"order_id"`
	OrderStep             sql.NullInt64   `db:"order_step"`
	PartnerPlu            sql.NullString  `db:"partner_plu"`
	Price                 sql.NullFloat64 `db:"price"`
	PurchaseCode          sql.NullString  `db:"purchase_code"`
	QtyUpdateStatus       sql.NullString  `db:"qty_update_status"`
	Quantity              sql.NullInt64   `db:"quantity"`
	RefundShippingCharges sql.NullInt64   `db:"refund_shipping_charges"`
	ResellerPrice         sql.NullFloat64 `db:"reseller_price"`
	RetryProcess          sql.NullInt64   `db:"retry_process"`
	ShippingCode          sql.NullString  `db:"shipping_code"`
	ShopItemID            sql.NullInt64   `db:"shop_item_id"`
	Sms                   sql.NullInt64   `db:"sms"`
	SpotID                sql.NullInt64   `db:"spot_id"`
	Status                sql.NullInt64   `db:"status"`
	UpdatedAt             mysql.NullTime  `db:"updated_at"`
	VendorID              sql.NullInt64   `db:"vendor_id"`
}

func TestTyping(t *testing.T) {
	it := CoreOrderItems4{}
	it.AddPartnerPrice = sql.NullInt64{Int64:0}
	it.AgentDistributorID = sql.NullInt64{Int64:0}
	it.Attributes = sql.NullString{String: "{\"customer\":{\"cellphone_number\":\"081220172375\"},\"purchase_referral\":{\"action\":\"pulsa\"}}"}
	it.CashierID = sql.NullInt64{Int64:5246686}
	it.CollectionID = sql.NullInt64{Int64:0}
	it.CreatedAt = mysql.NullTime{Time:time.Now()}
	it.Description = sql.NullString{String:"Voucher Rp100.000 (Test)"}
	it.DeviceID = sql.NullString{String:"{{device_id}}"}
	it.ID = sql.NullInt64{Int64:0}
	it.ItemCardCharging = sql.NullFloat64{Float64:0}
	it.ItemDeleteStatus = sql.NullString{String:""}
	it.ItemDiscount = sql.NullFloat64{Float64:0}
	it.ItemDiscountShipping = sql.NullFloat64{Float64:0}
	it.ItemImage = sql.NullString{String:"https://dev-static.kudo.co.id/api/images/category/ic_pulsa_telkomsel.png"}
	it.ItemKomisi = sql.NullFloat64{Float64:0}
	it.ItemKudoFee = sql.NullFloat64{Float64:0}
	it.ItemName = sql.NullString{String:"Voucher Rp25.000 (081220172375)"}
	it.ItemReferenceID = sql.NullString{String:"TSV100"}
	it.ItemShipping = sql.NullFloat64{Float64:0}
	it.KudoboxID = sql.NullInt64{Int64:0}
	it.MerchantTrxID = sql.NullString{String:""}
	it.OrderID = sql.NullInt64{Int64:238442290, Valid:true}
	it.OrderStep = sql.NullInt64{Int64:0}
	it.PartnerPlu = sql.NullString{String:""}
	it.Price = sql.NullFloat64{Float64:79000}
	it.PurchaseCode = sql.NullString{String:""}
	it.QtyUpdateStatus = sql.NullString{String:""}
	it.Quantity = sql.NullInt64{Int64:1}
	it.RefundShippingCharges = sql.NullInt64{Int64:0}
	it.ResellerPrice = sql.NullFloat64{Float64:0}
	it.RetryProcess = sql.NullInt64{Int64:0}
	it.ShippingCode = sql.NullString{String:""}
	it.ShopItemID = sql.NullInt64{Int64:26178240}
	it.Sms = sql.NullInt64{Int64:0}
	it.SpotID = sql.NullInt64{Int64:190}
	it.Status = sql.NullInt64{Int64:1}
	it.VendorID = sql.NullInt64{Int64:83}
	it.UpdatedAt = mysql.NullTime{Time: time.Now()}
	t.Log("Testing ...")
	{
		typ := NewSqlTyping(InsertQuery)
		queries, err := typ.Typing(it)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}

		if len(queries) == 0 {
			t.Fatalf("%s expected queries not empty", failed)
		}

		if len(queries) > 0 {
			t.Logf("Queries %v", queries)
		}
	}

	t.Log("Testing type iterator")
	{
		output := bytes.NewBufferString("")
		err := TypeIterator(it, output)

		if err != nil {
			t.Fatalf("%s expected error nil, got %s", failed, err.Error())
		}

		if IsEmpty(output) {
			t.Fatalf("%s expected output not empty", failed)
		}

		if !IsEmpty(output) {
			t.Logf("Result data: %sv", output)
		}
	}
}