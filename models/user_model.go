package models

import (
	"github.com/rudolfoborges/go-gin-gorm-blog-api/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	Posts     []Post
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashPassword)
	return
}

func (u User) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"email":      u.Email,
		"posts":      utils.ToSerializedArray(u.Posts),
	}
}
