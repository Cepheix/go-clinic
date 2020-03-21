package domain

type PatientRepository interface {
	Add(patient Patient)
}
