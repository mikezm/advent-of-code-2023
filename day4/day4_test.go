package day4

import (
	"reflect"
	"testing"
)

func Test_readCardFromLine(t *testing.T) {
	type test map[string]struct {
		line string
		want card
	}

	tests := test{
		"first test": {
			line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: card{
				num:  1,
				win:  []int{41, 48, 83, 86, 17},
				have: []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
		},
		"long test": {
			line: "Card  88: 47 49 95 31 36 53 37 86 92 42 | 58 22  6 14 62 50 93 23 43 11 90 67 60 56 40 81 75 91  2 45 65 25 69  1  5",
			want: card{
				num:  88,
				win:  []int{47, 49, 95, 31, 36, 53, 37, 86, 92, 42},
				have: []int{58, 22, 6, 14, 62, 50, 93, 23, 43, 11, 90, 67, 60, 56, 40, 81, 75, 91, 2, 45, 65, 25, 69, 1, 5},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := readCardFromLine(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCardFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_calculatePoints(t *testing.T) {
	type test map[string]struct {
		c    card
		want int
	}
	tests := test{
		"test one": {
			c: card{
				num:  1,
				win:  []int{41, 48, 83, 86, 17},
				have: []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			want: 8,
		},
		"test two": {
			c: card{
				num:  2,
				win:  []int{13, 32, 20, 16, 61},
				have: []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			want: 2,
		},
		"test four matches": {
			c: card{
				num:  2,
				win:  []int{13, 32, 20, 16, 61},
				have: []int{61, 13, 68, 82, 17, 32, 24, 20},
			},
			want: 8,
		},
		"test no matches": {
			c: card{
				num:  2,
				win:  []int{13, 32, 20, 16, 61},
				have: []int{2, 5, 68, 82, 17, 11, 24, 66},
			},
			want: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.c.calculatePoints(); got != tt.want {
				t.Errorf("calculatePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countCards(t *testing.T) {
	type test map[string]struct {
		lines []string
		want  int
	}
	tests := test{
		"test from AoC example": {
			lines: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			want: 30,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := countCards(tt.lines); got != tt.want {
				t.Errorf("countCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
