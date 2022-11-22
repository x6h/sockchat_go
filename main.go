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
)

func main() {
    // connect to server
    server, server_error := net.Dial(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)

    if server_error != nil {
        fmt.Printf("failed to connect to server. (error: %s)\n", server_error)
        os.Exit(1)
    }

    // setup input scanner (fmt.Scanln does exactly opposite of what is says it does. awesome.)
    input_scanner := bufio.NewScanner(os.Stdin)

    // input loop
    for
    {
        // scan for user input
        input_scanner.Scan()
        fmt.Printf("%s\n", input_scanner.Text())

        // write input to server
        _, write_error := server.Write(input_scanner.Bytes())

        if write_error != nil {
            fmt.Printf("failed to send data. (error: %s)\n", write_error)
            os.Exit(1)
        }
    }
}
