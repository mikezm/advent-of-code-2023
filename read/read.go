package read

import (
	"bufio"
	"log"
	"os"
)

type InputReader struct {
	f       os.File
	scanner bufio.Scanner
}

func NewReader(f string) (*InputReader, error) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	return &InputReader{
		f:       *file,
		scanner: *scanner,
	}, nil
}

func (ir *InputReader) Lines() []string {
	defer ir.close()
	var lines []string
	for ir.scanner.Scan() {
		lines = append(lines, ir.scanner.Text())

	}
	return lines
}

func (ir *InputReader) close() {
	err := ir.f.Close()
	if err != nil {
		return
	}
}
