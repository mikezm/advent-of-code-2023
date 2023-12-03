package day2

import (
	"reflect"
	"testing"
)

func Test_countRound(t *testing.T) {

	tests := []struct {
		name    string
		ln      string
		want    round
		wantErr bool
	}{
		{
			name: "test all values are 9",
			ln:   "9 red, 9 green, 9 blue",
			want: round{
				"red":   9,
				"green": 9,
				"blue":  9,
			},
			wantErr: false,
		},
		{
			name: "test all random values",
			ln:   "2 red, 1 green, 4 blue",
			want: round{
				"red":   2,
				"green": 1,
				"blue":  4,
			},
			wantErr: false,
		},
		{
			name: "test some values",
			ln:   "2 red, 1 green",
			want: round{
				"red":   2,
				"green": 1,
				"blue":  0,
			},
			wantErr: false,
		},
		{
			name: "test some values",
			ln:   " red, 1en",
			want: round{
				"red":   0,
				"green": 0,
				"blue":  0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countRound(tt.ln)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGameRounds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countRound() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readGameFromString(t *testing.T) {
	tests := []struct {
		name, line string
		want       game
		wantErr    bool
	}{
		{
			name: "one game round",
			line: "Game 90: 8 red, 7 blue; 4 green, 3 red, 1 blue; 5 blue, 2 green",
			want: game{
				num: 90,
				rounds: []round{
					{"red": 8, "green": 0, "blue": 7},
					{"red": 3, "green": 4, "blue": 1},
					{"red": 0, "green": 2, "blue": 5},
				},
				ln: "8 red, 7 blue; 4 green, 3 red, 1 blue; 5 blue, 2 green",
			},
			wantErr: false,
		},
		{
			name: "one game round",
			line: "Game : 8 red, 7 blue; 4 green, 3 red, 1 blue; 5 blue, 2 green",
			want: game{
				num:    0,
				rounds: nil,
				ln:     "Game : 8 red, 7 blue; 4 green, 3 red, 1 blue; 5 blue, 2 green",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readGameFromString(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("readGameFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readGameFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round_isPossible(t *testing.T) {

	tests := []struct {
		name string
		r    round
		test round
		want bool
	}{
		{
			name: "possible round exact",
			r:    round{"red": 12, "green": 13, "blue": 14},
			test: round{"red": 12, "green": 13, "blue": 14},
			want: true,
		},
		{
			name: "possible round with 0",
			r:    round{"red": 12, "green": 13, "blue": 0},
			test: round{"red": 12, "green": 12, "blue": 0},
			want: false,
		},
		{
			name: "impossible round",
			r:    round{"red": 10, "green": 13, "blue": 14},
			test: round{"red": 12, "green": 12, "blue": 12},
			want: false,
		},
		{
			name: "possible round with 0s",
			r:    round{"red": 10, "green": 0, "blue": 0},
			test: round{"red": 12, "green": 12, "blue": 12},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.isPossible(tt.test); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_isPossible(t *testing.T) {

	tests := []struct {
		name string
		g    game
		test round
		want bool
	}{
		{
			name: "possible game",
			g: game{
				num: 1,
				rounds: []round{
					{"red": 4, "green": 0, "blue": 3},
					{"red": 1, "green": 2, "blue": 6},
					{"red": 0, "green": 0, "blue": 2},
				},
				ln: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			test: round{"red": 12, "green": 13, "blue": 14},
			want: true,
		},
		{
			name: "impossible game",
			g: game{
				num: 1,
				rounds: []round{
					{"red": 20, "green": 8, "blue": 14},
					{"red": 4, "green": 13, "blue": 5},
					{"red": 1, "green": 5, "blue": 0},
				},
				ln: "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			},
			test: round{"red": 12, "green": 13, "blue": 14},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.isPossible(tt.test); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_findMinSet(t *testing.T) {
	tests := []struct {
		name string
		line string
		want round
	}{
		{
			name: "example 1",
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: round{"red": 4, "green": 2, "blue": 6},
		},
		{
			name: "example 2",
			line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: round{"red": 1, "green": 3, "blue": 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, gErr := readGameFromString(tt.line)
			if gErr != nil {
				t.Errorf("failed running readGameFromString() with %v", gErr)
			}
			if got := g.findMinSet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
