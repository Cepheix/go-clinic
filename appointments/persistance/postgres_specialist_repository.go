package persistance

import (
	"github.com/go-clinic/appointments/domain/model"
	"github.com/go-clinic/appointments/domain/repository"
	"github.com/jinzhu/gorm"
)

type PostgresSpecialistRepository struct {
	db *gorm.DB
}

func (repository PostgresSpecialistRepository) Add(specialist model.Specialist) (int, error) {
	repository.db.Create(&specialist)
	repository.db.Save(&specialist)

	return specialist.ID, repository.db.Error
}

func (repository PostgresSpecialistRepository) Save(specialist model.Specialist) error {
	repository.db.Save(&specialist)
	return repository.db.Error
}

func (repository PostgresSpecialistRepository) Find(id int) (*model.Specialist, error) {
	specialist := &model.Specialist{}
	err := repository.db.First(specialist, id).Error

	return specialist, err
}

func NewSpecialistRepository(db *gorm.DB) repository.SpecialistRepository {
	return PostgresSpecialistRepository{db: db}
}
