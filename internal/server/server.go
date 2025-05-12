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
	ws.mux.HandleFunc("/initGame", ws.initGame)

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
}

func (ws *WsServer) initGame(w http.ResponseWriter, r *http.Request) {
	//todo: создаём отдельному игроку его контроллер
	//todo: возможно создаём горутину???
}
