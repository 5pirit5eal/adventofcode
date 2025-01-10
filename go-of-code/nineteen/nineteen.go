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
	logger.Printf("Starting day 19")

	start := time.Now()
	log.Println(Nineteen())
	elapsed := time.Since(start)
	logger.Printf("Day 19 took %s", elapsed)
}

func loadInput(filename string) ([]string, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	substrings := make([]string, 0)
	sequences := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.Contains(line, ",") {
			for _, s := range strings.Split(line, ",") {
				substrings = append(substrings, strings.TrimSpace(s))
			}
		} else {
			sequences = append(sequences, line)
		}

	}

	return sequences, substrings, nil
}

func Nineteen() (int, int) {
	sequences, substrings, err := loadInput("../../inputs/19th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sequences)
	log.Println(substrings)

	result := 0
	result2 := 0

	//var composed int
	var composedWithCount int

	for _, sequence := range sequences {
		// composed = composable(sequence, substrings)
		composedWithCount = composableWithCount(sequence, substrings)
		if composedWithCount > 0 {
			result += 1
		}
		result2 += composedWithCount
	}
	return result, result2
}

// Checks if the sequence can be composed of the substrings
func composable(sequence string, substrings []string) int {
	d := make([]bool, len(sequence)+1)
	d[0] = true

	for i := range d {
		for _, sub := range substrings {
			if i >= len(sub) && sequence[i-len(sub):i] == sub {
				d[i] = d[i] || d[i-len(sub)]
				if d[i] {
					break
				}
			}
		}
	}

	if d[len(d)-1] {
		return 1
	} else {
		return 0
	}
}

// Checks if the sequence can be composed of the substrings and returns the number of combinations
func composableWithCount(sequence string, substrings []string) int {
	d := make([]int, len(sequence)+1)
	d[0] = 1

	for i := range d {
		for _, sub := range substrings {
			if i >= len(sub) && sequence[i-len(sub):i] == sub {
				d[i] = d[i] + d[i-len(sub)]
			}
		}
	}

	return d[len(d)-1]
}
