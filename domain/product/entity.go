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
	Sku         string `json:"sku" gorm:"unique" `
	ProductName string `json:"productName" `
	Description string `json:"description" `
	Color       string `json:"color"`
	// https://stackoverflow.com/questions/9452897/how-to-decode-json-with-type-convert-from-string-to-float64
	Price      int `json:"price,string" `
	StockCount int `json:"stockCount,string"`
	//https://gorm.io/docs/belongs_to.html
	//This one is our foreign key
	CategoryID int `json:"categoryId,string"`
	//BTW we dont need hardcode reference:ID cause looks like its default
	Category category.Category `json:"category"`        //`gorm:"foreignKey:CategoryID;references:ID"`
	StoreID  int               `json:"storeId,string" ` //This one is our foreign key for store
	Store    store.Store       `json:"store"`           //`gorm:"foreignKey:StoreID;references:ID"`
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

func FromReqCreateDTO(reqProduct ReqCreateDTO) Product {
	return Product{
		ProductName: reqProduct.ProductName,
		Description: reqProduct.Description,
		Color:       reqProduct.Color,
		Price:       reqProduct.Price,
		CategoryID:  reqProduct.CategoryID,
		StockCount:  reqProduct.StockCount,
		StoreID:     reqProduct.StoreID,
		Sku:         reqProduct.Sku,
	}
}
