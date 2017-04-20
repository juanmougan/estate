package main

import "time"

type Estate struct {
    Id      int       `json:"id"`
    Name		string		`json:"name"`
    Lat 		float64		`json:"lat"`
    Long 		float64		`json:"long"`
    Created 	time.Time	`json:"created"`
}

type Estates []Estate
