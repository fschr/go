package server

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/fschr/go/engine"
	"github.com/gorilla/websocket"
)

const (
	ServerErr     = "500: server error"
	BadRequestErr = "400: malformed message"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	state []engine.Board
	port  string
}

func NewServer(port string) *Server {
	return &Server{state: make([]engine.Board, 0), port: port}
}

func (s *Server) Run() {
	http.HandleFunc("/websocket/v1", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Warn(err)
			return
		}
		defer conn.Close()
		for {
			msgType, msgBytes, err := conn.ReadMessage()
			if err != nil {
				log.Warn(err)
				break
			}
			var m engine.Message
			err = json.Unmarshal(msgBytes, &m)
			if err != nil {
				log.WithFields(log.Fields{
					"msgBytes": string(msgBytes),
				}).Warn(err)
				err = conn.WriteMessage(msgType, ([]byte)(BadRequestErr))
				if err != nil {
					log.Fatal(err)
					return
				}
				continue
			}
			log.WithFields(log.Fields{
				"m": m,
			}).Info("received message")
			var lastBoard *engine.Board = nil
			if len(s.state) > 0 {
				lastBoard = &s.state[len(s.state)-1]
			}
			newState, retMsg := engine.Reduce(lastBoard, &m)
			s.state = append(s.state, *newState)
			retStr, err := json.Marshal(retMsg)
			if err != nil {
				log.WithFields(log.Fields{
					"retMsg": retMsg,
				}).Warn(err)
				err = conn.WriteMessage(msgType, ([]byte)(ServerErr))
				if err != nil {
					log.Fatal(err)
					return
				}
				continue
			}
			err = conn.WriteMessage(msgType, retStr)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	})
	if err := http.ListenAndServe(s.port, nil); err != nil {
		log.Fatal(err)
	}
}
