package server

import (
	"artificialLifeGo/internal/ALSS"
	"artificialLifeGo/internal/config"
	"artificialLifeGo/internal/logger/baseLogger"
	"context"
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
	srv  *http.Server
	mux  *http.ServeMux
	upg  *websocket.Upgrader
	conf *config.Config
	l    *baseLogger.Logger
}

func New(conf *config.Config, l *baseLogger.Logger) Server {
	mux := http.NewServeMux()
	return &WsServer{
		srv: &http.Server{
			Addr:    conf.IP + ":" + conf.Port,
			Handler: mux,
		},
		mux:  mux,
		upg:  &websocket.Upgrader{},
		conf: conf,
		l:    l,
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
	init := ALSS.Message{}
	if err := conn.ReadJSON(&init); err != nil {
		log.Println(err.Error())
		return
	}

	controller := ALSS.NewController(ws.conf, ws.l, init.Count, init.Sea, init.Sea, init.Age, init.Energy)
	_ = controller
	frameChan := make(chan *ALSS.Frame)
	//controlChan := make(chan *ALSS.Message) //todo: заменить

	controller.InitModel()

	go controller.Run(frameChan, context.TODO())
	for {
		msg, err := ws.getMessage(conn, context.TODO())
		if err != nil {
			log.Println(err.Error())
			return
		}
		_ = msg

		frame := <-frameChan

		if err := ws.sendMessage(conn, frame); err != nil {
			log.Printf(err.Error())
		}

	}
}

func (ws *WsServer) getMessage(conn *websocket.Conn, ctx context.Context) (*ALSS.Message, error) {
	message := ALSS.Message{}
	if err := conn.ReadJSON(&message); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &message, nil
}

func (ws *WsServer) sendMessage(conn *websocket.Conn, v interface{}) error {
	err := conn.WriteJSON(v)
	return err
}
