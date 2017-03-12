import { createStore, applyMiddleware, combineReducers } from 'redux'
import undoable, { includeAction, excludeAction } from 'redux-undo';
import thunkMiddleware from 'redux-thunk'

function switchTurn(currentPlayer) {
  switch(currentPlayer) {
    case "white":
      return "black"
    case "black":
      return "white"
  }
}

export const boardReducer = (state = { 
  pieces: [],
  currentTurn: "black"
 }, action) => {
  switch (action.type) {
    case 'PLACE':
      return Object.assign({},state,
      {
        pieces: state.pieces.concat([
          {
            x: action.x,
            y: action.y,
            color: action.color
          }
        ]),
        currentTurn: switchTurn(state.currentTurn)
      })
    case 'SET':
      return Object.assign({}, state,
      {
        pieces: action.pieces,
        currentTurn: action.currentTurn
      })
    default: 
      return state
  }
}

export const profileReducer = (state = { 
  color: "black"
 }, action) => {
  switch (action.type) {
    case 'SET_PROFILE':
      return Object.assign({}, state, {
        color: action.color
      })
    default:
      return state
  }
}

export const reducer = combineReducers({
  board: undoable(boardReducer, { filter: includeAction('PLACE') }),
  profile: profileReducer
})



export const initStore = (initialState) => {
  return createStore(reducer, initialState, applyMiddleware(thunkMiddleware))
}