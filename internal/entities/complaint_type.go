package entities

import "gorm.io/gorm"

type ComplaintType struct {
	*gorm.Model
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"varchar(255)"`
	Description string `gorm:"text"`
}
