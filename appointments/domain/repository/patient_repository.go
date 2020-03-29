package repository

import "github.com/go-clinic/appointments/domain/model"

type PatientRepository interface {
	Add(patient model.Patient) (int, error)
}
