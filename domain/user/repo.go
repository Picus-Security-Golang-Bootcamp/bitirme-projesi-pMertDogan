package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *UserRepository = nil

//initilaze the repo with gorm db
func UserRepoInit(db *gorm.DB) *UserRepository {

	if singleton == nil {
		singleton = &UserRepository{db}
	}
	return singleton
}

//Before using this you need initialize the repo with UserRepoInit
func Repo() *UserRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *UserRepository) Migrations() {
	c.db.AutoMigrate(&User{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}