package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("Missing input file")
	}
	f, err := os.Open(flag.Arg(0))
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
