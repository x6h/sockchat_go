package main

import (
    "fmt"
    "net"
    "os"
)

// client settings
const (
    SERVER_TYPE = "tcp"
    SERVER_HOST = "localhost"
    SERVER_PORT = "8080"
)

func main() {
    // connect to server
    server, server_error := net.Dial(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)

    if server_error != nil {
        fmt.Printf("failed to connect to server. (error: %s)\n", server_error)
        os.Exit(1)
    }

    send_channel := make(chan int)

    // start send and receive threads
    go func() {
        SendMessages(server)
        send_channel <- 1
    }()
    go ReceiveMessages(server)

    // wait for send goroutine to finish
    <- send_channel
}
