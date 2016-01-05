package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"../helpers/"
	"github.com/gorilla/mux"
)

func GetUserList(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		type User struct {
			Id       int    `json:"id"`
			Username string `json:"username"`
			Fullname string `json:"fullname"`
		}

		rows, err := db.Query("select rowid, username, fullname from users")
		switch {
		case err == sql.ErrNoRows:
			resp := helpers.ShowResponse(404, "", map[string]interface{}{})
			helpers.WriteResponse(resp, w, r)
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
				Users = append(Users, &User{
					Id:       finalid,
					Username: username,
					Fullname: fullname,
				})
			}
			result := map[string]interface{}{
				"users": Users,
			}
			resp := helpers.ShowResponse(200, "User Lists", result)
			helpers.WriteResponse(resp, w, r)
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
			resp := helpers.ShowResponse(404, "", map[string]interface{}{})
			helpers.WriteResponse(resp, w, r)
		case err != nil:
			log.Fatal(err)
		default:
			finalid, _ := strconv.Atoi(rowid)
			result := map[string]interface{}{
				"id":       finalid,
				"username": username,
				"fullname": fullname,
			}
			resp := helpers.ShowResponse(200, "User Detail", result)
			helpers.WriteResponse(resp, w, r)
		}

	}
}
