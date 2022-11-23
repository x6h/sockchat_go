package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

const CLIENT_MSG_LENGTH = 512

var should_shutdown = false

func SendMessages(server net.Conn) {
    // setup input scanner (fmt.Scanln does exactly opposite of what is says it does. awesome.)
    input := bufio.NewScanner(os.Stdin)

    for {
        if should_shutdown {
            break
        }

        input.Scan()

        if input.Text() == "/quit" {
            should_shutdown = true
            continue
        }

        // write input to server
        _, write_error := server.Write(input.Bytes())

        if write_error != nil {
            fmt.Printf("failed to send data. (error: %s)\n", write_error)
            should_shutdown = true
            break
        }
    }
}

func ReceiveMessages(server net.Conn) {
    for {
        if should_shutdown {
            break
        }

        message := make([]byte, CLIENT_MSG_LENGTH)
        _, read_error := server.Read(message)

        if read_error != nil {
            fmt.Printf("failed to receive data. (error: %s)\n", read_error)
            should_shutdown = true
            break
        }

        fmt.Printf("-> %s\n", message)
    }
}
