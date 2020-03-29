package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments/domain/model"
	"github.com/go-clinic/appointments/domain/repository"
)

// CreateSpecialist godoc
// @Summary Create a new specialist
// @Description add new specialist by json
// @Tags specialist
// @Accept  json
// @Produce  json
// @Param specialist body CreateSpecialistCommand true "specialist"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /appointments/specialists [post]
func (controller SpecialistController) CreateSpecialist(context *gin.Context) {
	var command CreateSpecialistCommand

	if err := context.ShouldBindJSON(&command); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler := NewCreateSpecialistCommandHandler(controller.specialistRepository)

	id, err := handler.CreateNewPatient(command)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"id": id})
}

type CreateSpecialistCommand struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type CreateSpecialistCommandHandler struct {
	specialistRepository repository.SpecialistRepository
}

func NewCreateSpecialistCommandHandler(repository repository.SpecialistRepository) CreateSpecialistCommandHandler {
	return CreateSpecialistCommandHandler{specialistRepository: repository}
}

func (handler CreateSpecialistCommandHandler) CreateNewPatient(command CreateSpecialistCommand) (int, error) {
	patient := model.Specialist{FirstName: command.FirstName, LastName: command.LastName}
	return handler.specialistRepository.Add(patient)
}
