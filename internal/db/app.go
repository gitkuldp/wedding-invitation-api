package db

import "gorm.io/gorm"

type Application struct {
	Env *Env
	Db  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Db = InitDB(app.Env)
	return *app
}
