package structs

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	id  uint64
	Name  string
	Url string
}

func(Menu) TableName() string {
	return "tbl_menu"
}