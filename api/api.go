/*
* Вопросы:
* Чем отличается fmt.Println от println
* Что ты имел ввиду "Но нужно просто считать не из http.Request. Тут есть примерчик" из переписки
* Я немного запутался в структуре папок - почему у меня тут package main
* Чем отличаются ковычки "" '' ``
* что такое mux
 */
// package main

// import (
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// var results []byte

// // GetHandler - Get request handler
// func GetHandler(w http.ResponseWriter, r *http.Request) {
// 	resp, err := http.Get("/")
// 	if err != nil {
// 		// handle error
// 	}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	defer resp.Body.Close()
// 	if err != nil {
// 		http.Error(w, "Error converting results to json",
// 			http.StatusInternalServerError)
// 	}
// 	// w.Write(jsonBody)
// 	println(string(body))
// }

// func main() {
// 	results = append(results, time.Now().Format(time.RFC3339)...)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/calcualte", GetHandler)

// 	log.Printf("listening on port %s", ":8000")
// 	http.ListenAndServe(":8000", mux)
// }

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// DefaultHandler - test
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	println("test")
}

func main() {
	http.HandleFunc("/", DefaultHandler)

	errServer := http.ListenAndServe(":8000", nil)
	if errServer != nil {
		panic(errServer)
	}

	res, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
	println("asdasd")
}
