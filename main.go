package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type alt struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Combo string `json:"combo"`
}

func main() {

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	r.HandleFunc("/gen", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		var (
			email string
			password string
			)

		if err := db.QueryRow("SELECT email, password FROM alts ORDER BY rand() LIMIT 1").Scan(&email, &password); err != nil {
			log.Fatal(err)
		}

		altGenerated := alt{
			Email:    email,
			Password: password,
			Combo: email + ":" + password,
		}

		json.NewEncoder(w).Encode(altGenerated)

	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
