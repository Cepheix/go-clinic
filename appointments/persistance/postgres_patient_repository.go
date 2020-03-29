package persistance

import (
	"github.com/go-clinic/appointments/domain/model"
	"github.com/go-clinic/appointments/domain/repository"
	"github.com/jinzhu/gorm"
)

type PostgresPatientRepository struct {
	db *gorm.DB
}

func (repository PostgresPatientRepository) Add(patient model.Patient) (int, error) {
	repository.db.Create(&patient)
	repository.db.Save(&patient)

	return patient.ID, repository.db.Error
}

func NewPatientRepository(db *gorm.DB) repository.PatientRepository {
	return PostgresPatientRepository{db: db}
}
