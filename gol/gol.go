package gol

import (
  "github.com/fosterdill/gol/utils"
)

func Run() {
  width, height := utils.ScreenSize()
  screen := utils.CreateScreen(width, height)
  game := utils.CreateGame(screen)
  game.Start()
}