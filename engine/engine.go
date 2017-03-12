package engine

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type position uint32

type color string
const (
	black color = "black"
	white color = "white"
)

type Piece struct {
	X     position
	Y     position
	Color color
}

type Board struct {
	Pieces []Piece
}

type actionType string

type sender string
const (
	client sender = "client"
	server sender = "server"
)

type Message struct {
	Action     actionType // action type
	Sender     sender     // who sent the message
	Data       string     // JSON string
}

func copyBoard(b *Board) *Board {
	newPieces := make([]Piece, len(b.Pieces))
	copy(newPieces, b.Pieces)
	return &Board{Pieces: newPieces}
}

func isValidMove(b *Board, p Piece) bool {
	return true
}

// move makes a copy of the Board and
// returns the copy with given piece played
// if it is a valid move, otherwise the copy
// without the place is returned
// `moveSuccessful` specifies whether the given
// piece was placed
func move(b *Board, p Piece) (newBoard *Board, moveSuccessful bool) {
	newBoard = copyBoard(b)
	if isValidMove(newBoard, p) {
		newBoard.Pieces = append(newBoard.Pieces, p)
		moveSuccessful = true
	}
	return newBoard, moveSuccessful
}

func Reduce(b *Board, m *Message) (newBoard *Board, retMsg *Message) {

	retMsg = &Message{Action: "CURRENT_STATE", Sender: server, Data: ""}

	if b == nil {
		b = &Board{Pieces: make([]Piece, 0)}
	}

	switch m.Action {
	case "GET_CURRENT_STATE":
		newBoard = copyBoard(b)
	case "MOVE":
		var p Piece
		err := json.Unmarshal(([]byte)(m.Data), &p)
		if err != nil {
			log.WithFields(log.Fields{
				"data": m.Data,
			}).Warn(err)
			retMsg.Action = "INVALID_DATA"
			retMsg.Data = "error: could not deserialize message data as a piece"
			return copyBoard(b), retMsg
		}
		var moveOk bool
		newBoard, moveOk = move(b, p)
		if !moveOk {
			retMsg.Action = "INVALID_MOVE"
		}
	default:
		retMsg.Action = "INVALID_ACTION"
		retMsg.Data = "error: invalid action type"
		log.WithFields(log.Fields{
			"board": b,
			"message": m,
		}).Warn("invalid action type")
		newBoard = copyBoard(b)
	}
	stateJSON, err := json.Marshal(*newBoard)
	if err != nil {
		log.WithFields(log.Fields{
			"newBoard": newBoard,
		}).Fatal(err)
	}
	retMsg.Data = string(stateJSON)
	return newBoard, retMsg
}
