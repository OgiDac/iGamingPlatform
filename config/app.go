package config

import "github.com/jmoiron/sqlx"

type Application struct {
	Env   *Env
	MySql *sqlx.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.MySql = NewDbConnection(app.Env)
	return *app
}

func (app *Application) CloseDatabaseConnection() {
	CloseDbConnection(app.MySql)
}
