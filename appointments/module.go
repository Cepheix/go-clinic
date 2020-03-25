package appointments

import (
	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments/application"
)

func RegisterModule(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/patients", application.CreatePatient)
}
