package school

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name      string    `gorm:"not null" json:"name"`
	ShortName string    `gorm:"not null" json:"short_name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type Classroom struct {
	gorm.Model
	Grade     string `gorm:"not null" json:"grade"`
	Section   string `gorm:"not null" json:"section"`
	SchoolID  uint
	School    School
	CreatedAt time.Time `json:"created_at"`
}
