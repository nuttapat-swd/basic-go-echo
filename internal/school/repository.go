package school

import "gorm.io/gorm"

type SchoolRepository struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	return &SchoolRepository{db: db}
}

func (r *SchoolRepository) Create(school *School) error {
	return r.db.Create(school).Error
}

func (r *SchoolRepository) List() ([]School, error) {
	var schools []School
	err := r.db.Find(&schools).Error
	return schools, err
}

func (r *SchoolRepository) Get(id uint) (*School, error) {
	var school School
	if err := r.db.First(&school, id).Error; err != nil {
		return nil, err
	}
	return &school, nil
}

func (r *SchoolRepository) Update(school *School) error {
	return r.db.Save(school).Error
}

func (r *SchoolRepository) Delete(id uint) error {
	result := r.db.Delete(&School{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
