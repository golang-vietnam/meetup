package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 400)
			return
		}

		var data Payload
		if err := proto.Unmarshal(buf, &data); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 400)
			return
		}

		pl, err := proto.Marshal(&data)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 400)
			return
		}

		w.Header().Set("Content-Type", "application/protobuf")
		w.Write(pl)
	})

	http.ListenAndServe(":9902", router)
}
