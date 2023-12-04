package day3

import (
	"reflect"
	"testing"
)

func Test_findNumbersInLine(t *testing.T) {
	type test map[string]struct {
		s        string
		want     []int
		hasError bool
	}
	tests := test{
		"Test single number": {
			s:        "617*......",
			want:     []int{617},
			hasError: false,
		},
		"Test two numbers": {
			s:        "617*...666...",
			want:     []int{617, 666},
			hasError: false,
		},
		"Test no numbers": {
			s:        "..*...&...",
			want:     []int{},
			hasError: true,
		},
		"Test many numbers": {
			s:        "730....138.30..455.....................589..",
			want:     []int{730, 138, 30, 455, 589},
			hasError: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := findNumbersInLine(tt.s)
			if !((err != nil) == tt.hasError) {
				t.Errorf("evaluateNumberMatch() = %v, want %v", err, tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evaluateNumberMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schematic_findNumAdjToSymbol(t *testing.T) {
	type test map[string]struct {
		s       schematic
		want    []int
		wantErr bool
	}
	tests := test{
		//"Single Line Test": {
		//	s: schematic{
		//		"617*.....12.",
		//	},
		//	want: []int{617},
		//},
		"Test example from the challenge": {
			s: schematic{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			want: []int{467, 35, 633, 617, 592, 755, 664, 598},
			//two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.s.findNumsAdjToSymbol(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumAdjToSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matchSymbol(t *testing.T) {
	type test map[string]struct {
		s    string
		want bool
	}
	tests := test{
		"matches $": {
			s:    "$",
			want: true,
		},
		"matches &": {
			s:    "&",
			want: true,
		},
		"does not match a .": {
			s:    ".",
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isSymbol(tt.s); got != tt.want {
				t.Errorf("isSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
