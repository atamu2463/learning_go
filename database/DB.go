package database

import (
	"log"

	"github.com/atamu2463/learning_go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	//SQLiteに接続
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB接続に失敗しました", err)
	}

	log.Println("DB接続に成功しました")
	return db
}

func MigrateDB(db *gorm.DB) {
	//マイグレーションを実行
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("マイグレーションに失敗しました", err)
	}

	log.Println("マイグレーションに成功しました")
}

func InsertSeedData(db *gorm.DB) {
	//シードデータを挿入
	sampleUser := models.User{
		ID:       1,
		Name:     "sampleuser",
		Password: "samplepassword",
	}

	result := db.Create(&sampleUser)
	if result.Error != nil {
		log.Fatal("シードデータの挿入に失敗しました", result.Error)
	}

	log.Println("シードデータの挿入に成功しました")
}
