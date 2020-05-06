package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/point", pointHandler)
	http.HandleFunc("/api/blabla", blablaHandler)
	//http.HandleFunc("/api/...", ...Handler)

	http.ListenAndServe(":8080", nil)
}

// The `json:"whatever"` bit is a way to tell the JSON
// encoder and decoder to use those names instead of the
// capitalised names
type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

var _p *point = &point{
	X: 3,
	Y: 28,
}

func pointHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// Just send out the JSON version of 'tom'
		j, _ := json.Marshal(_p)
		w.Header().Set("Content-Type", "application/json")

		w.Write(j)
	case "POST":
		// Decode the JSON in the body and overwrite 'tom' with it
		d := json.NewDecoder(r.Body)
		p := &point{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_p = p

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func blablaHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// Just send out the JSON version of 'tom'
		j, _ := json.Marshal(point)
		w.Header().Set("Content-Type", "application/json")

		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}
