/*
Autoescape demonstrates automatic HTML escaping in html/template.

$ go run src/ch4/autoescape/main.go > src/ch4/autoescape/index.html
$ open src/ch4/autoescape/index.html
*/

package main

import (
    "html/template"
    "log"
    "os"
)

func main() {
    const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
    t := template.Must(template.New("escape").Parse(templ))

    var data struct {
        A string        // untrusted plain text
        B template.HTML // trusted HTML
    }

    data.A = "<b>Hello!</b>"
    data.B = "<b>Hello!</b>"

    if err := t.Execute(os.Stdout, data); err != nil {
        log.Fatal(err)
    }
}
