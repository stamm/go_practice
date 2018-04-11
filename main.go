//Ну а теперь задание.
//Нужно будет написать хендлер, который принимает на вход json вида {"name": "Victor"} и возвращает {"response": "Hello, Victor"}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Name     string `json:"name"`
	Response string `json:"response"`
}

// curl localhost:8000 -d '{"name":"Hello"}'
func VictorHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	m := map[string]string{}
	m["response"] = "Hello, " + msg.Name

	output, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func main() {
	http.HandleFunc("/", VictorHandler)
	address := ":8000"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
