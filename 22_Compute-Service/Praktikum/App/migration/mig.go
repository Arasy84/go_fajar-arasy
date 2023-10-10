package migration

import (
	"Docker/model"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	// auto migrate untuk table book
}