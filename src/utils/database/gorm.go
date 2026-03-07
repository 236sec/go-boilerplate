package database

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

type GormDB struct {
	DB *gorm.DB
}

// Create implements the Database interface
func (db *GormDB) Create(value any) error {
	return db.DB.Create(value).Error
}

// Find implements the Database interface
func (db *GormDB) Find(dest any, conds ...any) error {
	return db.DB.Find(dest, conds...).Error
}

// First implements the Database interface
func (db *GormDB) First(dest any, conds ...any) error {
	return db.DB.First(dest, conds...).Error
}

// Where implements the Database interface
func (db *GormDB) Where(query any, args ...any) Database {
	return &GormDB{DB: db.DB.Where(query, args...)}
}

var (
	dbInstance *GormDB
	dbOnce     sync.Once
)

// initDatabase initializes the database connection
func initDatabase() *GormDB {
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

	return &GormDB{DB: db}
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
func (db *GormDB) GetGormDB() *gorm.DB {
	return db.DB
}