package lib

type Screen struct {
  Tiles []string
}

func (screen *Screen) Revive(x, y Coord) {
  screen.setValue(x, y, FullSpace)
}

func (screen *Screen) Kill(x, y Coord) {
  screen.setValue(x, y, EmptySpace)
}

func (screen *Screen) At(x, y Coord) (value string) {
  return string(screen.Tiles[y][x])
}

func (screen *Screen) LiveCells() (liveCells int) {
  liveCells = 0

  for _, row := range screen.Tiles {
    for _, char := range row {
      if string(char) == FullSpace {
        liveCells++
      }
    }
  }

  return
}

//private
func (screen *Screen) setValue(x, y Coord, value string) {
  row := &screen.Tiles[y]
  *row = (*row)[:x] + value + (*row)[x + 1:]
}