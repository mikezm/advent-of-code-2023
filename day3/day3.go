package day3

import (
	"errors"
	"fmt"
	"github.com/mikezm/advent-of-code-2023/read"
	"math"
	"regexp"
	"strconv"
)

const INPUTS = "./day3/input.txt"

type Challenge struct{}

func (d Challenge) A() {
	s, err := readSchematic()
	if err != nil {
		fmt.Println("error reading input file")
	}

	for _, l := range s {
		fmt.Println(l)
	}

}

func (d Challenge) B() {
	fmt.Println("function B() not yet implemented")
}

type schematic []string

func readSchematic() (schematic, error) {
	ir, err := read.NewReader(INPUTS)

	if err != nil {
		return schematic{}, nil
	}

	s := schematic{}
	for _, line := range ir.Lines() {
		s = append(s, line)
	}

	return s, nil
}

func (s schematic) findNumsAdjToSymbol() []int {
	var results []int
	for lineIndex, line := range s {
		nums, err := findNumbersInLine(line)
		if err != nil {
			return results
		}

		for _, n := range nums {
			numStr := fmt.Sprintf("%d", n)
			re := regexp.MustCompile(numStr)
			loc := re.FindStringIndex(line)

			lnBefore := int(math.Max(float64(0), float64(lineIndex-1)))
			lnAfter := int(math.Min(float64(len(s)), float64(lineIndex+1))) + 1
			before := int(math.Max(float64(loc[0]-1), float64(0)))
			after := int(math.Min(float64(loc[1]+1), float64(len(line)))) + 1

			part := s[lnBefore:lnAfter]
			for _, ll := range part {
				for _, char := range ll[before:after] {
					c := string(char)
					if isSymbol(c) && !inArray(results, n) {
						results = append(results, n)
					}
				}
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

func findNumbersInLine(l string) ([]int, error) {
	re := regexp.MustCompile(`(\d+)`)
	nums := re.FindAllStringSubmatch(l, -1)

	if nums == nil {
		return []int{}, errors.New("could not find any numbers")
	}

	var results []int
	for _, row := range nums {
		for _, n := range row {
			num, err := strconv.Atoi(n)
			if err != nil {
				return []int{}, errors.New("error converting number to int")
			}

			if !inArray(results, num) {
				results = append(results, num)
			}

		}
	}

	//re := regexp.MustCompile(reg)
	//match := re.FindStringSubmatch(s)

	return results, nil
}

func isSymbol(s string) bool {
	re := regexp.MustCompile(`(\W)`)
	match := re.FindString(s)

	return match != "" && match != "."
}
