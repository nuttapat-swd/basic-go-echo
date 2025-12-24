package school

type SchoolService struct {
	repo *SchoolRepository
}

func NewSchoolService(repo *SchoolRepository) *SchoolService {
	return &SchoolService{repo: repo}
}

func (s *SchoolService) CreateSchool(name, short_name, address string) (*School, error) {
	school := &School{Name: name, ShortName: short_name, Address: address}
	if err := s.repo.Create(school); err != nil {
		return nil, err
	}
	return school, nil
}

func (s *SchoolService) ListSchools() ([]School, error) {
	return s.repo.List()
}

func (s *SchoolService) GetSchool(id uint) (*School, error) {
	school, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (s *SchoolService) UpdateSchool(id uint, name, short_name, address string) (*School, error) {
	school, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	school.Name = name
	school.ShortName = short_name
	school.Address = address
	if err := s.repo.Update(school); err != nil {
		return nil, err
	}
	return school, nil
}

func (s *SchoolService) DeleteSchool(id uint) error {
	return s.repo.Delete(id)
}
