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
func (c *ProductRepository) SearchProducts(searchText, page, pageSize string) (Products, error) {

	//https://www.compose.com/articles/mastering-postgresql-tools-full-text-search-and-phrase-search/
	var products Products

	//ILIKE is case insensitive
	result := c.db.
		Where("product_name ILIKE ?", "%"+searchText+"%").
		Or("products.description ILIKE ?", "%"+searchText+"%").
		Or("color ILIKE ?", "%"+searchText+"%").
		Or("sku ILIKE ?", "%"+searchText+"%").
		//hardcoded store name search :/
		Or(" \"Store\".\"name\" ILIKE ?", "%"+searchText+"%").
		Or("\"Category\".\"category_name\" ILIKE ?", "%"+searchText+"%").
		Scopes(domain.Paginate(page, pageSize)).
		Joins("Store").Joins("Category").Find(&products).Limit(10)

	// result := c.db.Raw("select * from products  	Join categories ON categories.id = products.category_id Join stores ON stores.id = products.store_id where sku ILIKE ? 	or product_name ILIKE ? 	or products.description ILIKE ?", "%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%").Scan(&products)	

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
		//with the help of the unsoped we can detect soft deleted products if we cant detect its try to create it 
		if c.db.Model(&v).Unscoped().Where("sku = ?", v.Sku).Updates(&v).RowsAffected == 0 {
			//zero means not found
			c.db.Create(&v)
		}
	}

}

//delete product by id
func (c *ProductRepository) Delete(id string) (Product, error) {
	var product Product
	result := c.db.Where("id = ?", id).First(&product)

	if result.Error != nil {
		return product, result.Error
	}

	result = c.db.Delete(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

//PATCH Product
func (c *ProductRepository) Update(id string, patched Product) (*Product, error) {

	var old Product
	//get the first product if exist
	result := c.db.Where("id = ?", id).First(&old)

	if result.Error != nil {
		return nil, result.Error
	}
	//patch the produc

	c.db.Model(&old).Updates(patched).Where("id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &old, nil
}

//Get single product by id
func (c *ProductRepository) GetById(id string) (Product, error) {
	var product Product
	result := c.db.Where("id = ?", id).First(&product)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}
