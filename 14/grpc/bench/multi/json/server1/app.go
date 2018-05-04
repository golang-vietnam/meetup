package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type payload struct {
	A string `json:"a"`
	B string `json:"b"`
	C int    `json:"c"`
	D string `json:"d"`
	E bool   `json:"e"`
	F int64  `json:"f"`
	G string `json:"g"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		var data payload
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 400)
			return
		}

		pl, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(pl)
	})

	http.ListenAndServe(":9901", router)
}
