package migration

import (
	smsmodule "airbnb-messaging-be/internal/app/sms"
	orm "airbnb-messaging-be/internal/pkg/gorm"
	"airbnb-messaging-be/internal/pkg/log"

	"gorm.io/gorm"
)

func MigrateUp(db gorm.DB) {
	models := []interface{}{
		&smsmodule.Sms{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal(orm.Instance, "failed to run migration", err)
	}
}
