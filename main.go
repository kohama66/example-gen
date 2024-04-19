package main

import (
	"api/domain/entity"
	"api/infrastructure/db"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Docker Compose with Go")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		conn := db.New()
		us := []*entity.User{}

		// query
		conn.Table("users").Where("email = 'tony@example.com'").Find(&us)

		for _, u := range us {
			fmt.Fprintf(w, "%v\n", u)
		}
	})

	http.ListenAndServe(":1323", nil)
}
