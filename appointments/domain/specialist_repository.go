package domain

type SpecialistRepository interface {
	Add(specialist Specialist) (int, error)
}