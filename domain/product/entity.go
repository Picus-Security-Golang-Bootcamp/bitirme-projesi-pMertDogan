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
	Sku         int    `json:"sku" binding:"required,alphanum"`
	ProductName string `json:"productName" binding:"required,min=3,alphanum"`
	Description string `json:"description" binding:"alphanum"`
	Color       string `json:"color"`
	Price       int    `json:"price" binding:"required,numeric"`
	StockCount  int    `json:"stockCount" binding:"required,numeric"`
	//https://gorm.io/docs/belongs_to.html
	//This one is our foreign key
	CategoryID int `json:"categoryId"`
	//BTW we dont need hardcode reference:ID cause looks like its default
	Category category.Category `json:"category"` //`gorm:"foreignKey:CategoryID;references:ID"`
	StoreID  int               `json:"storeId"`  //This one is our foreign key for store
	Store    store.Store       `json:"store"`    //`gorm:"foreignKey:StoreID;references:ID"`
	//Owner + ID -> OwnerID , Example Store + ID -> StoreID
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
