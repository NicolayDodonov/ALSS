package server

import (
	"artificialLifeGo/internal/ALSS"
	"artificialLifeGo/internal/config"
	"artificialLifeGo/internal/logger/baseLogger"
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	web = "./web/"
)

type Server interface {
	Start() error
}

// WsServer структура сервера.
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

// Start настраивает адреса сервера и запускает сервер в постоянном ожидании
// новых сообщений.
func (ws *WsServer) Start() error {
	ws.l.Info("Starting server...")

	//запускаем файл сервер
	ws.mux.Handle("/", http.FileServer(http.Dir(web)))
	//определяем канал web-socket
	ws.mux.HandleFunc("/ws", ws.wsHandler)

	ws.l.Info("Listening on http://" + ws.srv.Addr)
	return ws.srv.ListenAndServe()
}

// wsHandler обрабатывает запросы к Web-Socket соединени, создавая его и запуская
// горутину общения клиента и экземпляра ALSS.
func (ws *WsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	// создаём новый web-socket канал
	conn, err := ws.upg.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ws.l.Info(conn.RemoteAddr().String())

	// и начинаем общаться в нём.
	go ws.commutation(conn)
}

// commutation управляет ходом работы сессии между клиентом и сервером
func (ws *WsServer) commutation(conn *websocket.Conn) {
	defer func() {
		ws.l.Info("CM1: commutation end")
		conn.Close()
	}()
	// ожидаем поступления инициализирующего сообщения от клиента
	init := ALSS.Message{}
	if err := conn.ReadJSON(&init); err != nil {
		ws.l.Error("CM2: " + err.Error())
		return
	}

	// создаём контроллер модели и канал обмена фреймами.
	controller := ALSS.NewController(ws.conf, ws.l, init.Count, init.Sea, init.Sea)
	frameChan := make(chan *ALSS.Frame)

	// настраиваем модель
	controller.InitModel()
	// задаём внешний контекст прекращения работы
	ctxRun, cancelRun := context.WithCancel(context.Background())
	// запускаем модель на исполнение
	go controller.Run(frameChan, ctxRun)
	// бесконечно обрабатываем сообщения от сервера к клиенту и от клиента к серверу
	for {
		// настраиваем контекст ожидания нового сообщения
		ctxMSG, cancelMSG := context.WithTimeout(
			context.Background(),
			time.Microsecond*time.Duration(
				ws.conf.TimeStop))
		defer cancelMSG()
		// ждём от клиента сообщение или ошибки закрытия канала.
		msg, err := ws.getMessage(conn, ctxMSG)
		if err != nil {
			ws.l.Error("CM2: " + err.Error())
			cancelRun()
			return
		}

		// работает с полученным сообщением
		// todo: создать обработчик сообщений к модели
		_ = msg

		// получаем от модели единичный фрейм
		frame := <-frameChan
		// и отправляем его клиенту
		if err := ws.sendMessage(conn, frame); err != nil {
			ws.l.Error("CM3: " + err.Error())
			cancelRun()
			return
		}
	}
}

// getMessage считывает от клиента входящее сообщение. Если сообщение не приходит за время работы контекста,
// завершает работу возвращает в обоих переменных nil. Либо может вернуть или сообщение от пользователя, либо
// ошибку закрытия канала.
func (ws *WsServer) getMessage(conn *websocket.Conn, ctx context.Context) (*ALSS.Message, error) {

	// создаём структуру нового сообщения
	message := ALSS.Message{}
	// и ждём чтения сообщения
	done := make(chan error)
	go func() {
		defer func() {
			// КОСТЫЛЬ!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
			_ = recover()
		}()

		if err := conn.ReadJSON(&message); err != nil {
			done <- err
			return
		}
	}()

	//ожидаем одно из двух событий
	select {
	// получили информацию от клиента
	case err := <-done:
		{
			if err != nil {
				return nil, err // возвращаем ошибку
			}
			return &message, nil // возвращаем сообщение
		}
	// контекст закончился
	case <-ctx.Done():
		{
			return nil, nil // возвращаем ничего
		}
	}
}

// sendMessage отправляет клиенту сообщение в JSON формате.
func (ws *WsServer) sendMessage(conn *websocket.Conn, v interface{}) error {
	defer func() {
		// КОСТЫЛЬ!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		_ = recover()
	}()

	err := conn.WriteJSON(v)
	return err
}
