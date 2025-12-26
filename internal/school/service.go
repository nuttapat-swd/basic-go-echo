package school

import (
	"go_poc/pkg/generic"
)

type SchoolService struct {
	*generic.BaseService[School]
}

func NewSchoolService(repo *SchoolRepository) *SchoolService {
	baseService := generic.NewService(repo.BaseRepository)
	return &SchoolService{
		BaseService: baseService,
	}
}

type ClassroomService struct {
	*generic.BaseService[Classroom]
	schoolRepo *SchoolRepository
}

func NewClassroomService(repo *ClassroomRepository, schoolRepo *SchoolRepository) *ClassroomService {
	baseService := generic.NewService(repo.BaseRepository)
	return &ClassroomService{
		BaseService: baseService,
		schoolRepo:  schoolRepo,
	}
}

func (s *ClassroomService) Create(entity *Classroom) (*Classroom, error) {

	_, err := s.schoolRepo.Get(entity.SchoolID)
	if err != nil {
		return nil, err
	}

	createdClassroom, err := s.BaseService.Create(entity)
	if err != nil {
		return nil, err
	}
	return s.Get(createdClassroom.ID)
}

func (s *ClassroomService) Get(id uint) (*Classroom, error) {
	classroom, err := s.BaseService.Get(id)
	if err != nil {
		return nil, err
	}

	school, err := s.schoolRepo.Get(classroom.SchoolID)
	if err != nil {
		return nil, err
	}

	classroom.School = *school
	return classroom, nil
}

func (s *ClassroomService) List() ([]Classroom, error) {
	classrooms, err := s.BaseService.List()
	if err != nil {
		return nil, err
	}

	schools, err := s.schoolRepo.List()
	if err != nil {
		return nil, err
	}

	schoolByID := make(map[uint]School, len(schools))
	for _, school := range schools {
		schoolByID[school.ID] = school
	}

	for i, entity := range classrooms {
		school, ok := schoolByID[entity.SchoolID]
		if !ok {
			return nil, ErrSchoolInvalid
		}
		classrooms[i].School = school
	}
	return classrooms, nil
}

func (s *ClassroomService) Update(id uint, data map[string]any) (*Classroom, error) {
	classroom, err := s.BaseService.Update(id, data)
	if err != nil {
		return nil, ErrNotFound
	}
	school, err := s.schoolRepo.Get(classroom.SchoolID)
	if err != nil {
		return nil, ErrSchoolInvalid
	}
	classroom.School = *school

	return classroom, nil
}
