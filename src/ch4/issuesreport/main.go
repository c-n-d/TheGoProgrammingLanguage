/*
IssuesReport prints a formated report of issues matching the search terms

$ go run src/ch4/issuesreport/main.go is:open json decoder
4357 issues:
---------------------------------------------
Number: 59
User:   adamliter
Title:  Problem with JSON decoder and sperimentschema.json
Age:    31383118769142 days
---------------------------------------------
Number: 1895
User:   maxstepanov
Title:  JSON Decoder destroys existing heka message fields
Age:    51134035437322 days
---------------------------------------------
Number: 31
User:   TechBK
Title:  add encoder, decoder paramater
Age:    158338743771685 days
---------------------------------------------
...
*/

package main

import (
    "html/template"
    "log"
    "os"
    "time"

    "ch4/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreateAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
    Funcs(template.FuncMap{"daysAgo": daysAgo}).
    Parse(templ))

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }

    if err := report.Execute(os.Stdout, result); err != nil {
        log.Fatal(err)
    }
}

func daysAgo(t time.Time) int {
    return int(time.Since(t) / 24)
}