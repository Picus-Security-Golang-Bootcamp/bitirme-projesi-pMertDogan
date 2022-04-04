package user

import (
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/crypto"
	"gorm.io/gorm"
)

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

//check if there is any user thats registered with email
func (c *UserRepository) CheckIsUserExistWithThisEmail(email string) (*User, error) {
	var user User
	result := c.db.Where("email = ?", email).First(&user)

	// if result.RowsAffected == 0 {
	// 	//zero means not found
	// 	return false, nil, nil
	// }
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	} else if result.Error != nil {
		return  nil, result.Error
	}
	//return true and password (Hash)
	return  &user, nil
}

//register user with his email and password(hash)
func (c *UserRepository) RegisterUser(email string, password string) error {
	//hash user password with bcrypt
	//https://godoc.org/golang.org/x/crypto/bcrypt

	passwordHashed, err := customCrypto.HashPassword(password)

	//if there is an error on hashing password
	if err != nil {
		return err
	}
	//save user to database
	c.db.Create(&User{Email: email, Password: passwordHashed})
	return nil
}
