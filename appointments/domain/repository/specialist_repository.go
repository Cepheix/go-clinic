package repository

import "github.com/go-clinic/appointments/domain/model"

type SpecialistRepository interface {
	Add(specialist model.Specialist) (int, error)
}
