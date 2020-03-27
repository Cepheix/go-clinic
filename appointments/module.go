package appointments

import (
	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments/application"
	"github.com/go-clinic/appointments/persistance"
	"github.com/jinzhu/gorm"
)

func RegisterModule(routerGroup *gin.RouterGroup, db *gorm.DB) {

	repository := persistance.NewPatientRepository(db)
	patientController := application.NewPatientController(repository)

	routerGroup.POST("/patients", patientController.CreatePatient)
}
