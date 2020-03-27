package persistance

import (
	"fmt"

	"github.com/go-clinic/appointments/domain"
	"github.com/jinzhu/gorm"
)

type PostgresPatientRepository struct {
	db gorm.DB
}

func (repository PostgresPatientRepository) Add(patient domain.Patient) {
	fmt.Printf("Storing %+v", patient)
}

func NewPatientRepository(db *gorm.DB) domain.PatientRepository {
	return PostgresPatientRepository{}
}
