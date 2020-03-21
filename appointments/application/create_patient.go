package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments/domain"
	"github.com/go-clinic/appointments/persistance"
)

func CreatePatient(context *gin.Context) {
	var command CreatePatientCommand

	if err := context.ShouldBindJSON(&command); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repository := persistance.NewPatientRepository()
	handler := NewCreatePatientCommandHandler(repository)

	handler.CreateNewPatient(command)

	context.JSON(http.StatusCreated, gin.H{})
}

type CreatePatientCommand struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type CreatePatientCommandHandler struct {
	patientRepository domain.PatientRepository
}

func NewCreatePatientCommandHandler(repository domain.PatientRepository) CreatePatientCommandHandler {
	return CreatePatientCommandHandler{patientRepository: repository}
}

func (handler *CreatePatientCommandHandler) CreateNewPatient(command CreatePatientCommand) {
	patient := domain.Patient{FirstName: command.FirstName, LastName: command.LastName}
	handler.patientRepository.Add(patient)
}
