package main

import (
	"database/sql"
	"fmt"
	"log"

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
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=5433 sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

//Run will start the App reference on the port passed in.
func (a *App) Run(addr string) {

}
