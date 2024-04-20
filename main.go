package main

import (
	"api/infrastructure/db"
	"api/infrastructure/db/query"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Docker Compose with Go")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		conn := db.New()
		q := query.Use(conn.DB)
		userQ := q.User

		us, err := userQ.Where(userQ.Username.Eq("tony")).Find()
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		for _, u := range us {
			fmt.Fprintf(w, "%v\n", u)
		}
	})

	http.ListenAndServe(":1323", nil)
}
