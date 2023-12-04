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
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.c.calculatePoints(); got != tt.want {
				t.Errorf("calculatePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
