package persistance

import (
	"github.com/go-clinic/appointments/domain"
	"github.com/jinzhu/gorm"
)

type PostgresSpecialistRepository struct {
	db *gorm.DB
}

func (repository PostgresSpecialistRepository) Add(specialist domain.Specialist) (int, error) {
	repository.db.Create(&specialist)
	repository.db.Save(&specialist)

	return specialist.ID, repository.db.Error
}

func NewSpecialistRepository(db *gorm.DB) domain.SpecialistRepository {
	return PostgresSpecialistRepository{db: db}
}
