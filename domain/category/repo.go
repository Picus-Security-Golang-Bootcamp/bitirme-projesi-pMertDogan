package category

import (
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type CategoryRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *CategoryRepository = nil

//initilaze the repo with gorm db
func CategoryRepoInit(db *gorm.DB) *CategoryRepository {
	if singleton == nil {
		singleton = &CategoryRepository{db}
	}
	return singleton
}

//Before using this you need initialize the repo
func Repo() *CategoryRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *CategoryRepository) Migrations() {
	c.db.AutoMigrate(&Category{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

//Save All Categories to SQL
//https://gorm.io/docs/create.html#Batch-Insert
func (c *CategoryRepository) CreateCategories(categories Categorys) {
	//categoryName is uniq
	// c.db.Create(&categories)

	// c.db.Clauses(clause.OnConflict{
	// 	UpdateAll: true,
	//   }).Create(&categories)
	// c.db.FirstOrCreate(&categories)

	for _, v := range categories {
		//https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
		//If its not exist just create it else update it
		if c.db.Model(&v).Where("category_name = ?", v.CategoryName).Updates(&v).RowsAffected == 0 {
			c.db.Create(&v)
			//zero means not found
		}
	}

}

//Get All Categories from DB
func (c *CategoryRepository) GetAllCategories() (Categorys, error) {
	var categories Categorys
	var result *gorm.DB
	result = c.db.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil

}

func (c *CategoryRepository) GetAllCategoriesWithLimit(limit int) (Categorys, error) {
	var categories Categorys
	var result *gorm.DB
	result = c.db.Find(&categories).Limit(limit)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil

}

//https://gorm.io/docs/advanced_query.html#Locking-FOR-UPDATE
func update() {}
