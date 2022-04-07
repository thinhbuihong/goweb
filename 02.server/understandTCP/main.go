package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	var c chan string = make(chan string, 1)

	request(conn, c)

	respond(conn, c)
}

func request(conn net.Conn, c chan<- string) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			fmt.Println("=============METHOD ", m)
			c <- strings.Fields(ln)[1] //path
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn, c <-chan string) {
	body := `
	<!DOCTYPE html>
	<html lang="en">

	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>hello world</title>
	</head>

	<body>
		<h1>Hello</h1>
		<h2>` + <-c + `</h2>
	</body>

	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, body)
}
