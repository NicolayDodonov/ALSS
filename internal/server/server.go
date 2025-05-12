package server

import (
	"artificialLifeGo/internal/ALSS"
	"artificialLifeGo/internal/config"
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
	mux  *http.ServeMux
	srv  *http.Server
	upg  *websocket.Upgrader
	conf *config.Config
}

func New(conf *config.Config) Server {
	mux := http.NewServeMux()
	return &WsServer{
		mux: mux,
		srv: &http.Server{
			Addr:    conf.IP + ":" + conf.Port,
			Handler: mux,
		},
		upg:  &websocket.Upgrader{},
		conf: conf,
	}
}

func (ws *WsServer) Start() error {
	ws.mux.Handle("/", http.FileServer(http.Dir(web)))

	ws.mux.HandleFunc("/ws", ws.wsHandler)

	return ws.srv.ListenAndServe()
}

func (ws *WsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upg.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(conn.RemoteAddr().String())
	go ws.commutation(conn)
}

func (ws *WsServer) commutation(conn *websocket.Conn) {
	init := Message{}
	if err := conn.ReadJSON(&init); err != nil {
		log.Println(err.Error())
		return
	}

	controller := ALSS.NewController(ws.conf, init.Count, init.Sea, init.Sea, init.Age, init.Energy)
	_ = controller
	//todo: make 2 chan
	for {
		//todo: check connection

		//todo: controller.Run(chanel, chanel, ctx)

	}
}
