package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 23")

	start := time.Now()
	log.Println(TwentyThree())
	elapsed := time.Since(start)
	logger.Printf("Day 23 took %s", elapsed)
}

func loadInput(filename string) (Network, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nums := make(Network, 0)

	for scanner.Scan() {
		num_pair := strings.Split(scanner.Text(), "-")

		if len(num_pair) != 2 {
			continue
		}

		left, right := num_pair[0], num_pair[1]

		if _, ok := nums[left]; !ok {
			nums[left] = make(map[string]bool, 0)
		}

		if _, ok := nums[right]; !ok {
			nums[right] = make(map[string]bool, 0)
		}

		nums[left][right] = true
		nums[right][left] = true
	}

	return nums, nil
}

func TwentyThree() (int, string) {
	nums, err := loadInput("../../inputs/23th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return First(nums), Second(nums)
}

type Network map[string]map[string]bool

// Find set of three connected computers by their id
func First(net Network) int {
	intercon := make(map[string]bool, 0)

	for a, cona := range net {
		for b := range cona {
			if a == b {
				continue
			}

			for ca := range cona {
				if _, ok := net[b][ca]; ok && ca != a && ca != b {
					for _, v := range []string{a, b, ca} {
						if strings.HasPrefix(v, "t") {
							intercon[CreateKey([]string{a, b, ca})] = true
							break
						}
					}
				}
			}
		}
	}

	return len(intercon)
}

func CreateKey(s []string) string {
	slices.Sort(s)
	return strings.Join(s, ",")
}

func (n Network) GatherPartners(partners []string, candidate string) []string {
	if slices.Contains(partners, candidate) {
		return partners
	}
	// Check if candidate is connected to all partners
	for _, partner := range partners {
		if _, ok := n[partner][candidate]; !ok {
			return partners
		}
	}
	partners = append(partners, candidate)

	for newCandidate := range n[candidate] {
		if !slices.Contains(partners, newCandidate) {
			partners = n.GatherPartners(partners, newCandidate)
		}
	}
	return partners
}

// Find the largest group of connected computers
func Second(net Network) string {
	intercon := map[string]int{}

	for a, cona := range net {
		for b := range cona {
			if a == b {
				continue
			}
			partners := net.GatherPartners([]string{a}, b)

			key := CreateKey(partners)
			intercon[key] = len(partners)
		}
	}

	max := 0
	key := ""
	for k, v := range intercon {
		if v > max {
			max = v
			key = k
		}
	}

	return key
}
