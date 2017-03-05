package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//App is the structure to hold our Application variables
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize will assign the database values to the App variable referenced.
func (a *App) Initialize(user, password, dbname string) {

}

//Run will start the App reference on the port passed in.
func (a *App) Run(addr string) {

}
