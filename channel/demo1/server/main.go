package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	var (
		err      error
		conn     net.Conn
		listener net.Listener
	)
	if listener, err = net.Listen("tcp", "localhost:8080"); err != nil {
		panic(err)
	}
	go broadcaster()
	for {
		if conn, err = listener.Accept(); err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	var (
		cli     client
		clients map[client]bool
		msg     string
	)
	clients = make(map[client]bool)
	for {
		select {
		case msg = <-messages:
			for cli = range clients {
				cli <- msg
			}
		case cli = <-entering:
			clients[cli] = true
		case cli = <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	var (
		who   string      = conn.RemoteAddr().String()
		ch    chan string = make(chan string)
		input *bufio.Scanner
	)
	go clientWriter(conn, ch)

	ch <- "你是 " + who
	messages <- who + "进入"
	entering <- ch

	input = bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + "：" + input.Text()
	}

	leaving <- ch
	messages <- who + "离开"
	conn.Close()

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
