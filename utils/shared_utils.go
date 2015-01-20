package utils

import (
  "github.com/fosterdill/gol/lib"
  "bytes"
)

func BuildString(length lib.Coord, char string) string {
  var buffer bytes.Buffer

  for i := lib.Coord(0); i < length; i++ {
    buffer.WriteString(char)
  }

  return buffer.String()
}