package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "../../.env")

	if err != nil {
		log.Fatal("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv("dsn")
		db, err = gorm.Open(mysql.Open(dsn))
	} else {
		dsn = os.Getenv("dsnTest")
		db, err = gorm.Open(mysql.Open(dsn))
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// if os.Getenv("AutoMigrateDb") == "true" {
	// 	db.AutoMigrate(&domain.User{}, &domain.Author{}, &domain.Challenge{}, &domain.ChallengeFile{})
	// 	db.Model(domain.ChallengeFile{}).AddForeignKey("challenge_id", "challenges (id)", "CASCADE", "CASCADE")
	// }

	return db
}
