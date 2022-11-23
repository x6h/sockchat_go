package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

// client settings
const (
    SERVER_TYPE = "tcp"
    SERVER_HOST = "localhost"
    SERVER_PORT = "8080"
    SERVER_MSG_LENGTH = 512
)

func main() {
    // connect to server
    server, server_error := net.Dial(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)

    if server_error != nil {
        fmt.Printf("failed to connect to server. (error: %s)\n", server_error)
        os.Exit(1)
    }

    // setup input scanner (fmt.Scanln does exactly opposite of what is says it does. awesome.)
    input := bufio.NewScanner(os.Stdin)

    // start thread to recieve messages on
    go recieve_messages(server)

    // input loop
    for {
        input.Scan()

        // write input to server
        _, write_error := server.Write(input.Bytes())

        if write_error != nil {
            fmt.Printf("failed to send data. (error: %s)\n", write_error)
            os.Exit(1)
        }
    }
}

func recieve_messages(server net.Conn) {
    for {
        message := make([]byte, SERVER_MSG_LENGTH)
        _, read_error := server.Read(message)

        if read_error != nil {
            fmt.Printf("failed to read message from the server. (error: %s)\n", read_error)
        }

        fmt.Printf("-> %s\n", message)
    }
}
