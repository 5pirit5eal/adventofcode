package ten

import (
	utils "advent-of-code/go-of-code/utils"
	"bufio"
	"errors"
	"log"
	"os"
)

func Ten() (int, int) {
	f, err := os.Open("../inputs/tenth_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	matrix := [][]uint8{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var row []uint8
		for _, c := range s.Text() {
			row = append(row, uint8(c-'0'))
		}
		matrix = append(matrix, row)
	}

	topograficMap := NewTopograficMap(matrix)

	return topograficMap.CalculatePaths()

}

type Point struct {
	x      int
	y      int
	height uint8
}

type Positions map[Point]bool

type Tree struct {
	branches []*Tree
	position Point
}

type TopograficMap struct {
	topography [][]uint8
	width      int
	height     int
}

func NewTopograficMap(topography [][]uint8) *TopograficMap {
	topograficMap := &TopograficMap{
		topography,
		len(topography),
		len(topography[0]),
	}
	return topograficMap
}

func (t *TopograficMap) GetPoint(x, y int) (Point, error) {
	if t == nil {
		return Point{}, errors.New("topografic map is nil")
	}
	if x < 0 || x >= t.height {
		return Point{}, errors.New("x is out of bounds")
	}
	if y < 0 || y >= t.width {
		return Point{}, errors.New("y is out of bounds")
	}
	return Point{x, y, t.topography[x][y]}, nil
}

// Builds the tree by walking over the positions and checking if the top,bottom and side neighbours
// are exactly 1 larger than the tree position value, if so adds them to its branches
func (t *TopograficMap) buildTree(tree *Tree) *Tree {
	if tree == nil || t == nil {
		return nil
	}
	if tree.position.height == 9 {
		return tree
	}
	cross := [][]int{
		{utils.Max(tree.position.x-1, 0), tree.position.y},
		{utils.Min(tree.position.x+1, len(t.topography)-1), tree.position.y},
		{tree.position.x, utils.Max(tree.position.y-1, 0)},
		{tree.position.x, utils.Min(tree.position.y+1, len(t.topography[0])-1)},
	}

	for _, p := range cross {
		point, err := t.GetPoint(p[0], p[1])
		if err != nil {
			continue
		}
		if point.height == tree.position.height+1 {
			tree.branches = append(tree.branches, t.buildTree(&Tree{[]*Tree{}, point}))
		}
	}

	return tree
}

// Counts paths to the top of the map and returns the result
func (t *TopograficMap) walkTreeHeights(tree *Tree) []Point {

	if tree == nil || t == nil {
		return nil
	}
	result := []Point{}
	if tree.position.height == 9 {
		result = append(result, tree.position)
	}

	for _, b := range tree.branches {
		result = append(result, t.walkTreeHeights(b)...)
	}

	return result
}

func (t *TopograficMap) calculatePath(tree *Tree, ch chan int) {
	defer close(ch)
	if tree == nil || t == nil {
		return
	}
	t.buildTree(tree)
	heights := t.walkTreeHeights(tree)

	uniqueHeights := make(Positions)

	// return number of unique paths to heights
	ch <- len(heights)

	for _, height := range heights {
		uniqueHeights[height] = true
	}

	ch <- len(uniqueHeights)
}

func (t *TopograficMap) CalculatePaths() (int, int) {
	channels := []chan int{}
	for i := range t.height {
		for j := range t.width {
			position, err := t.GetPoint(i, j)
			if position.height != 0 || err != nil {
				continue
			}
			ch := make(chan int)
			channels = append(channels, ch)
			go t.calculatePath(&Tree{[]*Tree{}, position}, ch)
		}
	}

	countPaths := 0
	countUnique := 0
	for _, ch := range channels {
		n := 0
		for msg := range ch {
			if n%2 == 0 {
				countPaths += msg
			} else {
				countUnique += msg
			}
			n++
		}
	}

	return countUnique, countPaths

}
