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
			hasError: false,
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
		"Single Line Test": {
			s: schematic{
				[]rune("617*.....12."),
			},
			want: []int{617},
		},
		"Test example from the challenge": {
			s: schematic{
				[]rune("467..114.."),
				[]rune("...*......"),
				[]rune("..35..633."),
				[]rune("......#..."),
				[]rune("617*......"),
				[]rune(".....+.58."),
				[]rune("..592....."),
				[]rune("......755."),
				[]rune("...$.*...."),
				[]rune(".664.598.."),
			},
			want: []int{467, 35, 633, 617, 592, 755, 664, 598},
			//two numbers are not part numbers because they are not adjacent to a symbolLoc: 114 (top right) and 58
		},
		"Bigger Test": {
			s: schematic{
				[]rune("..172..............................454..46.......507..........809......923.778..................793..............137.............238........"),
				[]rune("............/.........712........=.......*................515.*...........*.......690.........../..........658.........=.........*.........."),
				[]rune(".........823.835........%.........710.....749........134..%............................#812...&.....925.../..........276.......386.........."),
				[]rune("519..................13......341.................481....=.....$............-.......211.......92.......*....................................*"),
				[]rune("............832*105..-........$..................*.........797.....535..932.........*....152...........123.........678.540...........-...6.7"),
			},
			want: []int{46, 809, 923, 778, 793, 238, 712, 515, 658, 823, 835, 710, 749, 134, 812, 925, 276, 386, 13, 341, 481, 211, 92, 832, 105, 797, 932, 123, 7},
			//two numbers are not part numbers because they are not adjacent to a symbolLoc: 114 (top right) and 58
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.s.findAdj(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumAdjToSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matchSymbol(t *testing.T) {
	type test map[string]struct {
		s    rune
		want bool
	}
	tests := test{
		"matches $": {
			s:    '$',
			want: true,
		},
		"matches &": {
			s:    '&',
			want: true,
		},
		"matches a .": {
			s:    '.',
			want: true,
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

func Test_schematic_findSymbolLocations(t *testing.T) {
	type test map[string]struct {
		s       schematic
		want    []symbolLoc
		wantErr bool
	}
	tests := test{
		"Single Line Test": {
			s: schematic{
				[]rune("$..*.....12."),
			},
			want: []symbolLoc{
				{0, 0},
				{0, 3},
			},
		},
		"several lines": {
			s: schematic{
				[]rune("467..114.."),
				[]rune("...*......"),
				[]rune("..35..633."),
				[]rune("......#..."),
				[]rune("617*......"),
				[]rune(".......58*"),
			},
			want: []symbolLoc{
				{1, 3},
				{3, 6},
				{4, 3},
				{5, 9},
			},
			//two numbers are not part numbers because they are not adjacent to a symbolLoc: 114 (top right) and 58
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.s.findSymbolLocations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSymbolLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schematic_findNumberLocations(t *testing.T) {
	type test map[string]struct {
		s       schematic
		want    []numberLoc
		wantErr bool
	}
	tests := test{
		"Single Line Test": {
			s: schematic{
				[]rune("$..*...12."),
			},
			want: []numberLoc{
				{0, 7, 8, 12},
			},
		},
		"several lines": {
			s: schematic{
				[]rune("467..114.."),
				[]rune("...*......"),
				[]rune("..35..633."),
				[]rune("......#..."),
				[]rune("617*......"),
				[]rune("........58"),
			},
			want: []numberLoc{
				{0, 0, 2, 467},
				{0, 5, 7, 114},
				{2, 2, 3, 35},
				{2, 6, 8, 633},
				{4, 0, 2, 617},
				{5, 8, 9, 58},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.s.findNumberLocations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumberLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schematic_findGearRatios(t *testing.T) {
	type test map[string]struct {
		s    schematic
		want []int
	}
	tests := test{
		"Single Line Test": {
			s: schematic{
				[]rune("617*12........"),
			},
			want: []int{7404},
		},
		"Test example from the challenge": {
			s: schematic{
				[]rune("467..114.."),
				[]rune("...*......"),
				[]rune("..35..633."),
				[]rune("......#..."),
				[]rune("617*......"),
				[]rune(".....+.58."),
				[]rune("..592....."),
				[]rune("......755."),
				[]rune("...$.*...."),
				[]rune(".664.598.."),
			},
			want: []int{16345, 451490},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.s.findGearRatios(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumAdjToSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAdjacent(t *testing.T) {
	type test map[string]struct {
		num  numberLoc
		sym  symbolLoc
		want bool
	}
	tests := test{
		"same line": {
			num:  numberLoc{0, 0, 2, 612},
			sym:  symbolLoc{0, 3},
			want: true,
		},
		"symbol diagonal below right": {
			num:  numberLoc{0, 0, 2, 467},
			sym:  symbolLoc{1, 3},
			want: true,
		},
		"symbol diagonal above right": {
			num:  numberLoc{2, 2, 3, 35},
			sym:  symbolLoc{1, 3},
			want: true,
		},
		"diagonal above right": {
			num:  numberLoc{2, 2, 3, 35},
			sym:  symbolLoc{1, 3},
			want: true,
		},
		"diagonal below too far lef": {
			num:  numberLoc{0, 5, 7, 114},
			sym:  symbolLoc{1, 3},
			want: false,
		},
		//	"Test example from the challenge": {
		//		s: schematic{
		//			[]rune("467..114.."),
		//			[]rune("...*......"),
		//			[]rune("..35..633."),
		//			[]rune("......#..."),
		//			[]rune("617*......"),
		//			[]rune(".....+.58."),
		//			[]rune("..592....."),
		//			[]rune("......755."),
		//			[]rune("...$.*...."),
		//			[]rune(".664.598.."),
		//		},
		//		want: []int{16345, 467835},
		//	},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isAdjacent(tt.num, tt.sym); got != tt.want {
				t.Errorf("isAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
