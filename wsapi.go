package regolt

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const WS_URL = "wss://ws.revolt.chat"

type Session struct {
	sync.RWMutex

	ws      *websocket.Conn
	wsMutex sync.RWMutex

	http *http.Client

	handlers map[string][]eventHandler

	Token string
}

func (s *Session) Open() error {
	var err error

	if s.Token == "" {
		return fmt.Errorf("no token")
	}

	s.Lock()
	defer s.Unlock()

	if s.ws != nil {
		return fmt.Errorf("ws open already")
	}

	s.ws, _, err = websocket.DefaultDialer.Dial(WS_URL, nil)
	if err != nil {
		return err
	}

	s.authenticate()

	close := make(chan struct{})
	go s.listener(close)
	go s.heartbeat(close)

	return nil

}

func (s *Session) listener(close <-chan struct{}) {
	for {
		_, message, err := s.ws.ReadMessage()
		if err != nil {
			return
		}
		select {
		case <-close:
			return
		default:
			go s.handler(message)
		}
	}
}

func (s *Session) heartbeat(close <-chan struct{}) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	h := struct {
		Type string `json:"type"`
		Data int    `json:"data"`
	}{
		Type: "Ping",
		Data: 0,
	}

	for {
		s.wsMutex.Lock()
		s.ws.WriteJSON(h)
		s.wsMutex.Unlock()

		select {
		case <-ticker.C:
		case <-close:
			return
		}
	}
}
