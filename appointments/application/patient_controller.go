package application

import "github.com/go-clinic/appointments/domain"

type PatientController struct {
	patientRepository domain.PatientRepository
}

func NewPatientController(repository domain.PatientRepository) PatientController {
	return PatientController{patientRepository: repository}
}
