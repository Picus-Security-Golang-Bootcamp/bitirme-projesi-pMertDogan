package order


import (
	"encoding/json"
	"gorm.io/gorm"
)

type Orders []Order

type Order struct {
	gorm.Model
	UserID int
	ProductID int
	Comment string
	ShippingAdress string
	BillingAddress string
	quantity int
}

// fromJson Order
func UnmarshalOrders(data []byte) (Orders, error) {
	var r Orders
	err := json.Unmarshal(data, &r)
	return r, err
}


//Order toJson
func (r *Orders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Order toJson
func (r *Order) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
