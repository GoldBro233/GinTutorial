package config

import (
	"GinWithGormTutorial/global"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func buildDSN(user, password, host, port, db, charset string) string {
	template := "%s:%s@tcp(%s:%s)/%s?charset=%s"
	return fmt.Sprintf(template, user, password, host, port, db, charset)
}

func initDB() {
	dsn := buildDSN(AppConfig.Database.User, AppConfig.Database.Password, AppConfig.Database.Host, AppConfig.Database.Port, AppConfig.Database.Name, AppConfig.Database.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database,%v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	if err != nil {
		log.Fatalf("Failed to configure to database,%v", err)
	}

	global.Db = db
}
