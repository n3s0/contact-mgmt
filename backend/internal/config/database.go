package config

import (
	"fmt"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host string
	Port int
	User string
	Password string
	DBName string
	SSLMode string
}

func NewDatabase(conf DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.DBName,
		conf.SSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database %w", err)
	}

	dbConnection, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get database instance: %w", err)
	}

	dbConnection.SetMaxIdleConns(10)
	dbConnection.SetMaxOpenConns(10)
	dbConnection.SetConnMaxLifetime(time.Hour)
	dbConnection.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}
