package utils

import (
  "github.com/fosterdill/gol/lib"
)

func CreateGame() *lib.Game {
  width, height := ScreenSize()
  blankScreen := CreateScreen(width, height)
  return &lib.Game{ Screen: blankScreen, NextScreen: blankScreen }
}