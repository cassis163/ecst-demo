package backend

import (
	"ecst-demo/cassis163/article-service/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitOrm() {
	db := createDatabaseConnection()
	db.AutoMigrate(&models.Article{})
}

func createDatabaseConnection() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
