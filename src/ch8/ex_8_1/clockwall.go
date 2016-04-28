/*
Exercise 8.1 - Modify clock2 to accept a port. Write `clockwall` that acts as a client of multiple
               clock servers and displays the result in a table.

$ go build -o clockwall src/ch8/ex_8_1/clockwall.go
$ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
NewYork     London      Tokyo       
16:56:46    21:56:46    05:56:46    
...
NewYork     London      Tokyo       
16:57:02    21:57:02    05:57:02    
*/

package main

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "strings"
    "time"
)

type TimeData struct {
    TZ, CurrentTime string
    Address         string
    Time           *bytes.Buffer
    Connection      net.Conn
}

func NewTimeData(s, sep string) *TimeData {
    var buf bytes.Buffer
    pair := strings.Split(s, sep)
    timezone := pair[0]
    address := pair[1]

    conn, err := net.Dial("tcp", address)
    if err != nil {
        log.Fatal(err)
    }

    return &TimeData{TZ: timezone, Address: address, Time: &buf, Connection: conn}
}

func (td *TimeData) Update() {
    td.Time.Reset()
    if _, err := io.CopyN(td.Time, td.Connection, 9); err != nil {
        log.Fatal(err)
    }
    td.CurrentTime = strings.TrimSpace(td.Time.String())
}

func main() {
    var allTimeData []*TimeData

    for _, arg := range os.Args[1:] {
        timeData := NewTimeData(arg, "=")
        allTimeData = append(allTimeData, timeData)

        go mustUpdate(timeData)
    }

    printHeader(allTimeData)

    for {
        printCurrentTime(allTimeData)
        time.Sleep(1 * time.Second)
    }
}

func mustUpdate(td *TimeData) {
    defer td.Connection.Close()
    for {
        td.Update()
        time.Sleep(500 * time.Millisecond)
    }
}

func printHeader(times []*TimeData) {
    for _, t := range times {
        fmt.Printf("%-12s", t.TZ)
    }
    fmt.Println()
}

func printCurrentTime(times []*TimeData) {
    fmt.Printf("\r")
    for _, t := range times {
        fmt.Printf("%-12s", t.CurrentTime)
    }
}
