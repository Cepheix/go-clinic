package application

import "github.com/go-clinic/appointments/domain/repository"

type PatientController struct {
	patientRepository repository.PatientRepository
}

func NewPatientController(repository repository.PatientRepository) PatientController {
	return PatientController{patientRepository: repository}
}
