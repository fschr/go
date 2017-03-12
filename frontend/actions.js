export function placePiece(x, y, color, ws) {
  return dispatch => {
    dispatch({
      type: 'PLACE',
      x,
      y,
      color
    })

    ws.send(JSON.stringify({
      action: '',
      sender: 'client',
      data: {
        x,
        y,
        color
      }
    }))
  }
}

export function setPieces(pieces, currentTurn) {
  console.log('setting pieces')
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