package backend

import (
	"ecst-demo/cassis163/article-service/backend/models"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitOrm() {
	db := createDatabaseConnection()
	log.Info("InitOrm")
	db.AutoMigrate(&models.Article{})
	article := models.Article{Name: "D42", Price: 100}
	db.Create(&article)
}

func createDatabaseConnection() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=article port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
