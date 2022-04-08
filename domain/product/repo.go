package product

import "gorm.io/gorm"

type ProductRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *ProductRepository = nil

//initilaze the repo with gorm db
func ProductRepoInit(db *gorm.DB) *ProductRepository {
	if singleton == nil {
		singleton = &ProductRepository{db}
	}
	return singleton
}

//Before using this you need initialize the repo
func Repo() *ProductRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *ProductRepository) Migrations() {
	c.db.AutoMigrate(&Product{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}


//Create single product
func (c *ProductRepository) Create(product Product) error {
	result := c.db.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Get single product by sku with relations
func (c *ProductRepository) GetBySkuWithRelations(sku string) (Product, error) {
	var product Product
	//get product by sku with relations
	result := c.db.Joins("Store").Joins("Category").Where("sku = ?", sku).First(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}