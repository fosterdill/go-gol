package utils

import (
  "github.com/fosterdill/gol/lib"
)

func CreateGame() *lib.Game {
  width, height := ScreenSize()
  return &lib.Game{
    Screen: CreateScreen(width, height),
    NextScreen: CreateScreen(width, height),
  }
}