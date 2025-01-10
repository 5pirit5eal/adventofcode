package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 20")

	start := time.Now()
	fmt.Println(Twenty())
	elapsed := time.Since(start)
	logger.Printf("Day 20 took %s", elapsed)
}

func Twenty() (int, int) {
	maze, start, end, err := loadInput("../../inputs/20th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//maze.Print()
	maze.CreateTrack(start, end)
	times := maze.Cheat(100)
	times2 := maze.Cheat2(100)

	return times, times2
}

type Location struct {
	x int
	y int
}

type Maze struct {
	visited map[Location]int
	maze    [][]string
}

func NewMaze(maze [][]string) *Maze {
	return &Maze{make(map[Location]int), maze}
}

func (m *Maze) Neighbors(location Location, walls bool, part2 bool) []Location {
	var targets []Location
	if part2 {
		targets = []Location{}
		for x := -20; x <= 20; x++ {
			if x < 0 {
				for y := -20 - x; y <= 20+x; y++ {
					targets = append(targets, Location{x, y})
				}
			} else {
				for y := -20 + x; y <= 20-x; y++ {
					targets = append(targets, Location{x, y})
				}
			}

		}
	} else {
		targets = []Location{
			{x: -1, y: 0},
			{x: 1, y: 0},
			{x: 0, y: -1},
			{x: 0, y: 1},
		}

	}

	var neighbors []Location
	for _, target := range targets {
		candidate := Location{
			x: location.x + target.x,
			y: location.y + target.y,
		}

		if candidate.x < 0 || candidate.x >= len(m.maze[0]) || candidate.y < 0 || candidate.y >= len(m.maze) {
			continue
		}

		if walls {
			if m.maze[candidate.y][candidate.x] == "#" {
				neighbors = append(neighbors, candidate)
			}
		} else if m.maze[candidate.y][candidate.x] != "#" {
			neighbors = append(neighbors, candidate)
		}
	}
	return neighbors
}

// Depth first search to create the track
func (m *Maze) CreateTrack(start Location, end Location) {
	m.visited[start] = 0
	queue := []Location{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, neighbor := range m.Neighbors(current, false, false) {
			if _, ok := m.visited[neighbor]; ok {
				continue
			}
			m.visited[neighbor] = m.visited[current] + 1
			queue = append(queue, neighbor)
		}
	}
}

// Iterate over the track and search for shortcuts by glitching into the wall
func (m *Maze) Cheat(threshold int) int {
	count := 0
	for location, traveled := range m.visited {
		for _, wall := range m.Neighbors(location, true, false) {
			for _, neighbor := range m.Neighbors(wall, false, false) {
				if m.visited[neighbor]-traveled >= threshold+2 {
					count++
				}
			}
		}
	}
	return count
}

// Iterate over the track and search for shortcuts by glitching into the wall up to 20 steps
func (m *Maze) Cheat2(threshold int) int {
	count := 0
	for location, traveled := range m.visited {
		for _, neighbor := range m.Neighbors(location, false, true) {
			var dx, dy int
			if neighbor.x < location.x {
				dx = location.x - neighbor.x
			} else {
				dx = neighbor.x - location.x
			}
			if neighbor.y < location.y {
				dy = location.y - neighbor.y
			} else {
				dy = neighbor.y - location.y
			}
			saved := m.visited[neighbor] - (traveled + dx + dy)
			if saved >= threshold {
				count++
			}
		}
	}
	return count
}

func (m *Maze) Print() {
	for _, row := range m.maze {
		var sb strings.Builder
		for _, cell := range row {
			sb.WriteString(cell)
		}
		log.Println(sb.String())
	}
	log.Println()
}

// Find all shortcuts in the maze by cheating two steps for each viable wall
func loadInput(filename string) (*Maze, Location, Location, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, Location{}, Location{}, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maze := make([][]string, 0)
	var start, end Location

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}
		maze = append(maze, make([]string, len(line)))
		for x, obj := range line {
			if string(obj) == "S" {
				start = Location{x, y}

			} else if string(obj) == "E" {
				end = Location{x, y}
			}
			maze[y][x] = string(obj)
		}
		y++
	}

	return NewMaze(maze), start, end, nil
}
