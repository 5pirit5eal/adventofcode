package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 17")

	start := time.Now()
	log.Println(Seventeen())
	elapsed := time.Since(start)
	logger.Printf("Day 17 took %s", elapsed)
}

var A, B, C int

func loadInput(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	program := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", -1)
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid line at %q", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Register A":
			A, _ = strconv.Atoi(value)
		case "Register B":
			B, _ = strconv.Atoi(value)
		case "Register C":
			C, _ = strconv.Atoi(value)
		case "Program":
			programParts := strings.Split(value, ",")
			for _, p := range programParts {
				i, err := strconv.Atoi(p)
				if err != nil {
					return nil, fmt.Errorf("Invalid integer %q: %w", p, err)
				}
				program = append(program, i)
			}
		default:
			return nil, fmt.Errorf("Invalid key at %q", key)
		}
	}

	return program, nil
}

func Seventeen() (string, int) {
	program, err := loadInput("../../inputs/17th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(A, B, C)

	result := ""

	i := 0
	for {
		if i >= len(program) {
			break
		}

		switch program[i] {
		case 0:
			adv(program[i+1])
		case 1:
			bxl(program[i+1])
		case 2:
			bst(program[i+1])
		case 3:
			j, jump := jnz(program[i+1])
			if jump {
				i = j - 2
			}
		case 4:
			bxc()
		case 5:
			result += strconv.Itoa(out(program[i+1])) + ","
		case 6:
			bdv(program[i+1])
		case 7:
			cdv(program[i+1])
		default:
			log.Fatalf("Unknown instruction %d", program[i])
		}
		i += 2
	}

	log.Println(A, B, C)
	return result, 0
}

func combo(n int) int {
	switch n {
	case 0, 1, 2, 3:
		return n
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		log.Fatalf("Unknown combo %d", n)
		return 0
	}
}

func adv(n int) {
	A = A >> uint(combo(n))
}

// Bitwise XOR of register B with n
func bxl(n int) {
	B = B ^ n
}

func bst(n int) {
	B = combo(n) % 8
}

func jnz(n int) (int, bool) {
	if A == 0 {
		return 0, false
	} else {
		return n, true
	}
}

func bxc() {
	B = B ^ C
}

func out(n int) int {
	return combo(n) % 8
}

func bdv(n int) {
	B = A >> combo(n)
}

func cdv(n int) {
	C = A >> combo(n)
}
