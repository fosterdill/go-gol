package lib

import (
  "fmt"
)

type Game struct {
  Screen *Screen
}

func (game *Game) Start() {
  game.ShowScreen()
}

func (game *Game) ShowScreen() {
  screen := game.Screen

  for y := len(screen.Tiles) - 1; y >= 0; y-- {
    fmt.Println(screen.Tiles[y])
  }
}