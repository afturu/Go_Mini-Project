package config

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "tukerin-platform/entities"
)

var DB *gorm.DB

func InitDB() {
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Format DSN (Data Source Name) MySQL
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPass, dbHost, dbPort, dbName,
    )

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Gagal terhubung ke database: %v", err)
    }

    fmt.Println("Terhubung ke database!")

    migrateDB()
}

func migrateDB() {
    err := DB.AutoMigrate(
        &entities.User{},
        &entities.UserProfile{},
        &entities.Item{},
        &entities.Category{},
        &entities.Transaction{},
    )
    if err != nil {
        log.Fatalf("Gagal melakukan migrasi database: %v", err)
    }
    fmt.Println("Migrasi database berhasil!")
}