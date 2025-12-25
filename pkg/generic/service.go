package generic

type BaseService[T any] struct {
	repo *BaseRepository[T]
}

func NewService[T any](repo *BaseRepository[T]) *BaseService[T] {
	return &BaseService[T]{repo: repo}
}

func (s *BaseService[T]) Create(input *T) (*T, error) {
	if err := s.repo.Create(input); err != nil {
		return nil, err
	}
	return input, nil
}

func (s *BaseService[T]) Get(id uint) (*T, error) {
	inst, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func (s *BaseService[T]) List() ([]T, error) {
	return s.repo.List()
}

func (s *BaseService[T]) Update(id uint, input *T) (*T, error) {
	if _, err := s.repo.Get(id); err != nil {
		return nil, err
	}
	if err := s.repo.Update(input); err != nil {
		return nil, err
	}
	return input, nil
}

func (s *BaseService[T]) Delete(id uint) error {
	return s.repo.Delete(id)
}
