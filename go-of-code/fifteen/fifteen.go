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
	log.SetPrefix("14:")
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 14")

	start := time.Now()
	fmt.Println(Fifteen())

	elapsed := time.Since(start)
	logger.Printf("Day 14 took %s", elapsed)
}

func Fifteen() (int, int) {
	warehouse, _, err := loadInput1("../../inputs/15th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	wideWarehouse, directions, err := loadInput2("../../inputs/15th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	warehouse.PrintRoom()
	log.Println()
	log.Println(directions)
	log.Println()

	for _, direction := range directions {
		warehouse.Move(direction)
	}
	warehouse.PrintRoom()
	log.Println()

	wideWarehouse.PrintRoom()
	log.Println()
	for _, direction := range directions {
		log.Println(direction)
		wideWarehouse.Move(direction)
		wideWarehouse.PrintRoom()
		log.Println()
	}
	wideWarehouse.PrintRoom()

	return warehouse.GPS(), wideWarehouse.GPS()
}

type Location struct {
	x int
	y int
}

type Warehouse struct {
	room  [][]string
	robot Location
}

func (w *Warehouse) Move(direction string) {
	w.move(direction, w.robot, "@")
}

func (w *Warehouse) move(direction string, xy Location, symbol string) {
	var target Location
	switch direction {
	case "<":
		target = Location{xy.x - 1, xy.y}
	case ">":
		target = Location{xy.x + 1, xy.y}
	case "v":
		target = Location{xy.x, xy.y + 1}
	case "^":
		target = Location{xy.x, xy.y - 1}
	default:
		log.Fatal("Invalid direction")
	}
	switch {
	case w.room[target.y][target.x] == "O":
		w.move(direction, target, "O")
	case w.room[target.y][target.x] == "[" && strings.ContainsAny(direction, "v^"):
		if !(w.peek(direction, target) && w.peek(direction, Location{target.x + 1, target.y})) {
			return
		}
		w.move(direction, target, "[")
		w.move(direction, Location{target.x + 1, target.y}, "]")
	case w.room[target.y][target.x] == "]" && strings.ContainsAny(direction, "v^"):
		if !(w.peek(direction, target) && w.peek(direction, Location{target.x - 1, target.y})) {
			return
		}
		w.move(direction, target, "]")
		w.move(direction, Location{target.x - 1, target.y}, "[")
	case strings.ContainsAny(w.room[target.y][target.x], "[]"):
		w.move(direction, target, w.room[target.y][target.x])
	}

	if w.room[target.y][target.x] == "." {
		w.room[target.y][target.x] = symbol
		w.room[xy.y][xy.x] = "."
		if symbol == "@" {
			w.robot = target
		}
	}

}

func (w *Warehouse) peek(direction string, xy Location) bool {
	var target Location
	switch direction {
	case "v":
		target = Location{xy.x, xy.y + 1}
	case "^":
		target = Location{xy.x, xy.y - 1}
	default:
		log.Fatal("Invalid direction")
	}
	switch {
	case w.room[target.y][target.x] == "[":
		left := w.peek(direction, target)
		right := w.peek(direction, Location{target.x + 1, target.y})
		return left && right
	case w.room[target.y][target.x] == "]":
		right := w.peek(direction, target)
		left := w.peek(direction, Location{target.x - 1, target.y})
		return right && left
	}
	return w.room[target.y][target.x] == "."
}

func (w *Warehouse) GPS() int {
	gps := 0
	for i := 0; i < len(w.room); i++ {
		for j := 0; j < len(w.room[0]); j++ {
			if strings.ContainsAny(w.room[i][j], "[O") {
				gps += 100*i + j
			}
		}
	}
	return gps
}

func (w *Warehouse) PrintRoom() {
	for _, row := range w.room {
		var sb strings.Builder
		for _, cell := range row {
			sb.WriteString(cell)
		}
		log.Println(sb.String())
	}
}

func loadInput1(filename string) (*Warehouse, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	directions := make([]string, 0)

	room := make([][]string, 0)
	var robot Location

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}
		room = append(room, make([]string, len(line)))
		for x, obj := range line {
			if string(obj) == "@" {
				robot = Location{x, y}
			}
			room[y][x] = string(obj)
		}
		y++
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		for i := 0; i < len(line); i++ {
			directions = append(directions, string(line[i]))
		}
	}

	return &Warehouse{room, Location{robot.x, robot.y}}, directions, nil
}

func loadInput2(filename string) (*Warehouse, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	directions := make([]string, 0)

	room := make([][]string, 0)
	var robot Location

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}
		room = append(room, make([]string, len(line)*2))
		for x, obj := range line {
			if string(obj) == "@" {
				robot = Location{x * 2, y}
			}
			switch string(obj) {
			case "O":
				room[y][x*2] = "["
				room[y][x*2+1] = "]"
			case "@":
				room[y][x*2] = "@"
				room[y][x*2+1] = "."
			default:
				room[y][x*2] = string(obj)
				room[y][x*2+1] = string(obj)
			}

		}
		y++
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		for i := 0; i < len(line); i++ {
			directions = append(directions, string(line[i]))
		}
	}

	return &Warehouse{room, Location{robot.x, robot.y}}, directions, nil
}
