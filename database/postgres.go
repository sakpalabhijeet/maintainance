package database

import (
	"Maintainance/internal/config"
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config)*gorm.DB{
	dsn:= fmt.Sprintf(
		"host=%s user =%s password =%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Kolkata",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
		cfg.DB.SSLMode,
	)
	db, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err!= nil{
		log.Fatalf("failed to get sqldb: %v", err)
	}
	sqlDb, err:= db.DB()
	if err!= nil{
		log.Fatalf("Failed to get sqldb: %v", err)
	}
	
	sqlDb.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	sqlDb.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(cfg.DB.MaxIdleTime)
// Health check
	if err:= sqlDb.Ping();err!= nil{
		log.Fatalf("Database Ping Failed! %v", err)
	}
	log.Println("Postgres Sql Connected Successfully!")
	return db
}