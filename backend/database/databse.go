package database

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {

	dsn := "User.db"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	if err := InitdefaultUser(db); err != nil {
		return nil, err
	}

	return db, nil
}

func InitdefaultUser(db *gorm.DB) error {
	var count int64

	err := db.Raw("SELECT COUNT(1) From [users]").Scan(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return db.Transaction(func(tx *gorm.DB) error {
			var userId int
			if err := tx.Raw(`INSERT INTO [users] (name,email,password,is_admin) VALUES('suresh','suresh@gmail.com','123qweasd',true)`).Scan(&userId).Error; err != nil {
				return err
			}
			return nil
		})
	}
	return err
}
