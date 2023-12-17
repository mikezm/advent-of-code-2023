package day4

import (
	"errors"
	"fmt"
	"github.com/mikezm/advent-of-code-2023/read"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const INPUTS = "./day4/input.txt"

type Challenge struct{}

func (c Challenge) A() {
	ir, err := read.NewReader(INPUTS)

	if err != nil {
		fmt.Println("failed to read inputs")
		return
	}

	lines := ir.Lines()

	results := 0
	for _, line := range lines {
		c := readCardFromLine(line)
		results += c.calculatePoints()
	}

	fmt.Println("Results for Day 4 Part A")
	fmt.Println("Total points for all cards: ", results)
}

type cardCount struct {
	count int
	c     *card
}

type cardMap map[int]cardCount

func (cm cardMap) increment(c *card) {
	cc, ok := cm[c.num]
	if !ok {
		cm[c.num] = cardCount{1, c}
	} else {
		cm[c.num] = cardCount{cc.count + 1, c}
	}
}

func countCards(lines []string) int {
	var results int
	cm := make(cardMap)

	for lnIdx, line := range lines {
		c := readCardFromLine(line)
		numMatches := c.calculateMatches()
		cm.increment(&c)

		if numMatches > 0 {
			maxLn := int(math.Min(float64(len(lines)-1), float64(lnIdx+numMatches+1)))
			minLn := int(math.Min(float64(len(lines)-1), float64(lnIdx+1)))

			for i := minLn; i < maxLn; i++ {
				curCount := cm[i].count
				for j := 0; j < curCount; j++ {
					cc := readCardFromLine(lines[i])
					cm.increment(&cc)
				}

			}
		}
	}

	for _, cd := range cm {
		results += cd.count
	}

	return results
}

func (c Challenge) B() {
	ir, err := read.NewReader(INPUTS)

	if err != nil {
		fmt.Println("failed to read inputs")
		return
	}

	lines := ir.Lines()

	results := countCards(lines)

	fmt.Println("Results for Day 4 Part B")
	fmt.Println("Total points for all cards: ", results)
}

type card struct {
	num  int
	win  []int
	have []int
}

func (cd card) calculatePoints() int {
	matches := 0
	for _, v := range cd.have {
		if inArray(cd.win, v) {
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

func (cd card) calculateMatches() int {
	matches := 0
	for _, v := range cd.have {
		if inArray(cd.win, v) {
			matches += 1
		}
	}

	return matches
}

func readCardFromLine(l string) card {
	cardNum, remainder, sErr := split(l, `Card(\s+)(\d+): (.+)$`)
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

	if len(match) != 4 {
		return "", "", errors.New("could not find a match")
	}

	return match[2], match[3], nil
}

func findNumbersInLine(l string) []int {
	re := regexp.MustCompile(`(\d+)`)
	nums := re.FindAllStringSubmatch(l, -1)

	if nums == nil {
		return []int{}
	}

	var results []int
	for _, row := range nums {
		num, err := strconv.Atoi(row[0])
		if err != nil {
			return []int{}
		}

		results = append(results, num)
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
