package main

import (
	"os"

	"github.com/gitkuldp/wedding-invitation-api/internal/db"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"

	"github.com/sirupsen/logrus"
)

type AppEnv string

func (e AppEnv) String() string {
	return string(e)
}

const (
	// DevEnv is an environment for development release
	DevEnv AppEnv = "dev"
	// add other environments below if needed
	// eg: ProdEnv AppEnv = "prod"
)

func main() {
	env := db.NewEnv()
	db := db.InitDB(env)

	tx := db.Begin()

	// for dev release, we will drop all tables and migrate again,
	// this is because gorm does not support alter tables properly,
	// and dropping tables helps to resolve conflicting issues
	//if strings.ToLower(env.AppEnv) == DevEnv.String() {
	err := tx.Migrator().DropTable(models.AllModels...)
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		os.Exit(1)
	}
	//}

	err = tx.AutoMigrate(models.AllModels...)
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		os.Exit(1)
	}
	tx.Commit()
}
