package migration

import (
	"20_Clean-Architecture/model"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}