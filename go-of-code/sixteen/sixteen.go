package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
	"time"
)

func main() {
	log.SetPrefix("16:")
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 16")

	start := time.Now()
	fmt.Println(Sixteen())
	elapsed := time.Since(start)
	logger.Printf("Day 16 took %s", elapsed)
}

func Sixteen() (int, int) {
	maze, start, end, err := loadInput("../../inputs/16th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// maze.Print()
	end = maze.Dijkstra(start, end)

	maze.PrintPath(end)
	maze.Print()

	return end.d, maze.CountPath(end)
}

type Location struct {
	x   int
	y   int
	dir string
}

type Node struct {
	location Location
	d        int
	parents  []*Node
}

// Creates a new node with maximum distance and no parent
func NewNode(location Location) *Node {
	return &Node{location, int((^uint(0) >> 1) / 2), []*Node{}}
}

// Calculates the cost to move from the parent to the node
func (n *Node) Cost(parent *Node) int {
	// compare the parent direction to the node direction
	switch {
	case n.location.dir == parent.location.dir:
		return 1 + parent.d
	case (n.location.dir == "<" && parent.location.dir == ">") || (n.location.dir == ">" && parent.location.dir == "<") || (n.location.dir == "^" && parent.location.dir == "v") || (n.location.dir == "v" && parent.location.dir == "^"):
		return 2001 + parent.d
	default:
		return 1001 + parent.d
	}
}

func (n *Node) SetParent(parent *Node) {
	n.parents = []*Node{parent}
	n.d = n.Cost(parent)
}

func (n *Node) AddParent(parent *Node) {
	n.parents = append(n.parents, parent)
}

func (n *Node) Neighbors(maze *Maze) []*Node {
	targets := []Location{
		{x: -1, y: 0, dir: "<"},
		{x: 1, y: 0, dir: ">"},
		{x: 0, y: -1, dir: "^"},
		{x: 0, y: 1, dir: "v"},
	}

	var neighbors []*Node
	for _, target := range targets {
		candidate := Location{
			x:   n.location.x + target.x,
			y:   n.location.y + target.y,
			dir: target.dir,
		}
		if maze.maze[candidate.y][candidate.x] == "#" {
			continue
		}

		neighbors = append(neighbors, NewNode(candidate))
	}
	return neighbors
}

type Maze struct {
	maze    [][]string
	visited map[Location]*Node
}

func (m *Maze) Dijkstra(start, end *Node) *Node {
	// create a queue of unvisited nodes
	unvisited := []*Node{start}

	for len(unvisited) > 0 {
		// find the node with the lowest distance
		sort.Slice(unvisited, func(i, j int) bool {
			return unvisited[i].d < unvisited[j].d
		})

		current := unvisited[0]
		unvisited = unvisited[1:]

		// check if we have reached the end
		if current.location.x == end.location.x && current.location.y == end.location.y {
			if end.d > current.d {
				end.d = current.d
				end.parents = current.parents
			}

		}

		// add the neighbors to the queue initialized with the current node cost/distance
		neighbors := current.Neighbors(m)
		for _, neighbor := range neighbors {
			known, ok := m.visited[neighbor.location]
			if ok && known.d > neighbor.Cost(current) {
				known.SetParent(current)
				unvisited = append(unvisited, known)
			} else if ok && known.d == neighbor.Cost(current) && !slices.Contains(known.parents, current) {
				known.AddParent(current)
				unvisited = append(unvisited, known)
			} else if !ok {
				neighbor.SetParent(current)
				unvisited = append(unvisited, neighbor)
				m.visited[neighbor.location] = neighbor
			}
		}
		// mark point in maze as visited with x
		// m.maze[current.location.y][current.location.x] = "x"
		// m.Print()
	}
	return end
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

func (m *Maze) PrintPath(end *Node) {
	var printPath func(node *Node)
	printPath = func(node *Node) {
		if node == nil || len(node.parents) == 0 {
			return
		}
		for _, parent := range node.parents {
			printPath(parent)
			log.Printf("%d, %d -> %s, (%d)", parent.location.x, parent.location.y, parent.location.dir, parent.d)
			m.maze[parent.location.y][parent.location.x] = parent.location.dir
		}
	}
	printPath(end)
	log.Println()
}

type Position struct {
	x int
	y int
}

func (m *Maze) CountPath(end *Node) int {
	unique := make(map[Position]bool)
	var countPath func(node *Node)
	countPath = func(node *Node) {
		if node == nil {
			return
		}
		unique[Position{node.location.x, node.location.y}] = true
		for _, parent := range node.parents {
			countPath(parent)
		}
	}
	countPath(end)
	return len(unique)
}

func loadInput(filename string) (*Maze, *Node, *Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maze := make([][]string, 0)
	visited := make(map[Location]*Node)
	var start, end *Node

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
				start = NewNode(Location{x, y, ">"})
				start.d = 0
				visited[start.location] = start
			} else if string(obj) == "E" {
				end = NewNode(Location{x, y, ">"})
			}
			maze[y][x] = string(obj)
		}
		y++
	}

	return &Maze{maze, visited}, start, end, nil
}
