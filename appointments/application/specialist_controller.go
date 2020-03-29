package application

import "github.com/go-clinic/appointments/domain"

type SpecialistController struct {
	specialistRepository domain.SpecialistRepository
}

func NewSpecialistController(specialistRepository domain.SpecialistRepository) SpecialistController {
	return SpecialistController{specialistRepository: specialistRepository}
}
