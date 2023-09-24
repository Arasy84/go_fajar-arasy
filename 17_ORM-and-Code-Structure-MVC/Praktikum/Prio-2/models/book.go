package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Tittle    string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}
