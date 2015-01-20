package gol

import (
  "github.com/fosterdill/gol/utils"
)

func Run() {
  game := utils.CreateGame()
  game.Start(100)
}