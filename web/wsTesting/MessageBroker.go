package ws

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	server("8080")
}

type hub struct {
	clients    map[string]*websocket.Conn
	deleteUser chan *websocket.Conn
	addUser    chan *websocket.Conn
	Message    chan string
}

func newHub() *hub {
	return &hub{map[string]*websocket.Conn{},
		make(chan *websocket.Conn),
		make(chan *websocket.Conn),
		make(chan string)}
}

func server(port string) error {
	h := newHub()
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		handler(ws, h)
	}))
	s := http.Server{Addr: ":" + port, Handler: mux}
	return s.ListenAndServe()
}

func (vs *hub) Run() {

	for {
		select {
		case val := <-vs.deleteUser:
			{
				vs.DeleteUser(val)
			}
		case val := <-vs.addUser:
			{
				vs.AddUser(val)
			}
		case val := <-vs.Message:
			{
				vs.SendMessageAll(val)
			}
		}
	}
}

func (cd *hub) DeleteUser(user *websocket.Conn) {
	delete(cd.clients, string(user.LocalAddr().String()))
}
func (cd *hub) AddUser(user *websocket.Conn) {
	cd.clients[string(user.LocalAddr().String())] = user
}
func (cd *hub) SendMessageAll(message string) {
	for _, client := range cd.clients {
		websocket.JSON.Send(client, message)
		fmt.Println("Message sent to :", client.LocalAddr().String())
	}
}

func handler(conu *websocket.Conn, hh *hub) {
	hh.Run()
	hh.addUser <- conu
	for {
		var message string
		err := websocket.JSON.Receive(conu, &message)
		if err != nil {
			hh.deleteUser <- conu
		}
		hh.Message <- message
	}
}
