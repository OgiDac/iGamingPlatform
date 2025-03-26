package config

import "database/sql"

type Application struct {
	MySql *sql.DB
}

func App() Application {
	app := &Application{}
	app.MySql = NewDbConnection()
	return *app
}

func (app *Application) CloseDatabaseConnection() {
	CloseDbConnection(app.MySql)
}
