package persistance

import (
	"github.com/go-clinic/appointments/domain"
	"github.com/jinzhu/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&domain.Patient{})
	db.AutoMigrate(&domain.Specialist{})
}
