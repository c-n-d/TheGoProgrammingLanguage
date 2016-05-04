/*
Exercise 8.15 - Modify broadcaster to skip a message rather than wait if a client writer in not ready
                to accept it. Add buffering on the client outgoing message channel so that most messages are not 
                dropped.

*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "strings"
    "time"
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
    Name     string
    Mailbox  chan<- string // an outgoing message channel
    lastSend time.Time
}

func (c *client) Send(message string) {
    messages <- c.Name + ": " + message
    c.lastSend = time.Now()
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
                    // Non blocking send to each client
                    select {
                        case cli.Mailbox <- msg:
                            // send
                        default:
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
    // Add buffer to mailbox (5)
    ch := make(chan string, 5) // outgoing clients messages
    go clientWriter(conn, ch)

    who := namePrompt(conn)
    me := client{ who, ch, time.Now() }

    go watchdog(&me, conn)

    ch <- "You are " + who
    messages <- who + " has arrived"
    entering <- me

    input := bufio.NewScanner(conn)
    for input.Scan() {
        me.Send(input.Text())
    }
    // Note: ignoring potential errors from input.Err()

    leaving <- me
    messages <- who + " has left"
    conn.Close()
}

// Prompts new users for a nickname. Does not check if name is already in use.
func namePrompt(conn net.Conn) string {
    fmt.Fprintf(conn, "Please enter your name: ")

    input := bufio.NewScanner(conn)
    for input.Scan() {
        return input.Text()
    }
    return conn.RemoteAddr().String()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg) // Note: ignoring network errors
    }
}

func watchdog(c *client, conn net.Conn) {
    timer := time.Tick(1 * time.Second)
    for {
        select {
            case <-timer:
                if time.Since(c.lastSend) > 5 * time.Minute {
                    messages <- c.Name + " idle. Kicking"
                    conn.Close()
                    return
                }
        }
    }
}
