package utils

import (
  "log"
  "golang.org/x/crypto/ssh/terminal"
  "github.com/fosterdill/gol/lib"
)

func ScreenSize() (lib.Coord, lib.Coord) {
  widthInt, heightInt, err := terminal.GetSize(0)

  if err != nil {
    log.Fatal(err)
  }

  return lib.Coord(widthInt), lib.Coord(heightInt)
}

func CreateScreen(width, height lib.Coord) *lib.Screen {
  screen := make([]string, height)

  for y := range screen {
    screen[y] = BuildString(width, lib.EmptySpace)
  }

  return &lib.Screen{ screen, width, height }
}