package persistance

import (
	"github.com/go-clinic/appointments/domain/model"
	"github.com/jinzhu/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&model.Patient{})
	db.AutoMigrate(&model.Specialist{})
	db.AutoMigrate(&model.SpecialistAbility{})
}
