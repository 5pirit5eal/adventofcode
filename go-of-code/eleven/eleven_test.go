package eleven

import (
	"math/big"
	"testing"
)

func TestBlinkSteps(t *testing.T) {
	stones := [][]*big.Int{
		{big.NewInt(125), big.NewInt(17)},
		{big.NewInt(253000), big.NewInt(1), big.NewInt(7)},
		{big.NewInt(253), big.NewInt(0), big.NewInt(2024), big.NewInt(14168)},
		{big.NewInt(512072), big.NewInt(1), big.NewInt(20), big.NewInt(24), big.NewInt(28676032)},
		{big.NewInt(512), big.NewInt(72), big.NewInt(2024), big.NewInt(2), big.NewInt(0), big.NewInt(2), big.NewInt(4), big.NewInt(2867), big.NewInt(6032)},
		{big.NewInt(1036288), big.NewInt(7), big.NewInt(2), big.NewInt(20), big.NewInt(24), big.NewInt(4048), big.NewInt(1), big.NewInt(4048), big.NewInt(8096), big.NewInt(28), big.NewInt(67), big.NewInt(60), big.NewInt(32)},
		{big.NewInt(2097446912), big.NewInt(14168), big.NewInt(4048), big.NewInt(2), big.NewInt(0), big.NewInt(2), big.NewInt(4), big.NewInt(40), big.NewInt(48), big.NewInt(2024), big.NewInt(40), big.NewInt(48), big.NewInt(80), big.NewInt(96), big.NewInt(2), big.NewInt(8), big.NewInt(6), big.NewInt(7), big.NewInt(6), big.NewInt(0), big.NewInt(3), big.NewInt(2)},
	}
	for i := 0; i < len(stones)-1; i++ {
		result := Blinker(stones[i], 1).Int64()
		if result != int64(len(stones[i+1])) {
			t.Errorf("Expected length %d, got %d", len(stones[i+1]), result)
		}
	}

}

func TestBlink(t *testing.T) {
	result := Blinker([]*big.Int{big.NewInt(125), big.NewInt(17)}, 25)
	expected := big.NewInt(55312)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSplitBigInt(t *testing.T) {
	tests := []struct {
		name  string
		input *big.Int
		want  []*big.Int
	}{
		{
			name:  "small number",
			input: big.NewInt(1234),
			want:  []*big.Int{big.NewInt(12), big.NewInt(34)},
		},
		{
			name:  "single digit",
			input: big.NewInt(15151515),
			want:  []*big.Int{big.NewInt(1515), big.NewInt(1515)},
		},
		{
			name:  "zero",
			input: big.NewInt(0),
			want:  []*big.Int{big.NewInt(0), big.NewInt(0)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitBigInt(tt.input)
			if len(got) != 2 {
				t.Errorf("splitBigInt() = %v, want 2 elements", got)
			}
			if !(got[0].Cmp(tt.want[0]) == 0) || !(got[1].Cmp(tt.want[1]) == 0) {
				t.Errorf("splitBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
