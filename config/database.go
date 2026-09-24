package config

import (
	"fmt"
	"log"
	"os"

	"evermos-backend/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	_ = godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "evermos_vix"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}

	log.Println("Database connection established successfully!")

	// Auto-migration
	err = DB.AutoMigrate(
		&model.User{},
		&model.Toko{},
		&model.Category{},
		&model.Alamat{},
		&model.Produk{},
		&model.FotoProduk{},
		&model.Trx{},
		&model.LogProduk{},
		&model.DetailTrx{},
	)
	if err != nil {
		log.Fatalf("Auto-migration failed: %v", err)
	}

	log.Println("Database migration completed!")
	return DB
}
