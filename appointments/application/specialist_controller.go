package application

import (
	"github.com/go-clinic/appointments/domain/repository"
)

type SpecialistController struct {
	specialistRepository repository.SpecialistRepository
}

func NewSpecialistController(specialistRepository repository.SpecialistRepository) SpecialistController {
	return SpecialistController{specialistRepository: specialistRepository}
}
