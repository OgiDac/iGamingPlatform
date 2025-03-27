package config

import "github.com/jmoiron/sqlx"

type Application struct {
	MySql *sqlx.DB
}

func App() Application {
	app := &Application{}
	app.MySql = NewDbConnection()
	return *app
}

func (app *Application) CloseDatabaseConnection() {
	CloseDbConnection(app.MySql)
}
