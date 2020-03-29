package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments/domain"
)

// CreatePatient godoc
// @Summary Create a new patient
// @Description add patient by json
// @Tags patients
// @Accept  json
// @Produce  json
// @Param patient body CreatePatientCommand true "patient"
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

	id, err := handler.CreateNewPatient(command)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"id": id})
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

func (handler *CreatePatientCommandHandler) CreateNewPatient(command CreatePatientCommand) (int, error) {
	patient := domain.Patient{FirstName: command.FirstName, LastName: command.LastName}
	return handler.patientRepository.Add(patient)
}
