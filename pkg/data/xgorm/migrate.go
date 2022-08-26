package xgorm

import (
	"fmt"
)

func InitMigrateCmd(config *Config, do string, tables ...interface{}) {
	migrateORM, cleanup, err := OpenMigrateDB(config)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	var orm *XGorm
	if do == "migrate" {
		tmpORM, cleanup2, err := OpenDB(config)
		if err != nil {
			panic(err)
		}
		orm = tmpORM
		defer cleanup2()
	}

	switch do {
	case "create":
		migrateORM.CreateDB()
	case "drop":
		migrateORM.DropDB()
	case "migrate":
		if err := orm.AutoMigrate(tables...); err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("no support do:%s", do))
	}
}
