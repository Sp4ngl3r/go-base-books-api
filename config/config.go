package config

import (
	"database/sql"
	"fmt"
	clog "log"
	"os"

	"github.com/unbxd/go-base/v2/log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	AppPort     string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	Environment string
	DB          *sql.DB
}

var AppConfig *Config
var AppLogger log.Logger

func LoadConfig() {
	_ = godotenv.Load()

	AppConfig = &Config{
		AppPort:     getEnv("APP_PORT", "4444"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "password"),
		DBName:      getEnv("DB_NAME", "postgres"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	AppConfig.initDB()
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func (c *Config) initDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		clog.Fatalf("Failed to connect to DB: %v", err)
	}

	c.DB = db
	clog.Println("✅ Database connection established")
}

func InitLogger() {
	var err error
	AppLogger, err = log.NewZeroLogger(
		log.ZeroLoggerWithLevel("debug"),
	)
	if err != nil {
		clog.Fatalf("error initializing logger: %v", err)
	}
}
