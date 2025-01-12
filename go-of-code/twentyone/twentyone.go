package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 21")

	start := time.Now()
	log.Println(TwentyOne())
	elapsed := time.Since(start)
	logger.Printf("Day 21 took %s", elapsed)
}

func loadInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	codes := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		codes = append(codes, line)
	}

	return codes, nil
}

func TwentyOne() (int, int) {
	codes, err := loadInput("../../inputs/21th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return First(codes), Second(codes)
}

// Abs returns the absolute value of a given integer.
func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func yPosition(num string) int {
	switch num {
	case "0", "A":
		return 3
	case "1", "2", "3":
		return 2
	case "4", "5", "6":
		return 1
	case "7", "8", "9":
		return 0
	}
	return -1
}

func xPosition(num string) int {
	switch num {
	case "1", "4", "7":
		return 0
	case "0", "2", "5", "8":
		return 1
	case "A", "3", "6", "9":
		return 2
	}
	return -1
}

func DoorSequence(start, end string) string {
	order := []string{"<", "v", "^", ">"}
	switch {
	case start == end:
		return "A"
	case (strings.Contains("0A", start) && strings.Contains("147", end)) || (strings.Contains("0A", end) && strings.Contains("147", start)):
		order = []string{"^", ">", "v", "<"}
	}

	stepsUpDown := yPosition(end) - yPosition(start)
	stepsLeftRight := xPosition(end) - xPosition(start)
	sequence := ""

	for _, i := range order {
		switch {
		case i == "<" && stepsLeftRight < 0:
			sequence += strings.Repeat("<", Abs(stepsLeftRight))
		case i == ">" && stepsLeftRight > 0:
			sequence += strings.Repeat(">", Abs(stepsLeftRight))
		case i == "^" && stepsUpDown < 0:
			sequence += strings.Repeat("^", Abs(stepsUpDown))
		case i == "v" && stepsUpDown > 0:
			sequence += strings.Repeat("v", Abs(stepsUpDown))
		}
	}

	return sequence + "A"
}

// Get the steps from on key to another.
// Valid keys are <, >, ^, v and A
var keyPadSequenceMap = map[string]map[string]string{
	"<": {
		">": ">>A",
		"A": ">>^A",
		"^": ">^A",
		"v": ">A",
	},
	">": {
		"<": "<<A",
		"A": "^A",
		"v": "<A",
		"^": "<^A",
	},
	"^": {
		"v": "vA",
		"A": ">A",
		">": "v>A",
		"<": "v<A",
	},
	"v": {
		"^": "^A",
		"A": "^>A",
		"<": "<A",
		">": ">A",
	},
	"A": {
		"<": "v<<A",
		">": "vA",
		"^": "<A",
		"v": "<vA",
	},
}

func KeyPadSequenceCost(start, end string, nRobots int) int {
	var seq string
	var ok bool
	if start == end {
		seq = "A"
		ok = true
	} else {
		seq, ok = keyPadSequenceMap[start][end]
	}

	if ok {
		if nRobots == 1 {
			return len(seq)
		} else {
			return Cost(seq, nRobots-1)
		}
	}

	log.Fatalf("Unknown start %s or end %s", start, end)
	return 0
}

var costCache sync.Map

type costKey struct {
	code    string
	nRobots int
}

func Cost(code string, nRobots int) int {
	cacheKey := costKey{code, nRobots}
	if val, ok := costCache.Load(cacheKey); ok {
		return val.(int)
	}
	result := 0
	code = "A" + code
	for i := 0; i < len(code)-1; i++ {
		// use i:i+1 to get a string instead of just i to get a rune
		result += KeyPadSequenceCost(code[i:i+1], code[i+1:i+2], nRobots)
	}

	costCache.Store(cacheKey, result)
	return result
}

func Complexity(code string, nRobots int) int {
	sequence := ""
	// Append A as start position to code
	code = "A" + code
	for i := 0; i < len(code)-1; i++ {
		sequence += DoorSequence(string(code[i]), string(code[i+1]))
	}

	// recursively find the cost of the shortest sequence
	cost := Cost(sequence, nRobots)
	// log.Println(sequence)
	num, err := strconv.Atoi(code[1 : len(code)-1])
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return cost * num
}

func First(codes []string) int {
	sum := 0
	for _, code := range codes {
		result := Complexity(code, 2)
		log.Println(code, result)
		sum += result
	}
	return sum
}

func Second(codes []string) int {
	sum := 0
	for _, code := range codes {
		result := Complexity(code, 25)
		log.Println(code, result)
		sum += result
	}
	return sum
}
