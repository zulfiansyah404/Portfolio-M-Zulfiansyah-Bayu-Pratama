package connection

import (
	"project/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func DatabaseConnect() {
	// Menyambungkan ke localhost:5432
	database, err := gorm.Open(postgres.Open("postgresql://postgres:yesnmahsheh@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Project{}, &models.User{})

	DB = database
}