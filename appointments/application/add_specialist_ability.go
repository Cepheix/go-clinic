package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments/domain/model"
	"github.com/go-clinic/appointments/domain/repository"
)

// AddAbility godoc
// @Summary Add ability for given specialist
// @Description add ability by json
// @Tags specialist
// @Accept  json
// @Produce  json
// @Param specialistAbility body AddSpecialitAbilityCommand true "Specialist Ability"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /appointments/specialists/abilities [post]
func (controller SpecialistController) AddSpeciality(context *gin.Context) {
	var command AddSpecialitAbilityCommand

	if err := context.ShouldBindJSON(&command); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler := NewAddSpecialistAbilityHandler(controller.specialistRepository)

	err := handler.AddNewAbility(command)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{})
}

type AddSpecialitAbilityCommand struct {
	SpecialistId int
	AbilityType  model.Procedure
}

type AddSpecialitAbilityCommandHandler struct {
	specialistRepository repository.SpecialistRepository
}

func NewAddSpecialistAbilityHandler(repository repository.SpecialistRepository) AddSpecialitAbilityCommandHandler {
	return AddSpecialitAbilityCommandHandler{specialistRepository: repository}
}

func (handler AddSpecialitAbilityCommandHandler) AddNewAbility(command AddSpecialitAbilityCommand) error {
	specialist, err := handler.specialistRepository.Find(command.SpecialistId)

	if err != nil {
		return err
	}

	specialist.AddAbility(command.AbilityType)

	return handler.specialistRepository.Save(*specialist)
}
