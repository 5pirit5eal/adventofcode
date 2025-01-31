package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 25")

	start := time.Now()
	log.Println(TwentyFive())
	elapsed := time.Since(start)
	logger.Printf("Day 25 took %s", elapsed)
}

func loadInput(filename string) (map[string][]Schema, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	schemas := map[string][]Schema{"key": make([]Schema, 0), "lock": make([]Schema, 0)}
	s := make(Schema, 5)
	var c string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			schemas[c] = append(schemas[c], s)
			s = make(Schema, 5)
			c = ""
			continue
		}

		for i, char := range line {
			if char == '#' {
				s[i]++
			}
		}
		if c == "" {
			fr := 0
			for _, v := range s {
				fr += v
			}
			if fr == 5 {
				c = "lock"
			} else {
				c = "key"
			}
		}
	}
	schemas[c] = append(schemas[c], s)

	return schemas, nil
}

type Schema []int

func Compare(key, lock Schema) bool {
	for i := range key {
		if (8-lock[i])-key[i] <= 0 {
			return false
		}
	}
	return true
}

func TwentyFive() (int, string) {
	schemas, err := loadInput("../../inputs/25th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return First(schemas), Second()
}

func First(schemas map[string][]Schema) int {
	sum := 0

	for _, s := range schemas["lock"] {
		for _, k := range schemas["key"] {
			if Compare(k, s) {
				sum++
			}
		}
	}
	return sum
}

func Second() string {
	return ""
}
