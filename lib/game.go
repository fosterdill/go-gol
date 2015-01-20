package lib

import (
  "fmt"
)

type Game struct {
  Screen *Screen
  NextScreen *Screen
}

func (game *Game) Start() {
  game.setup()
  game.showScreen()

  for game.Screen.LiveCells() > 0 {
    game.iterate()
    game.showScreen()
  }
}

func (game *Game) showScreen() {
  screen := game.Screen
  for y := len(screen.Tiles) - 1; y >= 0; y-- {
    fmt.Println(screen.Tiles[y])
  }
}

func (game *Game) iterate() {
}

func (game *Game) setup() {
  game.Screen.Revive(1, 1)
}