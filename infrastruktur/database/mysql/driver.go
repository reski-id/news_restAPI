package mysql

import (
	"fmt"
	"log"
	"portal/config"

	contentData "portal/feature/content/data"
	detailData "portal/feature/detail/data"
	userData "portal/feature/user/data"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(userData.User{}, contentData.Content{}, detailData.Detail{})
}

func SeedUsers(db *gorm.DB) {
	users := []userData.User{
		{
			Username: "jhon",
			Email:    "jhon@gmail.com",
			Password: "123456",
			FullName: "Jhon Wick",
			Role:     "user",
		},
		{
			Username: "jhane",
			Email:    "jhane@gmail.com",
			Password: "jhane123",
			FullName: "Jhane Do",
			Role:     "user",
		},
		{
			Username: "reski",
			Email:    "programmer.reski@gmail.com",
			Password: "reski1234",
			FullName: "Ahmad Reski",
			Role:     "admin",
		},
	}

	for _, user := range users {
		var existingUser userData.User
		result := db.Where("email = ?", user.Email).First(&existingUser)

		if result.Error != nil {
			// If there is no error, it means the user already exists in the database
			if result.RowsAffected > 0 {
				continue
			}

			// If the user doesn't exist, hash the password and insert the data
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal("Error hashing password: ", err)
			}
			user.Password = string(hashedPassword)

			result := db.Create(&user)
			if result.Error != nil {
				log.Fatal("Error seeding user: ", result.Error)
			}
		}
	}
}

func SeedContent(db *gorm.DB) {
	contentList := []contentData.Content{
		{
			Title:   "New iPhone 14 leaks",
			Content: "Rumors suggest that Apple is working on a foldable iPhone that could launch in 2023.",
			Images:  "https://example.com/iphone14.jpg",
			UserID:  1,
		},
		{
			Title:   "Google announces new Pixel 6",
			Content: "The new Pixel 6 will feature Google's custom Tensor chip and an improved camera system.",
			Images:  "https://example.com/pixel6.jpg",
			UserID:  1,
		},
		{
			Title:   "Amazon's new Echo Show 15",
			Content: "The Echo Show 15 features a massive 15-inch display and improved voice recognition.",
			Images:  "https://example.com/echoshow15.jpg",
			UserID:  1,
		},
		{
			Title:   "Facebook announces Ray-Ban smart glasses",
			Content: "Facebook has partnered with Ray-Ban to create smart glasses that can take photos and make calls.",
			Images:  "https://example.com/rayban.jpg",
			UserID:  1,
		},
		{
			Title:   "Microsoft announces Surface Laptop 4",
			Content: "The Surface Laptop 4 features upgraded internals and improved battery life.",
			Images:  "https://example.com/surface.jpg",
			UserID:  1,
		},
	}

	for _, content := range contentList {
		result := db.Create(&content)

		if result.Error != nil {
			log.Fatal("Error seeding content: ", result.Error)
		}
	}
}
