package database

import (
	"diaxel/internal/config"
	. "diaxel/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Connect(settings *config.Settings) {
	host := settings.DbHost
	user := settings.DbUser
	password := settings.DbPassword
	dbname := settings.DbName
	port := settings.DbPort
	sslmode := settings.Ssl

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	//err = db.Migrator().DropTable(&Chat{}, &Message{})
	//if err != nil {
	//	log.Fatalf("Не удалось удалить таблицы: %v", err)
	//}
	//log.Println("Таблицы успешно удалены")

	err = db.AutoMigrate(&Chat{}, &Message{})
	if err != nil {
		log.Fatalf("Не удалось создать таблицы: %v", err)
	}

	log.Println("Таблицы успешно созданы")
}

func Disconnect() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close the database connection:", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
