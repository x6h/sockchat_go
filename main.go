package main

import (
    "fmt"
    "net"
    "os"
    "time"
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

    // start send and receive threads
    go SendMessages(server)
    go ReceiveMessages(server)

    // pause main thread until either send or receive thread signals a shutdown
    for !should_shutdown {
        time.Sleep(100 * time.Millisecond)
    }
}
