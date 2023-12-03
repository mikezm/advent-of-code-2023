package day2

import (
	"errors"
	"fmt"
	"github.com/mikezm/advent-of-code-2023/read"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "./day2/input.txt"

type round map[string]int

func newRound() round {
	return round{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
}

func (r round) isPossible(tr round) bool {
	return r["red"] <= tr["red"] && r["green"] <= tr["green"] && r["blue"] <= tr["blue"]
}

type game struct {
	num    int
	rounds []round
	ln     string
}

func (g game) isPossible(tr round) bool {
	for _, r := range g.rounds {
		if !r.isPossible(tr) {
			return false
		}
	}
	return true
}

func (g game) findMinSet() round {
	nr := newRound()
	for _, r := range g.rounds {
		for _, c := range []string{"red", "green", "blue"} {
			if r[c] > nr[c] {
				nr[c] = r[c]
			}
		}
	}
	return nr
}

type Challenge struct{}

func (d Challenge) A() {
	test := round{"red": 12, "green": 13, "blue": 14}
	ir, err := read.NewReader(inputFile)

	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	var result int

	for _, l := range ir.Lines() {
		g, gErr := readGameFromString(l)
		if gErr != nil {
			fmt.Println("failed to execute readGameFromString()")
			return
		}

		if g.isPossible(test) {
			result += g.num
		}
	}

	fmt.Println("Total is: ", result)
}

func (d Challenge) B() {
	ir, err := read.NewReader(inputFile)

	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	var result int

	for _, l := range ir.Lines() {
		g, gErr := readGameFromString(l)
		if gErr != nil {
			fmt.Println("failed to execute readGameFromString()")
			return
		}

		ms := g.findMinSet()
		result += ms["red"] * ms["green"] * ms["blue"]
	}

	fmt.Println("Total is: ", result)
}

func split(s, reg string) (string, string, error) {
	re := regexp.MustCompile(reg)
	match := re.FindStringSubmatch(s)

	if len(match) != 3 {
		return "", "", errors.New("could not find a match")
	}

	return match[1], match[2], nil
}

func countRound(l string) (round, error) {
	nr := newRound()
	colorSplits := strings.Split(l, ",")

	for _, c := range colorSplits {
		amtMatch, color, cErr := split(c, `(\d+) (red|blue|green)`)
		if cErr != nil {
			return nr, errors.New("regex failed to find a match for a color")
		}
		amt, amErr := strconv.Atoi(amtMatch)
		if amErr != nil {
			return nr, errors.New("regex failed to find int")
		}
		nr[color] = amt
	}

	return nr, nil
}

func readGameFromString(l string) (game, error) {
	numMatch, rMatch, sErr := split(l, `Game (\d+): (.+)$`)
	if sErr != nil {
		return game{ln: l}, errors.New(fmt.Sprintf("could not find any games: %s", sErr))
	}

	num, err := strconv.Atoi(numMatch)
	if err != nil {
		return game{ln: l}, err
	}

	g := game{
		num:    num,
		rounds: nil,
		ln:     rMatch,
	}

	roundSplits := strings.Split(l, ";")
	for _, r := range roundSplits {
		nr, cErr := countRound(r)
		if cErr != nil {
			return game{ln: l}, cErr
		}

		g.rounds = append(g.rounds, nr)
	}

	return g, nil
}
