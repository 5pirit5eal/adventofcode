package eleven

import (
	"bufio"
	"log"
	"os"
)

func eleven() {
	f, err := os.Open("inputs/eleventh_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		log.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
