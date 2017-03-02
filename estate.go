package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

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
    fmt.Fprintln(w, "Estate Index!")
}

func EstateShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    estateId := vars["estateId"]
    fmt.Fprintln(w, "Estate show:", estateId)
}
