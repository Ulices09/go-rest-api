package db

import (
	"go-rest-api/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Post{})

	// defer func() {
	// 	sqlDB, err := db.DB()

	// 	if err != nil {
	// 		log.Fatalf("FATAL: Fail on close DB Connection, %s\n", err)
	// 	}

	// 	err = sqlDB.Close()

	// 	if err != nil {
	// 		log.Fatalf("FATAL: Fail on close DB Connection, %s\n", err)
	// 	}
	// }()

	return db
}
