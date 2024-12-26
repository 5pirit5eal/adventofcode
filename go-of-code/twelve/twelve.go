package twelve

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Twelve() (int, int) {
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

func CalculateFencePrice(garden [][]string) (int, int) {
	visitedLocations := make(map[Location]bool)
	price := 0
	priceDiscounted := 0

	for x := 0; x < len(garden); x++ {
		for y := 0; y < len(garden[0]); y++ {
			if visitedLocations[Location{x, y}] {
				continue

			}
			area, perimeter, uniqueBorders := countAreaAndPerimeter(garden, Location{x, y}, &visitedLocations)
			price += area * perimeter
			priceDiscounted += area * uniqueBorders
		}
	}
	return price, priceDiscounted
}

func countAreaAndPerimeter(garden [][]string, startLocation Location, visitedLocations *map[Location]bool) (int, int, int) {
	area := 0
	perimeter := 0
	uniqueBorders := 0
	locations := Stack{startLocation}
	plant := garden[startLocation.x][startLocation.y]
	scannedLocations := make(map[Location]bool)
	tempVisitedLocations := make(map[Location]bool)
	for len((locations)) > 0 {
		location := locations.Pop()
		if scannedLocations[location] {
			continue
		}

		if !checkLocation(garden, location, plant) {
			// Start scanning the border of the area
			scanner := NewScanner(garden, location, plant, &scannedLocations, &tempVisitedLocations)
			if !scanner.ScanBorder(garden, &locations) {
				log.Fatalf("Could not find border for location %v", location)
				break
			}
			uniqueBorders += scanner.turns

		} else if !tempVisitedLocations[location] {
			area++
			locations.Push(Location{location.x - 1, location.y})
			locations.Push(Location{location.x + 1, location.y})
			locations.Push(Location{location.x, location.y - 1})
			locations.Push(Location{location.x, location.y + 1})
			tempVisitedLocations[location] = true
		}
	}
	// join temp and visited locations
	for k, v := range tempVisitedLocations {
		(*visitedLocations)[k] = v
	}
	return area, perimeter, uniqueBorders
}

// checkLocation verifies if the given location is within the bounds of the garden and if the plant at that location matches the specified plant.
//
// Parameters:
// - garden: The 2D slice representing the garden.
// - location: The Location struct representing the current coordinates in the garden.
// - plant: The string representing the plant to be checked against.
//
// Returns:
// - A boolean indicating if the location is within bounds and the plant matches.
func checkLocation(garden [][]string, location Location, plant string) bool {
	// Check if the location is out of bounds
	if location.x < 0 || location.y < 0 || location.x >= len(garden) || location.y >= len(garden[0]) {
		return false
	}
	// Check if the plant at the location is not the same as the given plant
	if garden[location.x][location.y] != plant {
		return false
	}
	return true
}

type Scanner struct {
	left      Location
	right     Location
	direction Location
	turns     int
	scanned   *map[Location]bool
}

// Initializes the direction by looking at the starting location
func NewScanner(garden [][]string, location Location, target string, scanned *map[Location]bool, visitedLocations *map[Location]bool) *Scanner {
	cross := [][]int{
		{location.x - 1, location.y},
		{location.x + 1, location.y},
		{location.x, location.y - 1},
		{location.x, location.y + 1},
	}
	for i := 0; i < len(cross); i++ {
		if !(*visitedLocations)[Location{cross[i][0], cross[i][1]}] {
			continue
		}
		if checkLocation(garden, Location{cross[i][0], cross[i][1]}, target) {
			scanner := Scanner{
				left:    location,
				right:   Location{cross[i][0], cross[i][1]},
				scanned: scanned,
			}
			(*scanned)[location] = true
			scanner.SetDirection()
			return &scanner
		}
	}
	return nil
}

// Scans the border of the area by dlf along individual borders by looking at a 2 wide cross
// It counts the number of turns necessary to get back to the starting location, the number of unique locations
// and the adds the unique locations to a map
func (s *Scanner) ScanBorder(garden [][]string, locations *Stack) bool {
	if s == nil {
		return false
	}
	goal := Scanner{
		left:  s.left,
		right: s.right,
	}
	rightPlant := garden[s.right.x][s.right.y]

	// Do first step away from starting location
	s.Step(garden, locations, rightPlant)

	// Add safety mechanism to stop when steps are above 1000000
	steps := 0
	for goal.left != (*s).left || goal.right != (*s).right {
		if steps > 1000000 {
			return false
		}
		s.Step(garden, locations, rightPlant)
		steps++
	}
	return true
}

func (s *Scanner) Step(garden [][]string, locations *Stack, plant string) {
	// check if the location is out of bounds
	aheadLeft := Location{s.left.x + s.direction.x, s.left.y + s.direction.y}
	aheadRight := Location{s.right.x + s.direction.x, s.right.y + s.direction.y}

	// Expect left not to match, therefore turn when inside area
	if !checkLocation(garden, aheadRight, plant) || (checkLocation(garden, aheadLeft, plant) && !checkLocation(garden, aheadRight, plant)) {
		s.turns++
		(*s).left = aheadRight
		s.SetDirection()
		(*locations).Push(s.right)
	} else if checkLocation(garden, aheadLeft, plant) {
		s.turns++
		(*s).right = aheadLeft
		s.SetDirection()
		// Expect right to match, therfore turn when not inside area
	} else {
		(*s).left = aheadLeft
		(*s).right = aheadRight
		(*locations).Push(s.right)
	}
	(*s.scanned)[s.left] = true
}

func (s *Scanner) SetDirection() {
	(*s).direction = Location{(*s).right.y - (*s).left.y, (*s).left.x - (*s).right.x}
}
