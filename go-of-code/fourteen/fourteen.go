package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
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
	fmt.Println(Fourteen())

	elapsed := time.Since(start)
	logger.Printf("Day 14 took %s", elapsed)
}

func Fourteen() int {
	robots, err := loadRobots("../../inputs/14th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	middleSafeties := []int{}

	for seconds := 0; seconds < 100*103; seconds++ {
		bathroom := CreateBathroom(101, 103)

		bathroom = CalculatePositions(robots, bathroom, seconds)
		middleSafety := CalculateMiddleSafety(bathroom)
		middleSafeties = append(middleSafeties, middleSafety)
		if middleSafety > 250 {
			log.Println("Seconds: ", seconds)
			PrintBathroom(bathroom)
		}

	}
	slices.Sort(middleSafeties)
	log.Println(middleSafeties[0], middleSafeties[len(middleSafeties)/2], middleSafeties[len(middleSafeties)-1])
	bathroom := CreateBathroom(101, 103)
	bathroom = CalculatePositions(robots, bathroom, 100)

	//PrintBathroom(bathroom)

	return CalculateSafety(bathroom)
}

type Robot struct {
	x  int
	y  int
	vX int
	vY int
}

func loadRobots(filename string) ([]Robot, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	robots := make([]Robot, 0)
	var robot Robot

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}
		robot = Robot{}

		re := regexp.MustCompile(`(?:v|p)=(\-*\d+),(\-*\d+)`)
		matches := re.FindAllStringSubmatch(line, -1)

		if len(matches) != 2 {
			return nil, fmt.Errorf("invalid line at %q", line)
		}

		robot.x, _ = strconv.Atoi(matches[0][1])
		robot.y, _ = strconv.Atoi(matches[0][2])
		robot.vX, _ = strconv.Atoi(matches[1][1])
		robot.vY, _ = strconv.Atoi(matches[1][2])
		robots = append(robots, robot)

	}

	return robots, nil
}

func CreateBathroom(dx, dy int) *[][]int {
	bathroom := make([][]int, dy)
	for i := range bathroom {
		bathroom[i] = make([]int, dx)
	}
	return &bathroom
}

func PrintBathroom(bathroom *[][]int) {
	for _, row := range *bathroom {
		var sb strings.Builder
		for _, cell := range row {
			if cell == 0 {
				sb.WriteString(".")
			} else {
				sb.WriteString(strconv.Itoa(cell))
			}
		}
		log.Println(sb.String())
	}
}

func CalculatePositions(robots []Robot, bathroom *[][]int, seconds int) *[][]int {
	for _, robot := range robots {
		// Calculate the position of the robot in the bathroom with wrap around
		newX := (robot.x + robot.vX*seconds) % len((*bathroom)[0])
		if newX < 0 {
			newX += len((*bathroom)[0])
		}

		newY := (robot.y + robot.vY*seconds) % len(*bathroom)
		if newY < 0 {
			newY += len(*bathroom)
		}

		(*bathroom)[newY][newX] += 1
	}
	return bathroom
}

func CalculateSafety(bathroom *[][]int) int {
	// For each quadrant multiply all values larger than 0
	// and return the sum of all quadrants
	quadrants := [][]int{
		{0, len(*bathroom) / 2, 0, len((*bathroom)[0]) / 2},
		{0, len(*bathroom) / 2, len((*bathroom)[0])/2 + 1, len((*bathroom)[0])},
		{len(*bathroom)/2 + 1, len(*bathroom), 0, len((*bathroom)[0]) / 2},
		{len(*bathroom)/2 + 1, len(*bathroom), len((*bathroom)[0])/2 + 1, len((*bathroom)[0])},
	}
	result := 1
	for _, quadrant := range quadrants {
		sum := 0
		for i := quadrant[0]; i < quadrant[1]; i++ {
			for j := quadrant[2]; j < quadrant[3]; j++ {
				if (*bathroom)[i][j] > 0 {
					sum += (*bathroom)[i][j]
				}
			}
		}
		if sum > 0 {
			result *= sum
		}
	}
	return result
}

func CalculateMiddleSafety(bathroom *[][]int) int {
	sum := 0
	for i := len(*bathroom) / 4; i < 3*len(*bathroom)/4; i++ {
		for j := len((*bathroom)[0]) / 4; j < 3*len((*bathroom)[0])/4; j++ {
			sum += (*bathroom)[i][j]
		}
	}
	return sum
}
