/*
Movie prints the Movie struct as JSON

$ go run src/ch4/movie/main.go
[
    {
        "Title": "Casablanca",
        "released": 1942,
        "Actors": [
            "Humphrey Bogart",
            "Ingrid Bergman"
        ]
    },
    {
        "Title": "Cool Hand Luke",
        "released": 1967,
        "color": true,
        "Actors": [
            "Paul Newman"
        ]
    },
    {
        "Title": "Bullit",
        "released": 1968,
        "color": true,
        "Actors": [
            "Steve McQueen",
            "Jacqueline Bisset"
        ]
    }
]

[{Casablanca} {Cool Hand Luke} {Bullit}]
*/

package main

import (
    "fmt"
    "log"
    "encoding/json"
    "os"
)

type Movie struct {
    Title  string
    Year   int      `json:"released"`
    Color  bool     `json:"color,omitempty"`
    Actors []string
}

func main() {
    var movies = []Movie{
        {Title: "Casablanca", Year: 1942, Color: false,
            Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
        {Title: "Cool Hand Luke", Year: 1967, Color: true,
            Actors: []string{"Paul Newman"}},
        {Title: "Bullit", Year: 1968, Color: true,
            Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
        }

    data, err := json.MarshalIndent(movies, "", "    ")
    if err != nil {
        log.Fatalf("JSON marshling failed: %s\n", err)
    }

    fmt.Printf("%s\n\n", data)

    var titles []struct{Title string}

    if err := json.Unmarshal(data, &titles); err != nil {
        log.Fatalf("JSON unmarshal failed: %s\n", err)
    }

    fmt.Println(titles)
}
