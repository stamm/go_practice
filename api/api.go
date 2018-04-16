/*
* Вопросы:
* Чем отличается fmt.Println от println
* Что ты имел ввиду "Но нужно просто считать не из http.Request. Тут есть примерчик" из переписки
* Я немного запутался в структуре папок - почему у меня тут package main
* Чем отличаются ковычки "" '' ``
* что такое mux
* как генерить айдишники для объектов структур
* 41 - зачем ставить запятую, если перечисления нет
* json.NewDecoder что это?
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ID - id of the ride
var ID int

// GOOGLEAPIKEY is GOOGLEAPIKEY
const GOOGLEAPIKEY string = "AIzaSyAEb0MBFh70ZZSVW13pRW0mA2ljc9oXSY4"

func main() {
	http.HandleFunc("/", GetHandler)
	http.ListenAndServe(":8000", nil)
}

// RideData - collecting data of the current trip
type RideData struct {
	ID        int    `json:"id"`
	StartLat  string `json:"start_lat"`
	EndLat    string `json:"end_lat"`
	StartLong string `json:"start_long"`
	EndLong   string `json:"end_long"`
}

//RideValues - calculate values
type RideValues struct {
	ID       int     `json:"id"`
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
}

type Response struct {
	Rows []Rows `json:"rows"`
}

type Rows struct {
	Elements []Elements `json:"elements"`
}

type Elements struct {
	Distance Distance `json:"distance"`
	Duration Duration `json:"duration"`
}

type Distance struct {
	Value float64 `json:"value"`
}

type Duration struct {
	Value float64 `json:"value"`
}

var rides []RideData

// GetHandler - handle GET params
func GetHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	if len(keys) < 1 {
		log.Println("No params passed")
		return
	}

	ride := RideData{
		ID:        ID,
		StartLat:  r.URL.Query().Get("start_lat"),
		EndLat:    r.URL.Query().Get("end_lat"),
		StartLong: r.URL.Query().Get("start_long"),
		EndLong:   r.URL.Query().Get("end_long"),
	}

	rideJSON, _ := json.Marshal(ride)

	rides = append(rides, ride)

	ID++
	fmt.Println("LOG: " + string(rideJSON))
	GetDistanceAndDuration(ride.StartLat, ride.EndLat, ride.StartLong, ride.EndLong)
}

//GetDistanceAndDuration - get distance and duration from google api
func GetDistanceAndDuration(slat string, endlat string, slng string, endlng string) {
	url := "https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=" + slat + "," + slng + "&destinations=" + endlat + "," + endlng + "&key=" + GOOGLEAPIKEY
	rs, err := http.Get(url)
	// Process response
	if err != nil {
		println("test1")
		panic(err)
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var resp Response
	json.Unmarshal(bodyBytes, &resp)

	rw := RideValues{
		ID:       rides[len(rides)-1].ID,
		Distance: resp.Rows[0].Elements[0].Distance.Value,
		Duration: resp.Rows[0].Elements[0].Duration.Value,
	}

	data, err := json.Marshal(rw)
	fmt.Println(string(data))
}
