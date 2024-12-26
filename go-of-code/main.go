package main

import (
	"fmt"
	"log"
	"os"
	"time"

	eleven "advent-of-code/go-of-code/eleven"
	ten "advent-of-code/go-of-code/ten"
	"advent-of-code/go-of-code/twelve"
)

func setupLogger(day int) *log.Logger {
	log.SetPrefix(string(day))
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	return logger
}

func main() {
	logger := setupLogger(1)
	// day, err := strconv.Atoi(os.Args[1])
	day := 12
	var err error = nil
	if err != nil {
		log.Fatalf("Could not parse day %q as int", os.Args[1])
	}
	logger.Printf("Starting day %d", day)

	start := time.Now()
	switch day {
	// case 1:
	// 	one()
	// case 2:
	// 	two()
	// case 3:
	// 	three()
	// case 4:
	// 	four()
	// case 5:
	// 	five()
	// case 6:
	// 	six()
	// case 7:
	// 	seven()
	// case 8:
	// 	eight()
	// case 9:
	// 	nine()
	case 10:
		fmt.Println(ten.Ten())
	case 11:
		fmt.Println(eleven.Eleven())
	case 12:
		fmt.Println(twelve.Twelve())
	default:
		logger.Fatalf("Not implemented yet.")
	}
	elapsed := time.Since(start)
	logger.Printf("Day %d took %s", day, elapsed)
}
