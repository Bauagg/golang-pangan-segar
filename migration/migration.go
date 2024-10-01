package migration

import (
	"pangan-segar/databases"
	modelglobal "pangan-segar/model/global"
)

func Migration() {
	db := databases.GetDB()
	err := db.AutoMigrate(
		modelglobal.Users{},
		modelglobal.Otps{},
		modelglobal.Address{},
	)

	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}
}
