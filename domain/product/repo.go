package product

import (
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"gorm.io/gorm"
)

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

//return all products with relations
func (c *ProductRepository) GetAllWithPagination(page, pageSize string) (Products, error) {

	var products Products
	// product := Product
	//resturn paginated data
	result := c.db.Scopes(domain.Paginate(page, pageSize)).Joins("Store").Joins("Category").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil

}

//return all products with relations
func (c *ProductRepository) SearchProducts( searchText string) (Products, error) {


	//https://www.compose.com/articles/mastering-postgresql-tools-full-text-search-and-phrase-search/
	var products Products
	// product := Product
	//resturn paginated data
	// result := c.db.Scopes(domain.Paginate("1", "10")).Joins("Store").Joins("Category").Find(&products)
	result := c.db.
	Where("product_name LIKE ?", "%"+searchText+"%").
	Or("products.description LIKE ?", "%"+searchText+"%").
	Or("color LIKE ?", "%"+searchText+"%").
	Or("sku LIKE ?", "%"+searchText+"%").
	//hardcoded store name search :/
	Or(" \"Store\".\"name\" LIKE ?", "%"+searchText+"%").
	Or("\"Category\".\"category_name\" LIKE ?", "%"+searchText+"%").
	Joins("Store").Joins("Category").Find(&products).Limit(10)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil

}

//create product 
func (c *ProductRepository) CreateBulkProduct(products Products) {
	//categoryName is uniq

	for _, v := range products {
		//https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
		//If its not exist just create it else update it
		//SKU is uniq
		if c.db.Model(&v).Where("sku = ?", v.Sku).Updates(&v).RowsAffected == 0 {
			c.db.Create(&v)
			//zero means not found
		}
	}

}