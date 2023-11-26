package ws

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"

	"golang.org/x/net/websocket"
)

func connect() (*websocket.Conn, error) {
	return websocket.Dial(fmt.Sprintf("ws://localhost:8080"), "tcp", "192.168.0.15")
}

func main2() {
	flag.Parse()
	// connect
	ws, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// receive
	var handler string
	go func() {
		for {
			err := websocket.JSON.Receive(ws, &handler)
			if err != nil {
				fmt.Println("Error receiving message: ", err.Error())
				break
			}
			fmt.Println("Message: ", handler)
		}
	}()
	// send
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		err = websocket.JSON.Send(ws, "Checking: "+string(rand.Int31()))
		if err != nil {
			fmt.Println("Error sending message: ", err.Error())
			break
		}
	}
}
