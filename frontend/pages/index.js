import React, { Component } from 'react';
import { Baduk, BadukBoard, Piece } from 'react-baduk';

import { initStore } from '../store'
import withRedux from 'next-redux-wrapper'
import { connect } from 'react-redux'
import { placePiece, setPieces } from '../actions'
import { ActionCreators } from 'redux-undo';

import _ from 'lodash'

class Index extends Component {
  constructor(){
    super()

    this._ws = undefined

    this.clickEmptyTile = this.clickEmptyTile.bind(this)
  }

  static getInitialProps ({store, isServer}) {
    return { isServer }
  }

  componentDidMount() {
    this._ws = new WebSocket('ws://localhost:3001/websocket/v1')

    this._ws.addEventListener('open', e => {
      this._ws.send(JSON.stringify({
        Action: 'GET_CURRENT_STATE',
        Sender: 'client'
      }))
    })

    this._ws.addEventListener('message', e => {
      const data = JSON.parse(e.data);
      console.log(data)

      if(data.sender !== 'client'){
        switch(data.Action) {
          case 'CURRENT_STATE':
            const serverData = JSON.parse(data.Data)
            this.props.dispatch(setPieces(serverData.Pieces.map(piece => _.mapKeys(piece, (val, key) => _.lowerCase(key))), "black"))
            break;
          case 'INVALID_MOVE':
            this.props.dispatch(ActionCreators.undo())
            break;
        }
      }
    })

    // this.props.dispatch(setPieces([
    //   {x: 0, y: 0, color: "black"},
    //   {x: 0, y: 1, color: "white"}
    // ], "black"))

    // setTimeout(() => {
    //   this.props.dispatch(ActionCreators.undo())
    // }, 3000)
  }

  clickEmptyTile(x, y) {
    if(this.props.profile.color === this.props.board.present.currentTurn)
      this.props.dispatch(placePiece(x, y, this.props.board.present.currentTurn, this._ws))
  }

  render() {
    return (
      <div>
        <style>{`
          .board {
            margin: 45px 45px 30px 30px;
          }
          .board svg {
            overflow-x: visible;
            overflow-y: visible;
            width: 400px;
            height: 400px;
          }
          .board svg rect {
            fill: rgba(255, 255, 255, 0.08);
          }

          .board svg line {
            stroke: rgba(0, 0, 0, 0.8);
            stroke-width: 0.05;
            stroke-linecap: square;
          }

          .table {
            box-shadow: 0 0 12px rgba(0, 0, 0, 0.4);
            display: inline-block;
            background-color: rgb(221, 180, 84);
          }

          .label {
            text-transform: uppercase;
            font-size: 0.8px;
            font-weight: bold;
            font-family: 'Helvetica', sans-serif;
          }

          .label.y-label {
            alignment-baseline: middle;
          }

          .piece {
            box-shadow: 0 0 2px rgba(0, 0, 0, 0.4);
          }

          .piece:hover {
            stroke: green;
            stroke-width: 0.1;
          }

          .piece.black {
            fill: black;
          }

          .piece.white {
            fill: white;
          }

          .piece-target {
            opacity: 0;
          }
        `}</style>
        <center>
          <h1>James; the Goat</h1>
          <BadukBoard size={19} labelStyle="hybrid"
            onClickEmpty={this.clickEmptyTile}
          >
          {
            this.props.board.present.pieces.map(piece => <Piece key={`${piece.color}${piece.x}${piece.y}`} {...piece} />)
          }
          </BadukBoard>
        </center>
      </div>
    );
  }
}

export default withRedux(initStore)(connect(state => state)(Index));
