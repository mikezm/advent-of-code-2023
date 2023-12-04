package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mikezm/advent-of-code-2023/day2"
	"github.com/mikezm/advent-of-code-2023/day3"
)

type challenge interface {
	A()
	B()
}

type challenges map[int]challenge

var challengeMap = challenges{
	2: day2.Challenge{},
	3: day3.Challenge{},
}

func main() {
	usage := "Usage: go run main.go <day> <part>"
	if len(os.Args) != 3 {
		fmt.Println("args provided: ", len(os.Args))
		fmt.Println(usage)
		return
	}

	// Parse the first command line argument as an integer
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing integer input:", err)
		return
	}

	var part string
	if len(os.Args) >= 2 {
		if os.Args[2] == "A" || os.Args[2] == "a" {
			part = "A"
		}

		if os.Args[2] == "B" || os.Args[2] == "b" {
			part = "B"
		}
	}

	if part == "" {
		fmt.Printf(usage)
		return
	}

	display := fmt.Sprintf("Running Day: %d part: %s", day, part)
	fmt.Println(display)
	if part == "A" {
		challengeMap[day].A()
	} else {
		challengeMap[day].B()
	}
}
