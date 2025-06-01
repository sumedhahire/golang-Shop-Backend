package dbConfig

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
	"log"
	"os"
)

type dbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Driver   string
}

func loadDBConfig() *dbConfig {
	return &dbConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
}

func (cfg *dbConfig) getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name)
}

func GetDB() *sql.DB {
	var err error
	config := loadDBConfig()

	db, err := sql.Open("mysql", config.getDSN())
	if err != nil {
		log.Fatalln("Database connection failed:", err)
	}

	runMigrations(db)

	return db
}

func runMigrations(db *sql.DB) {
	goose.SetDialect("mysql")

	if err := goose.Up(db, "./migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations ran successfully")
}
