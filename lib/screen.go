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

//helpers
func (screen *Screen) setValue(x, y Coord, value string) {
  row := &screen.Tiles[y]
  *row = (*row)[:x] + value + (*row)[x + 1:]
}