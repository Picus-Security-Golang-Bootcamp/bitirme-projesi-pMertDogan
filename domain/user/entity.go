package user

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Users []User

type User struct {
	gorm.Model
	UserName string 
	Email string 	`gorm:"unique"` //make sure email is unique
	Password string
	IsAdmin  bool 
	// FalseLoginCount int
	// ExpiresAt string
}

// fromJson users
func UnmarshalBooks(data []byte) (Users, error) {
	var r Users
	err := json.Unmarshal(data, &r)
	return r, err
}


//Users toJson
func (r *Users) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//User toJson
func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
