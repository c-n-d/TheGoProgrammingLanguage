/*
Issues prints a table fo GitHub issues matching the search terms.

$ go run src/ch4/issues/main.go is:open json decoder
4347 issues:
#59    adamliter Problem with JSON decoder and sperimentschema.json
#4           bep Support custom encoder/decoder
#1895  maxstepan JSON Decoder destroys existing heka message fields
#31       TechBK add encoder, decoder paramater
#389   GoogleCod Problems with 64-bit numbers in JSON decoder
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe
#347     chancez generator: Use a JSON decoder with UseNumber enabled fo
#389   GoogleCod Problems with 64-bit numbers in JSON decoder
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe
#389   GoogleCod Problems with 64-bit numbers in JSON decoder
#256   GoogleCod * Failed to use GoogleLocAPI: simplejson.decoder.JSONDe
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
        fmt.Printf("#%-5d %9.9s %55.55s\n",
            item.Number, item.User.Login, item.Title)
    }
}
