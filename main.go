package main

import (
	"api/infrastructure/db"
	"api/infrastructure/repository"
	"fmt"
	"net/http"
)

func main() {
	db.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Docker Compose with Go")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		repository := repository.NewUser()

		res, err := repository.GetByName("tony")
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		// for _, u := range res {
		// 	fmt.Fprintf(w, "%v\n", u)
		// }
		fmt.Fprintf(w, "%v\n", res)
	})

	http.ListenAndServe(":1323", nil)
}
