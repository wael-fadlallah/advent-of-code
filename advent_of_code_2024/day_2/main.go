package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func convertStringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints[i] = num
	}
	return ints, nil
}

func isReportSafe(levels []int) bool {
	var increasing, decreasing = true, true
	var isSafe = true

	if levels[0] < levels[1] {
		decreasing = false
	} else {
		increasing = false
	}

	for j := 0; j < len(levels)-1; j++ {
		num := levels[j]
		num1 := levels[j+1]

		// increasing
		if increasing && !(num < num1 && (num1-num) > 0 && (num1-num) <= 3) {
			isSafe = false
		} else if decreasing && !(num > num1 && (num-num1) > 0 && (num-num1) <= 3) {
			isSafe = false
		}

	}
	return isSafe
}

func main() {
	var safeReports int

	data, _ := readInputFile("input.txt")
	reports := strings.Split(data, "\n")

	for i := range reports {
		levelsStr := strings.Split(reports[i], " ")
		levels, _ := convertStringsToInts(levelsStr)

		if isReportSafe(levels) {
			safeReports++
		} else {
			for j := 0; j < len(levels); j++ {
				newReport := append([]int(nil), levels[:j]...)
				newReport = append(newReport, levels[j+1:]...)
				if isReportSafe(newReport) {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Println(safeReports)
}
