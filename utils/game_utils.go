package utils

import (
  "github.com/fosterdill/gol/lib"
)

func CreateGame(screen *lib.Screen) *lib.Game {
  return &lib.Game{ screen }
}