package product

import (
	"encoding/json"

	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/store"
	"gorm.io/gorm"
)

type Products []Product

type Product struct {
	gorm.Model
	Sku         int
	ProductName        string
	Description string
	Color       string
	Price       int
	StockCount  int
	//https://gorm.io/docs/belongs_to.html
	CategoryID int //This one is our foreign key
	//BTW we dont need hardcode reference:ID cause looks like its default
	Category category.Category `gorm:"foreignKey:CategoryID;references:ID"`
	StoreID  int               //This one is our foreign key for store
	Store    store.Store       `gorm:"foreignKey:StoreID;references:ID"`
}

// fromJson Product
func UnmarshalStore(data []byte) (Products, error) {
	var r Products
	err := json.Unmarshal(data, &r)
	return r, err
}

//Products toJson
func (r *Products) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Product toJson
func (r *Product) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
