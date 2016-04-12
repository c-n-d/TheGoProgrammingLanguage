/*
Exercise 4.10 - Modifies 'issues' to add a 'days since creation' column to the issues reported

$ go run src/ch4/ex_4_10/main.go is:open json decoder
4350 issues:
#59    adamliter      Problem with JSON decoder and sperimentschema.json     8 days ago
#4           bep                          Support custom encoder/decoder     0 days ago
#1895  maxstepan      JSON Decoder destroys existing heka message fields    14 days ago
#31       TechBK                          add encoder, decoder paramater    43 days ago
#389   GoogleCod            Problems with 64-bit numbers in JSON decoder    14 days ago
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe     0 days ago
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe     1 days ago
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe     1 days ago
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe     2 days ago
#347     chancez generator: Use a JSON decoder with UseNumber enabled fo    41 days ago
#389   GoogleCod            Problems with 64-bit numbers in JSON decoder    23 days ago
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe     2 days ago
#389   GoogleCod            Problems with 64-bit numbers in JSON decoder    23 days ago
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe     4 days ago
...
*/

package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "ch4/github"
)

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%d issues:\n", result.TotalCount)

    for _, item := range result.Items {
        fmt.Printf("#%-5d %9.9s %55.55s %-14s\n",
            item.Number, item.User.Login, item.Title, daysSince(item.CreateAt))
    }
}

func daysSince(t time.Time) string {
    return fmt.Sprintf("%5.0f days ago", (time.Since(t).Hours() / 24.0))
}
