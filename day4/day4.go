package day4

import (
	"errors"
	"fmt"
	"github.com/mikezm/advent-of-code-2023/read"
	"regexp"
	"strconv"
	"strings"
)

const INPUTS = "./day4/input.txt"

type Challenge struct{}

func (d Challenge) A() {
	ir, err := read.NewReader(INPUTS)

	if err != nil {
		fmt.Println("failed to read inputs")
		return
	}

	results := 0
	for _, line := range ir.Lines() {
		c := readCardFromLine(line)
		results += c.calculatePoints()
	}

	fmt.Println("Results for Day 4 PaRT A")
	fmt.Println("Total points for all cards: ", results)
}

func (d Challenge) B() {
	fmt.Println("function B() not yet implemented")
}

type card struct {
	num  int
	win  []int
	have []int
}

func (c card) calculatePoints() int {
	matches := 0
	for _, v := range c.have {
		if inArray(c.win, v) {
			matches += 1
		}
	}

	if matches == 0 {
		return 0
	}

	points := 1
	for i := 0; i < matches-1; i++ {
		points = points * 2
	}

	return points
}

func readCardFromLine(l string) card {
	cardNum, remainder, sErr := split(l, `Card (\d+): (.+)$`)
	if sErr != nil {
		return card{}
	}

	num, err := strconv.Atoi(cardNum)
	if err != nil {
		return card{}
	}

	c := card{
		num: num,
	}

	list := strings.Split(remainder, "|")
	if len(list) != 2 {
		return c
	}

	c.win = findNumbersInLine(list[0])
	c.have = findNumbersInLine(list[1])

	return c
}

func split(s, reg string) (string, string, error) {
	re := regexp.MustCompile(reg)
	match := re.FindStringSubmatch(s)

	if len(match) != 3 {
		return "", "", errors.New("could not find a match")
	}

	return match[1], match[2], nil
}

func findNumbersInLine(l string) []int {
	re := regexp.MustCompile(`(\d+)`)
	nums := re.FindAllStringSubmatch(l, -1)

	if nums == nil {
		return []int{}
	}

	var results []int
	for _, row := range nums {
		for _, n := range row {
			num, err := strconv.Atoi(n)
			if err != nil {
				return []int{}
			}

			if !inArray(results, num) {
				results = append(results, num)
			}

		}
	}

	return results
}

func inArray(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}
