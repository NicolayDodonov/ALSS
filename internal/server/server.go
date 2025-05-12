package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	web = "./web/"
)

type Server interface {
	Start() error
}

type WsServer struct {
	mux *http.ServeMux
	srv *http.Server
	Upg *websocket.Upgrader
}

func New(adr string) Server {
	mux := http.NewServeMux()
	return &WsServer{
		mux: mux,
		srv: &http.Server{
			Addr:    adr,
			Handler: mux,
		},
		Upg: &websocket.Upgrader{},
	}
}

func (ws *WsServer) Start() error {
	ws.mux.Handle("/", http.FileServer(http.Dir(web)))

	ws.mux.HandleFunc("/ws", ws.wsHandler)

	return ws.srv.ListenAndServe()
}

func (ws *WsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Upg.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(conn.RemoteAddr().String())
	go ws.commutation(conn)
}

func (ws *WsServer) commutation(conn *websocket.Conn) {
	for {
		msg := Message{}
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(msg.WorldSeason + " " + msg.StartCountAgent)

		conn.WriteJSON(Message{
			msg.WorldSeason + "12",
			msg.StartCountAgent + "13",
		})
	}
}
