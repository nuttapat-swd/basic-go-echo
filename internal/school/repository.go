package school

import (
	"go_poc/pkg/generic"

	"gorm.io/gorm"
)

type SchoolRepository struct {
	*generic.BaseRepository[School]
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	baseRepo := generic.NewBaseRepository[School](db)
	return &SchoolRepository{
		BaseRepository: baseRepo,
	}
}
