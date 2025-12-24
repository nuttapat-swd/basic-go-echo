package school

import (
	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	ShortName string `gorm:"not null" json:"short_name"`
	Address   string `json:"address"`
}

type Classroom struct {
	gorm.Model
	Grade    string `gorm:"not null" json:"grade"`
	Section  string `gorm:"not null" json:"section"`
	SchoolID uint
	School   School
}
