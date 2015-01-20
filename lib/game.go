package lib

import (
  "fmt"
  "bytes"
  "time"
)

type Game struct {
  Screen *Screen
  NextScreen *Screen
}

func (game *Game) Start(speed time.Duration) {
  var buffer bytes.Buffer

  game.setup()
  game.showScreen(&buffer)

  for game.Screen.LiveCells() > 0 {
    game.iterate()
    game.showScreen(&buffer)
    time.Sleep(speed * time.Millisecond)
  }
}

func (game *Game) showScreen(buffer *bytes.Buffer) {
  screen := game.Screen

  for y := len(screen.Tiles) - 1; y > 0; y-- {
    buffer.WriteString(screen.Tiles[y] + "\n")
  }

  buffer.WriteString(screen.Tiles[0])
  fmt.Print(buffer.String())

  buffer.Reset()
}

func (game *Game) iterate() {

}

func (game *Game) setup() {
}