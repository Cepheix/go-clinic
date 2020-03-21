package persistance

import (
	"fmt"

	"github.com/go-clinic/appointments/domain"
)

type PostgresPatientRepository struct {
}

func (repository PostgresPatientRepository) Add(patient domain.Patient) {
	fmt.Printf("Storing %+v", patient)
}

func NewPatientRepository() domain.PatientRepository {
	return PostgresPatientRepository{}
}
