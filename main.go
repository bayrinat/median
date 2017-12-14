package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bayrinat/median/metric"
)

func usage() string {
	return "Usage: median /path/to/file.csv /path/to/out"
}

func inFile(path string) (*os.File, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return os.Open(path)
}

// Open out file, rewrite file if exists
func outFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0660)
}

func readCSV(file io.Reader) ([][]string, error) {
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func main() {
	if len(os.Args) != 3 {
		log.Printf("wrong args count, want 3, got %v \n", len(os.Args))
		fmt.Println(usage())
		return
	}

	inFile, err := inFile(os.Args[1])
	if err != nil {
		log.Printf("failed to open inFile: %v \n", err)
		fmt.Println(usage())
		return
	}
	defer inFile.Close()

	lines, err := readCSV(inFile)
	if err != nil {
		log.Printf("failed to read inFile as csv inFile: %v \n", err)
		fmt.Println(usage())
		return
	}

	delays, err := metric.NewDelay(5)
	if err != nil {
		log.Printf("failed to create Delay instance %v \n", err)
		fmt.Println(usage())
		return
	}

	outFile, err := outFile(os.Args[2])
	if err != nil {
		log.Printf("failed to create or open output file with: %v", err)
		return
	}
	defer outFile.Close()

	for row, line := range lines {
		for column, word := range line {
			delay, err := strconv.Atoi(strings.Trim(word, "\r "))
			if err != nil {
				log.Printf("failed to convert word to int, row: %v, column: %v, got: %v", row, column, err)
				return
			}
			delays.AddDelay(delay)
		}

		// Write median to file
		outFile.WriteString(fmt.Sprintf("%v\n", delays.GetMedian()))
	}
}
