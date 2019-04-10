package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	closeGPIO := setUpGPIO()
	defer func() {
		log.Println("close GPIO")
		closeGPIO()
	}()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(map[string]string{"status": "up"})
		if err != nil {
			panic(err)
		}
		w.Write(res)
	})

	http.Handle("/ws", webSocketHandler())

	log.Fatal(http.ListenAndServe(":80", nil))
}
