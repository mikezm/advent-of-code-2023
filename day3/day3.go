package day3

import (
	"errors"
	"fmt"
	"github.com/mikezm/advent-of-code-2023/read"
	"math"
	"regexp"
	"strconv"
	"unicode"
)

const INPUTS = "./day3/input.txt"

type Challenge struct{}

func (d Challenge) A() {
	s, err := readSchematic()
	if err != nil {
		fmt.Println("error reading input file")
	}

	nums := s.findAdj()
	results := 0

	for _, n := range nums {
		results += n
	}

	fmt.Println("Results for Day 3 Part A: ", results)

}

func (d Challenge) B() {
	s, err := readSchematic()
	if err != nil {
		fmt.Println("error reading input file")
	}

	nums := s.findGearRatios()
	results := 0

	for _, n := range nums {
		results += n
	}

	fmt.Println("Results for Day 3 Part B: ", results)
}

type schematic [][]rune

func readSchematic() (schematic, error) {
	ir, err := read.NewReader(INPUTS)

	if err != nil {
		return schematic{}, nil
	}

	s := schematic{}
	for _, line := range ir.Lines() {
		s = append(s, []rune(line))
	}

	return s, nil
}

func (s schematic) findAdj() []int {
	var results []int
	var numStr string

	for rowIndex, chars := range s {
		for charIndex, char := range chars {
			if unicode.IsDigit(char) {
				numStr = fmt.Sprintf("%s%s", numStr, string(char))
			} else if len(numStr) > 0 {
				before := int(math.Max(float64(charIndex-len(numStr)-1), float64(0)))
				rowBefore := int(math.Max(float64(0), float64(rowIndex-1)))
				rowAfter := int(math.Min(float64(len(s)-1), float64(rowIndex+1)))

				isPart := false
				for i := before; i <= charIndex; i++ {
					if isPartSymbol(chars[i]) || isPartSymbol(s[rowBefore][i]) || isPartSymbol(s[rowAfter][i]) {
						isPart = true
						break
					}
				}

				if isPart {
					num, err := strconv.Atoi(numStr)
					if err == nil {
						results = append(results, num)
					}
				}

				numStr = ""
			}

			if len(numStr) > 0 && charIndex == len(chars)-1 {
				isPart := false
				for i := charIndex - len(numStr) - 1; i <= charIndex; i++ {
					if isPartSymbol(chars[i]) ||
						isPartSymbol(s[int(math.Max(float64(0), float64(rowIndex-1)))][i]) ||
						isPartSymbol(s[int(math.Min(float64(len(s)-1), float64(rowIndex+1)))][i]) {
						isPart = true
						break
					}
				}

				if isPart {
					num, err := strconv.Atoi(numStr)
					if err == nil {
						results = append(results, num)
					}
				}

				numStr = ""
			}
		}
	}

	return results
}

func findNumbersInLine(l string) ([]int, error) {
	re := regexp.MustCompile(`(\d+)`)
	nums := re.FindAllStringSubmatch(l, -1)

	if nums == nil {
		return []int{}, nil
	}

	var results []int
	for _, row := range nums {
		num, err := strconv.Atoi(row[0])
		if err != nil {
			return []int{}, errors.New("error converting number to int")
		}

		results = append(results, num)
	}

	return results, nil
}

func isSymbol(s rune) bool {
	re := regexp.MustCompile(`(\W)`)
	match := re.FindString(string(s))

	return match != ""
}

func isPartSymbol(r rune) bool {
	re := regexp.MustCompile(`(\W)`)
	match := re.FindString(string(r))

	return match != "" && match != "."
}

type symbolLoc struct {
	row, col int
}

func (s schematic) findSymbolLocations() []symbolLoc {
	var symbols []symbolLoc
	for rowIdx, row := range s {
		for colIdx, char := range row {
			if isPartSymbol(char) {
				symbols = append(symbols, symbolLoc{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}
	return symbols
}

type numberLoc struct {
	row, start, end, value int
}

func (s schematic) findNumberLocations() []numberLoc {
	var numbers []numberLoc
	var numStr string

	for rowIdx, row := range s {
		for colIdx, col := range row {
			if unicode.IsDigit(col) {
				numStr = fmt.Sprintf("%s%s", numStr, string(col))
			} else if len(numStr) > 0 {
				num, err := strconv.Atoi(numStr)
				if err == nil {
					numbers = append(numbers, numberLoc{
						row:   rowIdx,
						start: colIdx - len(numStr),
						end:   colIdx - 1,
						value: num,
					})
				}
				numStr = ""
			}

			if len(numStr) > 0 && colIdx == len(row)-1 {
				num, err := strconv.Atoi(numStr)
				if err == nil {
					numbers = append(numbers, numberLoc{
						row:   rowIdx,
						start: colIdx - len(numStr) + 1,
						end:   colIdx,
						value: num,
					})
				}
				numStr = ""
			}
		}
	}

	return numbers
}

func isAdjacent(n numberLoc, s symbolLoc) bool {
	if n.row == s.row {
		if s.col == n.start-1 || s.col == n.end+1 {
			return true
		}
	}

	if s.row == n.row-1 || s.row == n.row+1 {
		if s.col >= n.start-1 && s.col <= n.end+1 {
			return true
		}
	}

	return false
}

func (s schematic) findGearRatios() []int {
	var data []int
	var gears []int

	for _, symbol := range s.findSymbolLocations() {
		gears = []int{}
		for _, number := range s.findNumberLocations() {
			if isAdjacent(number, symbol) {
				gears = append(gears, number.value)
			}
		}

		if len(gears) == 2 {
			data = append(data, gears[0]*gears[1])
		}
	}

	return data
}
