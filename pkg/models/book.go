package models

import (
	"go-mysql/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ID     		uint   `json:"id"`
	Title  		string `json:"title"`
	Author 		string `json:"author"`
	Year   		string `json:"year"`
	ISBN		string `json:"isbn"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}