package main

import (
	"log"
	"testing"
)

func TestGenerateSecret(t *testing.T) {
	inputs := [][]int{
		{1, 8685429},
		{10, 4700978},
		{100, 15273692},
		{2024, 8667524},
	}

	for _, input := range inputs {
		got := generateSecret(input[0], 2000)
		if got != input[1] {
			t.Errorf("generateSecret(%d) = %d, want %d", input[0], got, input[1])
		}
	}
}

func TestGenerateSecretSequence(t *testing.T) {
	inputs := []int{
		123,
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}

	for i := 0; i < len(inputs)-1; i++ {
		got := generateSecret(inputs[i], 1)
		log.Printf("generateSecret(%b) = %b, want %b", inputs[i], got, inputs[i+1])
		if got != inputs[i+1] {
			t.Errorf("generateSecret(%d) = %d, want %d", inputs[0], got, inputs[i+1])
		}
	}
}

func TestOneDigit(t *testing.T) {
	inputs := [][]int{
		{123, 3},
		{15887950, 0},
		{16495136, 6},
	}

	for _, input := range inputs {
		got := getOneDigit(input[0])
		if got != input[1] {
			t.Errorf("oneDigit(%d) = %d, want %d", input[0], got, input[1])
		}
	}
}
