/*
Exercise 8.12 - Make the broadcaster announce the current set of clients to each new arrival

$ go build ch8/ex_8_12$
go build ch8/netcat3
$ ex_8_12 &

$ ./netcat3
You are 127.0.0.1:58013
Online: 127.0.0.1:58013
127.0.0.1:58019 has arrived
127.0.0.1:58019: Hi!
Howdy!
127.0.0.1:58013: Howdy!
127.0.0.1:58024 has arrived
127.0.0.1:58024: Any one here?
^C

$ netcat3
You are 127.0.0.1:58019
Online: 127.0.0.1:58013, 127.0.0.1:58019
Hi!
127.0.0.1:58019: Hi!
127.0.0.1:58013: Howdy!
127.0.0.1:58024 has arrived
127.0.0.1:58024: Any one here?
127.0.0.1:58013 has left
^C

$ netcat3
You are 127.0.0.1:58024
Online: 127.0.0.1:58013, 127.0.0.1:58019, 127.0.0.1:58024
Any one here?
127.0.0.1:58024: Any one here?
127.0.0.1:58013 has left
127.0.0.1:58019 has left
^C

$ killall ex_8_12
*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "strings"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }

    go broadcaster()
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go handleConn(conn)
    }
}

type client struct {
    Name    string
    Mailbox chan<- string // an outgoing message channel
}

var (
    entering = make(chan client)
    leaving  = make(chan client)
    messages = make(chan string) // all incoming client messages
)

func broadcaster() {
    clients := make(map[string]client) // all connected clients
    for {
        select {
            case msg := <-messages:
                // Broadcast incoming messages to all clients' outgoing message channels
                for _, cli := range clients {
                    cli.Mailbox <- msg
                }
            case cli := <- entering:
                clients[cli.Name] = cli
                var current []string
                for _, cli := range clients {
                    current = append(current, cli.Name)
                }
                cli.Mailbox <- "Online: " + strings.Join(current, ", ")
            case cli := <- leaving:
                delete(clients, cli.Name)
                close(cli.Mailbox)
        }
    }
}

func handleConn(conn net.Conn) {
    ch := make(chan string) // outgoing clients messages
    go clientWriter(conn, ch)

    who := conn.RemoteAddr().String()
    me := client{ who, ch }

    ch <- "You are " + who
    messages <- who + " has arrived"
    entering <- me

    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages <- who + ": " + input.Text()
    }
    // Note: ignoring potential errors from input.Err()

    leaving <- me
    messages <- who + " has left"
    conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg) // Note: ignoring network errors
    }
}
