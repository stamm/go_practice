/*
* Вопросы:
* Чем отличается fmt.Println от println
* Что ты имел ввиду "Но нужно просто считать не из http.Request. Тут есть примерчик" из переписки
* Я немного запутался в структуре папок - почему у меня тут package main
* Чем отличаются ковычки "" '' ``
 */
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var results []string

// GetHandler - Get request handler
func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

func main() {
	results = append(results, time.Now().Format(time.RFC3339))

	http.HandleFunc("/calculate", GetHandler)

	log.Printf("listening on port %s", ":8000")
	log.Printf(`Use: curl localhost:8000/calculate -d '{"name":"Hello"}`)
	http.ListenAndServe(":8000", nil)
}
