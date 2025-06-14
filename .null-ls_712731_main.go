package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"message": "hello world!",
		}
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	})
	router.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"message": "hi world!",
		}
		err := json.NewEncoder(w).Encode(data)
		log.Fatal(err)
	})
	fmt.Println("Server started at port 8414")
	log.Fatal(http.ListenAndServe(":8414", router))
}
