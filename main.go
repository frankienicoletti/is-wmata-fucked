package main

import (
	"encoding/json"
	"fmt"
	"is-wmata-fucked/wmata"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		railIncidents, err := wmata.GetRailIncidents()
		if err != nil {
			fmt.Printf("error getting rail incidents: %v", err)
			res.WriteHeader(500)
		}

		busIncidents, err := wmata.GetBusIncidents()
		if err != nil {
			fmt.Printf("error getting bus incidents: %v", err)
			res.WriteHeader(500)
		}

		incidents := struct {
			Rail []wmata.RailIncindent `json:"rail,omitempty"`
			Bus  []wmata.BusIncindent  `json:"bus,omitempty"`
		}{
			Rail: railIncidents,
			Bus:  busIncidents,
		}
		b, err := json.Marshal(incidents)
		if err != nil {
			fmt.Printf("error marshaling incidents: %v", err)
			res.WriteHeader(500)
		}
		res.Write(b)
	})
	log.Fatal(http.ListenAndServe(":8000", r))
}
