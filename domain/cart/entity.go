package cart


import (
	"encoding/json"
	"gorm.io/gorm"
)

type Carts []Cart

type Cart struct {
	gorm.Model
	UserID int
	ProductID int
	Total int
}

// fromJson Order
func UnmarshalOrders(data []byte) (Carts, error) {
	var r Carts
	err := json.Unmarshal(data, &r)
	return r, err
}


//Order toJson
func (r *Carts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Order toJson
func (r *Cart) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
