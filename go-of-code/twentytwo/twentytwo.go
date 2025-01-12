package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(0)
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.Printf("Starting day 22")

	start := time.Now()
	log.Println(TwentyTwo())
	elapsed := time.Since(start)
	logger.Printf("Day 22 took %s", elapsed)
}

func loadInput(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nums := make([]int, 0)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return nil, fmt.Errorf("invalid line at %q", scanner.Text())
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func TwentyTwo() (int, int) {
	nums, err := loadInput("../../inputs/22th_day_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return First(nums), Second(nums)
}

func generateSecret(num, iter int, ch chan int) {
	defer close(ch)
	for range iter {
		// times 64 moves bits by 6 to the left
		// % removes bits after 25
		num = (num ^ (num << 6)) & 0xFFFFFF
		// move 5 digits to the right
		num = (num ^ (num >> 5)) & 0xFFFFFF
		// move 11 digits to the left
		num = (num ^ (num << 11)) & 0xFFFFFF
	}
	ch <- num
}

func First(nums []int) int {
	sum := 0
	channels := make([]chan int, len(nums))
	for i, num := range nums {
		ch := make(chan int)
		channels[i] = ch
		go generateSecret(num, 2000, ch)
		// sum += <-ch
	}
	for _, ch := range channels {
		sum += <-ch
	}
	return sum
}

type PriceSequence struct {
	pos0, pos1, pos2, pos3 int
}

func generateSecretIter(num, iter int) <-chan int {
	nums := make(chan int, iter)

	go func() {
		defer close(nums)
		for range iter {
			// times 64 moves bits by 6 to the left
			// % removes bits after 25
			num = (num ^ (num << 6)) & 0xFFFFFF
			// move 5 digits to the right
			num = (num ^ (num >> 5)) & 0xFFFFFF
			// move 11 digits to the left
			num = (num ^ (num << 11)) & 0xFFFFFF

			nums <- num % 10
		}
	}()

	return nums
}

// Calculates the price sequences and their prices for a buyer
func BuyerPriceSequence(num int, ch chan map[PriceSequence]int) {
	defer close(ch)
	seq := make(map[PriceSequence]int)
	var num0, num1, num2, num3, num4 int
	i := 0
	for v := range generateSecretIter(num, 2000) {
		i++
		num0 = num1
		num1 = num2
		num2 = num3
		num3 = num4
		num4 = v
		if i < 4 {
			continue
		}
		pSeq := PriceSequence{num1 - num0, num2 - num1, num3 - num2, num4 - num3}

		_, ok := seq[pSeq]
		if !ok {
			seq[pSeq] = v
		}
	}
	ch <- seq
}

// Calculates the optimal PriceSequence so that the sum of all prices is maximized,
// then returns the sum of all prices for this sequence.
// The price sequence is the same for all buyers.
func Second(nums []int) int {
	channels := make([]chan map[PriceSequence]int, len(nums))
	for i, num := range nums {
		ch := make(chan map[PriceSequence]int)
		go BuyerPriceSequence(num, ch)
		channels[i] = ch
	}

	// Create a set of all sequences
	seqs := make(map[PriceSequence]int)
	for _, ch := range channels {
		for seq, price := range <-ch {
			seqs[seq] += price
		}
	}

	// Find the sequence with the highest price
	maxPrice := 0
	for price := range maps.Values(seqs) {
		if price > maxPrice {
			maxPrice = price
		}
	}

	return maxPrice
}
