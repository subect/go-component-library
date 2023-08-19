package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func NewDbDsn(dbName, dbUser, dbPwd, dbAddr string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=30s&charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=Local&parseTime=true", dbUser, dbPwd, dbAddr, dbName)
}

func NewClient(dsn string, devMode bool) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("db: %+v err:%v", db, err)
		return nil, err
	}

	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 5)

	// 启用Logger，显示详细日志
	db.LogMode(devMode)
	return db, nil
}
