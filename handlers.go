package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func EstateIndex(w http.ResponseWriter, r *http.Request) {
    estates := Estates {
        Estate { Name: "Departamento en Palermo", Lat: -34.594142, Long: -58.422036 },
        Estate { Name: "Casa en San Isidro", Lat: -34.467610, Long: -58.510191 },
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
    if err := json.NewEncoder(w).Encode(estates); err != nil {
        panic(err)
    }
}

func EstateShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    estateId := vars["estateId"]
    fmt.Fprintln(w, "Estate show:", estateId)
}
