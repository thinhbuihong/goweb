package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "how is you day?")
		fmt.Fprintf(conn, "%v", "well, i hope!")
		conn.Close()
	}
}
