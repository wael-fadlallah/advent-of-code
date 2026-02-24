package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func readInputFile(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return "", errors.New("error opening file")
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return "", errors.New("error reading the file")
	}
	return string(content), nil
}

func main() {
	content, err := readInputFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// Regular expressions to match instructions
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	// Find all matches in the content
	mulMatches := reMul.FindAllStringSubmatchIndex(content, -1)
	doMatches := reDo.FindAllStringIndex(content, -1)
	dontMatches := reDont.FindAllStringIndex(content, -1)

	// Combine all matches into a single slice and sort by position
	allMatches := append(mulMatches, doMatches...)
	allMatches = append(allMatches, dontMatches...)
	sort.Slice(allMatches, func(i, j int) bool {
		return allMatches[i][0] < allMatches[j][0]
	})

	var result int
	enabled := true

	for _, match := range allMatches {
		// Check if it's a do() instruction
		if len(match) == 2 && string(content[match[0]:match[1]]) == "do()" {
			enabled = true
			continue
		}

		// Check if it's a don't() instruction
		if len(match) == 2 && string(content[match[0]:match[1]]) == "don't()" {
			enabled = false
			continue
		}

		// Check if it's a mul(x,y) instruction
		if len(match) == 6 {
			if enabled {
				x, err1 := strconv.Atoi(content[match[2]:match[3]])
				y, err2 := strconv.Atoi(content[match[4]:match[5]])
				if err1 != nil || err2 != nil {
					continue
				}
				result += x * y
				fmt.Printf("mul(%d,%d) = %d\n", x, y, result)
			}
		}
	}

	fmt.Printf("Total result: %d\n", result)
}
