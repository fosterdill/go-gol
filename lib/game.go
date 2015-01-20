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
  game.showScreen(&buffer, speed)

  for game.Screen.LiveCells() > 0 {
    game.iterate()
    game.showScreen(&buffer, speed)
  }
}

func (game *Game) showScreen(buffer *bytes.Buffer, speed time.Duration) {
  screen := game.Screen

  for y := len(screen.Tiles) - 1; y > 0; y-- {
    buffer.WriteString(screen.Tiles[y] + "\n")
  }

  buffer.WriteString(screen.Tiles[0])
  fmt.Print(buffer.String())

  buffer.Reset()
  time.Sleep(speed * time.Millisecond)
}

func (game *Game) iterate() {
  for y, row := range game.Screen.Rows() {
    for x := range row {
      isDead := game.Screen.IsDeadAt(Coord(x), Coord(y))
      numberOfLivingNeighbors := game.Screen.LivingNeighbors(Coord(x), Coord(y))
      shouldSpawn := numberOfLivingNeighbors == 3
      shouldDie := numberOfLivingNeighbors < 2 || numberOfLivingNeighbors > 3

      if isDead {
        if shouldSpawn {
          game.NextScreen.Keep(Coord(x), Coord(y))
        } else {
          game.NextScreen.Kill(Coord(x), Coord(y))
        }
      } else {
        if shouldDie {
          game.NextScreen.Kill(Coord(x), Coord(y))
        } else {
          game.NextScreen.Keep(Coord(x), Coord(y))
        }
      }
    }
  }

  game.swapScreens()
}

func (game *Game) swapScreens() {
  game.Screen, game.NextScreen = game.NextScreen, game.Screen
}

func (game *Game) setup() {
  game.Screen.Keep(2, 0)
  game.Screen.Keep(2, 1)
  game.Screen.Keep(2, 2)
  game.Screen.Keep(1, 2)
  game.Screen.Keep(0, 1)

  game.Screen.Keep(2+5, 0+5)
  game.Screen.Keep(2+5, 1+5)
  game.Screen.Keep(2+5, 2+5)
  game.Screen.Keep(1+5, 2+5)
  game.Screen.Keep(0+5, 1+5)

  game.Screen.Keep(2+8, 0+5)
  game.Screen.Keep(2+8, 1+5)
  game.Screen.Keep(2+8, 2+5)
  game.Screen.Keep(1+8, 2+5)
  game.Screen.Keep(0+8, 1+5)
}
