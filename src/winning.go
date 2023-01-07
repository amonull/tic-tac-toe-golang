package main

func isSame(board []rune, player rune) bool {
  for _, element := range board {
    if element != player {
      return false
    }
  }
  return true
}

func isWon(board []rune, player rune) bool {
  
  rowWin := func() bool {
    top := isSame(board[6:9], player)
    middle := isSame(board[3:6], player)
    bottom := isSame(board[0:3], player)
    
    return top || middle || bottom
  }()

  colWin := func() bool {
    left := isSame([]rune{board[0], board[3], board[6]}, player)
    middle := isSame([]rune{board[1], board[4], board[7]}, player)
    right := isSame([]rune{board[2], board[5], board[8]}, player)

    return left || middle || right
  }()

  diagWin := func() bool {
    left := isSame([]rune{board[0], board[4], board[8]}, player)
    right := isSame([]rune{board[2], board[4], board[6]}, player)

    return left || right
  }()

  return rowWin || colWin || diagWin
}
