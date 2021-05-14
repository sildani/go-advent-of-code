package three

import (
  "testing"
)

func TestFindTreesInTerrain(t *testing.T) {
  var terrain []string
  var slope []int
  var expected int
  var actual int

  terrain = []string{
    "..##.......",
    "#...#...#..",
    ".#....#..#.",
    "..#.#...#.#",
    ".#...##..#.",
    "..#.##.....",
    ".#.#.#....#",
    ".#........#",
    "#.##...#...",
    "#...##....#",
    ".#..#...#.#",
  }
  slope = []int{3, 1}
  expected = 7
  actual = FindTreesInTerrain(terrain, slope)
  if expected != actual {
    t.Errorf("expected %v, got %v\n", expected, actual)
  }

  terrain = []string{
    "..##.......\r",
    "#...#...#..\r",
    ".#....#..#.\r",
    "..#.#...#.#\r",
    ".#...##..#.\r",
    "..#.##.....\r",
    ".#.#.#....#\r",
    ".#........#\r",
    "#.##...#...\r",
    "#...##....#\r",
    ".#..#...#.#",
  }
  slope = []int{3, 1}
  expected = 7
  actual = FindTreesInTerrain(terrain, slope)
  if expected != actual {
    t.Errorf("expected %v, got %v\n", expected, actual)
  }
}
