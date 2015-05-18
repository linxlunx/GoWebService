package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"database/sql"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GetUserList(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("select rowid, username, fullname from users")
		switch {
			case err == sql.ErrNoRows:
				resp := NotFound{
					Status: 404, 
					Message: "Not Found",
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				error := json.NewEncoder(w).Encode(resp)
				if error != nil {
					log.Fatal(error)
				}
			case err != nil:
				log.Fatal(err)
			default:
				Users := []*User{}
				for rows.Next() {
					var rowid string
					var username string
					var fullname string

					rows.Scan(&rowid, &username, &fullname)
					finalid, _ := strconv.Atoi(rowid)
					Users = append(Users, &User{Rowid: finalid, Username: username, Fullname: fullname})
				}


				resp := ResponseList{
					Status: 200, 
					Message: "User List",
					Result: Users,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				error := json.NewEncoder(w).Encode(resp)
				if error != nil {
					log.Fatal(error)
				}
		}	
	}
}

func GetUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["Id"]

		var rowid string
		var username string
		var fullname string

		err := db.QueryRow("select rowid, username, fullname from users where rowid = ?", id).Scan(&rowid, &username, &fullname)
		switch {
			case err == sql.ErrNoRows:
				resp := NotFound{
					Status: 404, 
					Message: "Not Found",
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				error := json.NewEncoder(w).Encode(resp)
				if error != nil {
					log.Fatal(error)
				}
			case err != nil:
				log.Fatal(err)
			default:
				finalid, _ := strconv.Atoi(rowid)
				resp := ResponseDetail{
					Status: 200, 
					Message: "User detail",
					Result: User{
						Rowid: finalid,
						Username: username,
						Fullname: fullname,
					},
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				error := json.NewEncoder(w).Encode(resp)
				if error != nil {
					log.Fatal(error)
				}
		}	
	}
}
