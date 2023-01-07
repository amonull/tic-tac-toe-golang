package main

import (
  "unicode"
)

func checkValid(board []rune, move uint) bool {
  if unicode.IsDigit(rune(board[move])) {
    return true
  }
  return false
}

func movePlayer(board []rune, player rune, move uint) {
  board[move] = player
}
