package three

import (
	"fmt"
	"strings"
)

func FindTreesInTerrain(terrain []string, slope []int) int {
  treeCount := 0
  x         := 0
  y         := 0
  stop      := false

  for i, line := range terrain {
    terrain[i] = strings.ReplaceAll(line, "\r", "")
  }

  for !stop {
    x = x+slope[0]
    y = y+slope[1]
    if y >= len(terrain) {
      stop = true
      fmt.Printf("FindTreesInTerrain; Breaking out of loop; x: %v; y: %v; treeCount: %v;\n", x, y, treeCount)
    } else {
      if x >= len(terrain[y]) {
        x = x - len(terrain[y])
      }
      if string(terrain[y][x]) == "#" {
        treeCount++
      }
      fmt.Printf("FindTreesInTerrain; Incrementing tree count; x: %v; y: %v; terrain[%v]: %v; treeCount: %v;\n", x, y, y, terrain[y], treeCount)
    }
  }

  return treeCount
}