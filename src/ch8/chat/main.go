/*
Chat is a server that lets clients chat with each other

$ go build ch8/chat
$ go build ch8/netcat3

$ ./netcat3 
You are 127.0.0.1:63969
127.0.0.1:63973 has arrived
Hello!
127.0.0.1:63969: Hello!
127.0.0.1:63973: How are you?
127.0.0.1:63973 has left

$ ./netcat3 
You are 127.0.0.1:63973
127.0.0.1:63969: Hello!
How are you?
127.0.0.1:63973: How are you?
^C
*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
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

type client chan<- string // an outgoing message channel

var (
    entering = make(chan client)
    leaving  = make(chan client)
    messages = make(chan string) // all incoming client messages
)

func broadcaster() {
    clients := make(map[client]bool) // all connected clients
    for {
        select {
            case msg := <-messages:
                // Broadcast incoming messages to all clients' outgoing message channels
                for cli := range clients {
                    cli <- msg
                }
            case cli := <- entering:
                clients[cli] = true
            case cli := <- leaving:
                delete(clients, cli)
                close(cli)
        }
    }
}

func handleConn(conn net.Conn) {
    ch := make(chan string) // outgoing clients messages
    go clientWriter(conn, ch)

    who := conn.RemoteAddr().String()
    ch <- "You are " + who
    messages <- who + " has arrived"
    entering <- ch

    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages <- who + ": " + input.Text()
    }
    // Note: ignoring potential errors from input.Err()

    leaving <- ch
    messages <- who + " has left"
    conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg) // Note: ignoring network errors
    }
}
