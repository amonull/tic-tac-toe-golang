package main

import (
  "fmt"
  "log"
)

const playerX rune = 'X'
const playerO rune = 'O'

func main() {
  board := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
  
  var running bool = true
  var gameCount = 0

  for running {
    if gameCount > 8 {
      fmt.Println("no more moves left")
      running = false
      break
    }

    printBoard(board)
    fmt.Println("Enter place: ")

    var userInput uint
    if _, err := fmt.Scan(&userInput); err != nil {
      log.Fatal("reading user input failed: ", err)
    }
    userInput-- // lowered for indexing
    
    // error checking
    if userInput > 8 {
      fmt.Println("invalid move try again")
      continue
    }

    // game
    if gameCount % 2 == 0 {
      if checkValid(board, userInput) {
        movePlayer(board, playerX, userInput)
      } else {
        fmt.Println("invalid move try again")
        continue
      }
      
      if isWon(board, playerX) {
        fmt.Println("Player: X Won")
        printBoard(board)
        running = false
        break
      }

    } else {
      if checkValid(board, userInput) {
        movePlayer(board, playerO, userInput)
      } else {
        fmt.Println("invalid move try again")
        continue
      }

      if isWon(board, playerO) {
        fmt.Println("Player: O Won")
        printBoard(board)
        running = false
        break
      }
    }

    gameCount++
  }
}

func printBoard(board []rune) {
  fmt.Println(string(board[0]), string(board[1]), string(board[2]))
  fmt.Println(string(board[3]), string(board[4]), string(board[5]))
  fmt.Println(string(board[6]), string(board[7]), string(board[8]))
}
