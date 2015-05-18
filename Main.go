package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", httpLogger(Index))
	router.HandleFunc("/users", httpLogger(GetUserList(db)))
	router.HandleFunc("/user/{Id}", httpLogger(GetUser(db)))

	log.Fatal(http.ListenAndServe(":9090", router))
}