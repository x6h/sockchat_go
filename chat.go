package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

const CLIENT_MSG_LENGTH = 512

func SendMessages(server net.Conn) {
    // setup input scanner (fmt.Scanln does exactly opposite of what is says it does. awesome.)
    input := bufio.NewScanner(os.Stdin)

    for {
        input.Scan()

        // write input to server
        _, write_error := server.Write(input.Bytes())

        if write_error != nil {
            fmt.Printf("failed to send data. (error: %s)\n", write_error)
            break
        }
    }
}

func ReceiveMessages(server net.Conn) {
    for {
        message := make([]byte, CLIENT_MSG_LENGTH)
        _, read_error := server.Read(message)

        if read_error != nil {
            fmt.Printf("failed to receive data. (error: %s)\n", read_error)
            break
        }

        fmt.Printf("-> %s\n", message)
    }
}
