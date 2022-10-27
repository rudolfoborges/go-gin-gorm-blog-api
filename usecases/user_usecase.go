package usecases

import (
	"errors"

	"github.com/rudolfoborges/go-gin-gorm-blog-api/database"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/models"
)

func CreateUserUseCase(user models.User) (uint, error) {
	var existingUser models.User

	if result := database.DB.Find(&existingUser, "email = ?", user.Email); result.RowsAffected > 0 {
		return 0, errors.New("email already exists")
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}
