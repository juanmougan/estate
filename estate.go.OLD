package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
	"encoding/json"

    "github.com/gorilla/mux"
)

type Estate struct {
    Name		string		`json:"name"`
    Lat 		float64		`json:"lat"`
    Long 		float64		`json:"long"`
    Created 	time.Time	`json:"created"`
}

type Estates []Estate

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/estates", EstateIndex)
    router.HandleFunc("/estates/{estateId}", EstateShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func EstateIndex(w http.ResponseWriter, r *http.Request) {
    estates := Estates {
        Estate { Name: "Departamento en Palermo", Lat: -34.594142, Long: -58.422036 },
        Estate { Name: "Casa en San Isidro", Lat: -34.467610, Long: -58.510191 },
    }

    json.NewEncoder(w).Encode(estates)
}

func EstateShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    estateId := vars["estateId"]
    fmt.Fprintln(w, "Estate show:", estateId)
}
