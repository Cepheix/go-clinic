package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments/domain"
)

// CreatePatient godoc
// @Summary Create a new patient
// @Description add patient by json
// @Accept  json
// @Produce  json
// @Param patient body CreatePatientCommand true "Add patient"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /appointments/patients [post]
func (patientController PatientController) CreatePatient(context *gin.Context) {
	var command CreatePatientCommand

	if err := context.ShouldBindJSON(&command); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler := NewCreatePatientCommandHandler(patientController.patientRepository)

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
