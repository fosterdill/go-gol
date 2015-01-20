package lib

type Screen struct {
  Tiles []string
  Width Coord
  Height Coord
}

func (screen *Screen) Keep(x, y Coord) {
  screen.set(x, y, FullSpace)
}

func (screen *Screen) Kill(x, y Coord) {
  screen.set(x, y, EmptySpace)
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

func (screen *Screen) LivingNeighbors(x, y Coord) (validNeighbors int) {
  neighbors := adjacentIndices(x, y)
  validNeighbors = 8

  for _, coords := range *neighbors {
    isOutOfBounds := coords[0] < 0 ||
                     coords[0] >= screen.Width ||
                     coords[1] < 0 ||
                     coords[1] >= screen.Height

    if isOutOfBounds || screen.IsDeadAt(coords[0], coords[1]) {
      validNeighbors--
    }
  }

  return
}

func (screen *Screen) Rows() []string {
  return screen.Tiles
}

func (screen *Screen) IsDeadAt(x, y Coord) bool {
  return screen.at(x, y) == EmptySpace
}

//private
func (screen *Screen) set(x, y Coord, value string) {
  row := &screen.Tiles[y]
  *row = (*row)[:x] + value + (*row)[x + 1:]
}

func (screen *Screen) at(x, y Coord) (value string) {
  return string(screen.Tiles[y][x])
}

//helpers
func adjacentIndices(x, y Coord) (indices *[8][2]Coord) {
  indices = &[8][2]Coord{
    [2]Coord{-1, -1},
    [2]Coord{-1, 1},
    [2]Coord{1, -1},
    [2]Coord{1, 1},
    [2]Coord{1, 0},
    [2]Coord{0, 1},
    [2]Coord{0, -1},
    [2]Coord{-1, 0},
  }

  for index, diff := range indices {
    (*indices)[index][0] = diff[0] + x
    (*indices)[index][1] = diff[1] + y
  }

  return
}