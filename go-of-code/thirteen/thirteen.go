package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.SetPrefix("13:")
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 13")

	start := time.Now()
	fmt.Println(Thirteen())

	elapsed := time.Since(start)
	logger.Printf("Day 13 took %s", elapsed)
}

func Thirteen() (int, int) {
	configs, err := loadConfigs("../../inputs/13th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(configs)
	return CalculatePrices(configs, 0), CalculatePrices(configs, 10000000000000)
}

type Location struct {
	x int
	y int
}

type Configuration struct {
	ButtonA Location
	ButtonB Location
	Price   Location
}

func loadConfigs(filename string) ([]Configuration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	configs := make([]Configuration, 0)
	config := Configuration{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			configs = append(configs, config)
			config = Configuration{}
			continue
		}
		parts := strings.SplitN(line, ":", -1)
		if len(parts) > 2 {
			return nil, fmt.Errorf("Invalid line at %q", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		re := regexp.MustCompile("(?:X|Y)((?:\\+|\\-)\\d+)|(?:\\=)(\\d+)")
		matches := re.FindAllStringSubmatch(value, -1)

		if len(matches) != 2 {
			return nil, fmt.Errorf("Invalid line at %q", line)
		}

		switch key {
		case "Button A":
			config.ButtonA.x, _ = strconv.Atoi(matches[0][1])
			config.ButtonA.y, _ = strconv.Atoi(matches[1][1])
		case "Button B":
			config.ButtonB.x, _ = strconv.Atoi(matches[0][1])
			config.ButtonB.y, _ = strconv.Atoi(matches[1][1])
		case "Prize":
			config.Price.x, _ = strconv.Atoi(matches[0][2])
			config.Price.y, _ = strconv.Atoi(matches[1][2])
		default:
			return nil, fmt.Errorf("Invalid key at %q", key)
		}
	}
	configs = append(configs, config)
	return configs, nil
}

func (c *Configuration) CalculateOptimalPrice() int {
	switch {
	case c.ButtonA.x > c.Price.x && c.ButtonB.x > c.Price.x:
		return -1
	case c.ButtonA.y > c.Price.y && c.ButtonB.y > c.Price.y:
		return -1
	case c.ButtonA.x == c.Price.x && c.ButtonA.y == c.Price.y:
		return 3
	case c.ButtonB.x == c.Price.x && c.ButtonB.y == c.Price.y:
		return 1
	}

	// Calculate the determinantes from the equation
	d1 := c.Price.x*c.ButtonB.y - c.ButtonB.x*c.Price.y
	d2 := c.ButtonA.x*c.Price.y - c.ButtonA.y*c.Price.x
	dp := c.ButtonA.x*c.ButtonB.y - c.ButtonB.x*c.ButtonA.y
	if d1%dp != 0 || d2%dp != 0 {
		return -1
	}
	a := d1 / dp
	b := d2 / dp

	// Check if a and b are integers
	if (c.Price.y == a*c.ButtonA.y+b*c.ButtonB.y) && (c.Price.x == a*c.ButtonA.x+b*c.ButtonB.x) && a >= 0 && b >= 0 {
		//log.Println("Unique Solution: ", a, b)
		return 3*a + b
	} else {
		//log.Println(a, b, "Equate to X:", c.Price.x, a*c.ButtonA.x+b*c.ButtonB.x, "Y:", c.Price.y, a*c.ButtonA.y+b*c.ButtonB.y)
		return -2.0
	}

}

func CalculatePrices(configs []Configuration, correction int) int {
	total := 0
	for _, config := range configs {
		config.Price.x += correction
		config.Price.y += correction
		optimalPrice := config.CalculateOptimalPrice()
		if optimalPrice > 0 {
			total += optimalPrice
		}
	}
	return total
}
