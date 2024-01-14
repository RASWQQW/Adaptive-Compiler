package gorTest

import (
	"fmt"
	"net"
	"os"
	"time"
)

// MAKE SERVER WHICH SENDS TO ALL CONNECTED MESSAGES IN SHORT MAKE REAL CHAT
func Mainerd() {
	service := ":1341"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 1024) // set maximum request length to 128B to prevent flood based attacks
	defer conn.Close()            // close connection before exit
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println("Inside error")
			fmt.Println(err)
			break
		}

		time.Sleep(time.Second * 3)
		if read_len > 0 {
			conn.Write([]byte("Message gotten"))
		} else {
			conn.Write([]byte("Message recieved"))
		}
	}
}

func checkErrorss(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
