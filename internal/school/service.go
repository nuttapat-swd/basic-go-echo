package school

import "go_poc/pkg/generic"

type SchoolService struct {
	*generic.BaseService[School]
}

func NewSchoolService(repo *SchoolRepository) *SchoolService {
	baseService := generic.NewService[School](repo.BaseRepository)
	return &SchoolService{
		BaseService: baseService,
	}
}
