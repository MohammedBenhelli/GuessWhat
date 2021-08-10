package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn, s *Server) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var rt RequestType
		if err := json.Unmarshal(p, &rt); err != nil {
			if err := conn.WriteMessage(1, []byte("Error can't read request")); err != nil {
				log.Fatal(err)
			}
		} else if route, ok := ROUTER[rt.Type]; ok {
			if err := route(s, conn, &p); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}

		log.Println("Client Connected")
		err = ws.WriteMessage(1, []byte("Hi Client!"))
		if err != nil {
			log.Println(err)
		}
		
		go reader(ws, s)
	}
}

func setupRoutes(s *Server) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint(s))
}

func Init() {
	s := Server{
		userList: UserList{make(map[string]*User)},
		channel:  Channel{make(map[string]*Lobby)},
	}
	fmt.Println("Hello World")
	setupRoutes(&s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
