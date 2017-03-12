package engine

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type position uint32

type color uint32
const (
	black color = iota
	white
)

type piece struct {
	x     position
	y     position
	color color
}

type Board struct {
	pieces []piece
}

type actionType string
type sender uint32
const (
	client sender = iota
	server
)

var strToSender = map[string]sender{
	"client": client,
	"server": server,
}

type Message struct {
	action     actionType // action type
	sender     sender     // who sent the message
	data       string     // JSON string
}

func copyBoard(b *Board) *Board {
	newPieces := make([]piece, len(b.pieces))
	copy(newPieces, b.pieces)
	return &Board{pieces: newPieces}
}

func isValidMove(b *Board, p piece) bool {
	return true
}

// move makes a copy of the Board and
// returns the copy with given piece played
// if it is a valid move, otherwise the copy
// without the place is returned
// `moveSuccessful` specifies whether the given
// piece was placed
func move(b *Board, p piece) (newBoard *Board, moveSuccessful bool) {
	newBoard = copyBoard(b)
	if isValidMove(newBoard, p) {
		newBoard.pieces = append(newBoard.pieces, p)
		moveSuccessful = true
	}
	return newBoard, moveSuccessful
}

func Reduce(b *Board, m *Message) (newBoard *Board, retMsg *Message) {
	retMsg = &Message{action: "CURRENT_STATE", sender: server, data: ""}
	switch m.action {
	case "GET_CURRENT_STATE":
		newBoard = copyBoard(b)
		return newBoard, retMsg
	case "MOVE":
		var p piece
		err := json.Unmarshal(([]byte)(m.data), &p)
		if err != nil {
			log.WithFields(log.Fields{
				"data": m.data,
			}).Warn(err)
			retMsg.action = "INVALID_DATA"
			retMsg.data = "error: could not deserialize message data as a piece"
			return copyBoard(b), retMsg
		}
		newBoard, moveOk := move(b, p)
		stateJSON, err := json.Marshal(newBoard)
		if err != nil {
			log.WithFields(log.Fields{
				"newBoard": newBoard,
			}).Fatal(err)
		}
		if moveOk {
			retMsg.action = "INVALID_MOVE"
		}
		retMsg.data = string(stateJSON)
		return newBoard, retMsg
	default:
		retMsg.action = "INVALID_ACTION"
		retMsg.data = "error: invalid action type"
		log.WithFields(log.Fields{
			"board": b,
			"message": m,
		}).Warn("invalid action type")
		return
	}
	log.Fatal("engine.Reduce: this should never happen")
	return
}
