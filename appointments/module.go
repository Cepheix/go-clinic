package appointments

import (
	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments/application"
	"github.com/go-clinic/appointments/persistance"
	"github.com/jinzhu/gorm"
)

func RegisterModule(routerGroup *gin.RouterGroup, db *gorm.DB) {

	patientRepository := persistance.NewPatientRepository(db)
	patientController := application.NewPatientController(patientRepository)

	routerGroup.POST("/patients", patientController.CreatePatient)

	specialistRepository := persistance.NewSpecialistRepository(db)
	specialistController := application.NewSpecialistController(specialistRepository)

	routerGroup.POST("/specialists", specialistController.CreateSpecialist)
	routerGroup.POST("/specialists/abilities", specialistController.AddSpeciality)
}
