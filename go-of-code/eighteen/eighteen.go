package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 18")

	start := time.Now()
	fmt.Println(Eighteen())
	elapsed := time.Since(start)
	logger.Printf("Day 18 took %s", elapsed)
}

func Eighteen() (int, int) {
	var end *Node
	var maze *Maze
	var block int
	for i := 1025; i < 3046; i++ {
		maze, err := loadInput("../../inputs/18th_day_input.txt", 71, i)
		if err != nil {
			log.Fatal(err)
		}
		// maze.Print()
		end = maze.Dijkstra()

		if len(end.parents) == 0 {
			block = i
			break
		}

	}
	maze, err := loadInput("../../inputs/18th_day_input.txt", 71, 2024)
	if err != nil {
		log.Fatal(err)
	}
	// maze.Print()
	end = maze.Dijkstra()
	maze.PrintPath(end, true)
	maze.Print()

	// The block coordinate == index+1
	return end.d, block
}

type Location struct {
	x int
	y int
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

func (n *Node) SetParent(parent *Node) {
	n.parents = []*Node{parent}
	n.d = parent.d + 1
}

func (n *Node) AddParent(parent *Node) {
	n.parents = append(n.parents, parent)
}

func (n *Node) Neighbors(maze *Maze) []*Node {
	targets := []Location{
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
	}

	var neighbors []*Node
	for _, target := range targets {
		candidate := Location{
			x: n.location.x + target.x,
			y: n.location.y + target.y,
		}
		if candidate.x < 0 || candidate.x >= len(*maze) || candidate.y < 0 || candidate.y >= len(*maze) {
			continue
		}
		if (*maze)[candidate.y][candidate.x] == "#" {
			continue
		}

		neighbors = append(neighbors, NewNode(candidate))
	}
	return neighbors
}

type Maze [][]string

func (m *Maze) Dijkstra() *Node {
	// create a queue of unvisited nodes
	start := &Node{Location{0, 0}, 0, []*Node{}}
	end := NewNode(Location{len(*m) - 1, len(*m) - 1})
	unvisited := []*Node{start}
	visited := make(map[Location]*Node)
	visited[start.location] = start

	for len(unvisited) > 0 {
		// find the node with the lowest distance
		sort.Slice(unvisited, func(i, j int) bool {
			return unvisited[i].d < unvisited[j].d
		})

		current := unvisited[0]
		unvisited = unvisited[1:]

		// check if we have reached the end
		if current.location == end.location {
			if end.d > current.d {
				end.d = current.d
				end.parents = current.parents
			} else if end.d == current.d {
				for _, parent := range current.parents {
					if !slices.Contains(end.parents, parent) {
						end.AddParent(parent)
					}
				}
			}
		}

		// add the neighbors to the queue initialized with the current node cost/distance
		neighbors := current.Neighbors(m)
		for _, neighbor := range neighbors {
			known, ok := visited[neighbor.location]
			if ok && known.d > current.d+1 {
				known.SetParent(current)
				unvisited = append(unvisited, known)
			} else if ok && known.d == current.d+1 && !slices.Contains(known.parents, current) {
				known.AddParent(current)
				unvisited = append(unvisited, known)
			} else if !ok {
				neighbor.SetParent(current)
				unvisited = append(unvisited, neighbor)
				visited[neighbor.location] = neighbor
			}
		}
		// mark point in maze as visited with x
		// m.maze[current.location.y][current.location.x] = "x"
		// m.Print()
	}
	return end
}

func (m *Maze) Print() {
	for _, row := range *m {
		var sb strings.Builder
		for _, cell := range row {
			sb.WriteString(cell)
		}
		log.Println(sb.String())
	}
	log.Println()
}

func (m *Maze) PrintPath(end *Node, all bool) {
	unique := make(map[Location]bool)
	var printPath func(node *Node)
	printPath = func(node *Node) {
		if node == nil || len(node.parents) == 0 {
			return
		}
		unique[node.location] = true
		if all {
			for _, parent := range node.parents {
				if unique[parent.location] {
					continue
				}
				printPath(parent)
				log.Printf("%d, %d -> (%d)", parent.location.x, parent.location.y, parent.d)
				(*m)[parent.location.y][parent.location.x] = "O"
			}
		} else {
			printPath(node.parents[0])
			log.Printf("%d, %d -> (%d)", node.parents[0].location.x, node.parents[0].location.y, node.parents[0].d)
			(*m)[node.parents[0].location.y][node.parents[0].location.x] = "O"
		}

	}
	printPath(end)
	log.Println()
}

func (m *Maze) CountPath(end *Node) int {
	unique := make(map[Location]bool)
	var countPath func(node *Node)
	countPath = func(node *Node) {
		if node == nil {
			return
		}
		unique[node.location] = true
		for _, parent := range node.parents {
			if unique[parent.location] {
				continue
			}
			countPath(parent)
		}
	}
	countPath(end)
	return len(unique)
}

func loadInput(filename string, size, limit int) (*Maze, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maze := make(Maze, size)

	for i := 0; i < size; i++ {
		maze[i] = make([]string, size)
		for j := 0; j < size; j++ {
			maze[i][j] = "."
		}
	}

	nBytes := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if nBytes == limit {
			break
		}

		if line == "" {
			break
		}

		nBytes++
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			// handle the case when the line does not have exactly 2 parts
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			// handle the case when the first part cannot be converted to an integer
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			// handle the case when the second part cannot be converted to an integer
			return nil, fmt.Errorf("invalid line: %q", line)
		}
		maze[y][x] = "#"
	}

	return &maze, nil
}
