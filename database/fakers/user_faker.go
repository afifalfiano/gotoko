package fakers

import (
	"time"

	"github.com/afifalfiano/gotoko/app/models"
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:               uuid.New().String(),
		Addresses:        nil,
		FirstName:        faker.FirstName(),
		LastName:         faker.LastName(),
		Email:            faker.Email(),
		Password:         "$2y$10$HimlYI49kuuJbpBkEIQ9Ge4ajwPMXspjgKOlrMRgYzVaDBrF3r.Cy", //password
		RememberPassword: "",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}
