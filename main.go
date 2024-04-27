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

		fmt.Fprintf(w, "%v\n", res)
	})

	http.HandleFunc("/users/credit_cards", func(w http.ResponseWriter, r *http.Request) {
		post := func() {
			ctx := r.Context()

			repository := repository.NewUser()
			u, err := repository.GetByName("tony")
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}

			err = repository.AddCreditCard(ctx, u.ID, "1234567890")
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}

			fmt.Fprintf(w, "Credit card added")
		}

		switch r.Method {
		case http.MethodPost:
			post()
			return
		case http.MethodGet:
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed) // 405 Method Not Allowed
			fmt.Fprintf(w, "Invalid request method")
			return
		}
	})

	http.ListenAndServe(":1323", nil)
}
