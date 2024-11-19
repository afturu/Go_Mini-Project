package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "tukerin-platform/entities"
)

var DB *gorm.DB

// LoadEnv memuat file .env
func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("Peringatan: Tidak dapat memuat file .env, menggunakan environment default.")
    }
}

// InitDB menginisialisasi koneksi database
func InitDB() {
    LoadEnv()

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
        log.Fatalf("Environment database tidak lengkap! Periksa file .env Anda.")
    }

    // Format DSN (Data Source Name) untuk MySQL
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPass, dbHost, dbPort, dbName,
    )

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Gagal terhubung ke database: %v", err)
    }

    log.Println("Terhubung ke database!")
    migrateDB()
}

// migrateDB melakukan migrasi entitas ke database
func migrateDB() {
    if DB == nil {
        log.Fatalf("Database belum diinisialisasi. Pastikan InitDB berhasil dijalankan.")
    }

    // Daftar entitas untuk migrasi
    entitiesToMigrate := []interface{}{
        &entities.User{},
        &entities.UserProfile{},
        &entities.Item{},
        &entities.Category{},
        &entities.Transaction{},
    }

    // Proses migrasi
    if err := DB.AutoMigrate(entitiesToMigrate...); err != nil {
        log.Fatalf("Gagal melakukan migrasi database: %v", err)
    }
    log.Println("Migrasi database berhasil!")
}