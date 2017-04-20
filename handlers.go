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
