package twelve

import (
	"bufio"
	"os"
	"strings"
)

func Twelve() int {
	f, _ := os.Open("../inputs/twelvth_day_input.txt")

	defer f.Close()
	var garden [][]string
	scanner := bufio.NewScanner(f)
	garden = make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		gardenRow := make([]string, len(line))
		for i, c := range line {
			gardenRow[i] = strings.ToUpper(string(c))
		}
		garden = append(garden, gardenRow)
	}

	return CalculateFencePrice(garden)
}

type Location struct {
	x int
	y int
}


type Stack []Location

func (s *Stack) Push(location Location) {
	*s = append(*s, location)
}

func (s *Stack) Pop() Location {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

func CalculateFencePrice(garden [][]string) int {
	visitedLocations := make(map[Location]bool)
	price := 0

	for x := 0; x < len(garden); x++ {
		for y := 0; y < len(garden[0]); y++ {
			if visitedLocations[Location{x, y}] {
				continue

			}
			area, perimeter := countAreaAndPerimeter(garden, Location{x, y}, &visitedLocations)
			price += area * perimeter
		}
	}
	return price
}

func countAreaAndPerimeter(garden [][]string, startLocation Location, visitedLocations *map[Location]bool) (int, int) {
	perimeter := make(map[])
	area := 0
	perimeter := 0
	locations := Stack{startLocation}
	plant := garden[startLocation.x][startLocation.y]
	borderLocations := make(map[Location]bool)
	for len((locations)) > 0 {
		location := locations.Pop()
		if !checkLocation(garden, location, plant) {
			// Start scanning the border of the area


		} else if !(*visitedLocations)[location] {
			area++
			locations.Push(Location{location.x - 1, location.y})
			locations.Push(Location{location.x + 1, location.y})
			locations.Push(Location{location.x, location.y - 1})
			locations.Push(Location{location.x, location.y + 1})
			(*visitedLocations)[location] = true
		}

	}
	return area, perimeter
}

func checkLocation(garden [][]string, location Location, plant string) bool {
	if location.x < 0 || location.y < 0 || location.x >= len(garden) || location.y >= len(garden[0]) {
		return false
	} else if garden[location.x][location.y] != plant {
		return false
	}
	return true
}

type Scanner struct {
	left Location
	right Location
	direction Location
}

// Initializes the direction by looking at the starting location
func NewScanner(garden [][]string, location Location, target string) *Scanner {
	cross := [][]int{
		{location.x - 1, location.y},
		{location.x + 1, location.y},
		{location.x, location.y - 1},
		{location.x, location.y + 1},
	}
	for i := 0; i < len(cross); i++ {
		if checkLocation(garden, Location{cross[i][0], cross[i][1]}, target) {
			return &Scanner{
				left: location, 
				right: Location{cross[i][0], cross[i][1]}, 
				direction: Location{}},
			}
		}
	}
	return nil
}


// Scans the border of the area by dlf along individual borders by looking at a 2 wide cross
// It counts the number of turns necessary to get back to the starting location, the number of unique locations
// and the adds the unique locations to a map
func scanBorder(garden [][]string, startLocation Location, visitedLocations *map[Location]bool) {

}