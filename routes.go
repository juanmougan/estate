package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
    Route {
        "Index",
        "GET",
        "/",
        Index,
    },
    Route {
        "EstateIndex",
        "GET",
        "/estates",
        EstateIndex,
    },
    Route {
        "EstateShow",
        "GET",
        "/estates/{estateId}",
        EstateShow,
    },
    Route {
        "EstateCreate",
        "POST",
        "/estates",
        EstateCreate,
    },
}
