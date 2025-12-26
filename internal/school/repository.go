package school

import (
	"go_poc/pkg/generic"

	"gorm.io/gorm"
)

type SchoolRepository struct {
	*generic.BaseRepository[School]
}

type ClassroomRepository struct {
	*generic.BaseRepository[Classroom]
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	baseRepo := generic.NewBaseRepository[School](db)
	return &SchoolRepository{
		BaseRepository: baseRepo,
	}
}

func NewClassroomRepository(db *gorm.DB) *ClassroomRepository {
	baseRepo := generic.NewBaseRepository[Classroom](db)
	return &ClassroomRepository{
		BaseRepository: baseRepo,
	}
}
