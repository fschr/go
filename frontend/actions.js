export function placePiece(x, y, color, ws) {
  return dispatch => {
    dispatch({
      type: 'PLACE',
      x,
      y,
      color
    })

    ws.send(JSON.stringify({
      Action: 'MOVE',
      Sender: 'client',
      Data: JSON.stringify({
        X: x,
        Y: y,
        Color: color
      })
    }))
  }
}

export function setPieces(pieces, currentTurn) {
  console.log('setting pieces')
  console.log(pieces)
  return ({
    type: 'SET',
    pieces,
    currentTurn
  })
}

export function setProfile(color) {
  return ({
    color
  })
}
