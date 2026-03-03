package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database wraps gorm.DB and implements repository interfaces
type Database struct {
	*gorm.DB
}

type GormDB interface {
	Create(value interface{}) error
	Where(query interface{}, args ...interface{}) GormDB
	First(dest interface{}, conds ...interface{}) error
}

// Create implements the GormDB interface
func (db *Database) Create(value interface{}) error {
	return db.DB.Create(value).Error
}

// Where implements the GormDB interface  
func (db *Database) Where(query interface{}, args ...interface{}) GormDB {
	return &Database{DB: db.DB.Where(query, args...)}
}

// First implements the GormDB interface
func (db *Database) First(dest interface{}, conds ...interface{}) error {
	return db.DB.First(dest, conds...).Error
}

var (
	dbInstance *Database
	dbOnce     sync.Once
)

// GetDatabase returns a singleton database connection
func GetDatabase() *Database {
	dbOnce.Do(func() {
		dbInstance = initDatabase()
	})
	return dbInstance
}

// initDatabase initializes the database connection
func initDatabase() *Database {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Get database configuration from environment
	config := getDatabaseConfig()

	// Build PostgreSQL DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.Host, config.User, config.Password, config.DBName, config.Port)

	// Configure GORM with custom logger
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Database connection established successfully")

	return &Database{DB: db}
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

// getDatabaseConfig reads database configuration from environment variables
func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		User:     getEnvOrDefault("DB_USER", "postgres"),
		Password: getEnvOrDefault("DB_PASS", "mypassword"),
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		Port:     getEnvOrDefault("DB_PORT", "5432"),
		DBName:   getEnvOrDefault("DB_NAME", "boilerplate"),
	}
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetGormDB returns the underlying gorm.DB instance for migrations or advanced operations
func (db *Database) GetGormDB() *gorm.DB {
	return db.DB
}