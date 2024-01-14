package gorTest

import (
	"fmt"
	"net"
	"os"
)

func Mainer() {
	fmt.Println(os.Args)
	var address string = "127.0.0.1:1200"
	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
	// 	os.Exit(1)
	// }
	// service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	for {
		var messages string = "Text sending message"
		fmt.Println("Type text: ")
		fmt.Scan(&messages)

		_, err = conn.Write([]byte(messages))
		checkError(err)
		res := make([]byte, 1024)

		_, err = conn.Read(res)
		checkError(err)

		fmt.Println("Sent message: ", messages)
		fmt.Println("Got message: ", string(res))
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
