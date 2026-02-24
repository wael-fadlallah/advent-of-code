package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
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

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func calcSimilarity(listA []int, listB []int) int {
	var li, ri int
	similarityScore := 0

	for i := range listA {
		li, ri = i, i
		similarityCount := 0

		// check the corresponding index
		if listB[i] == listA[i] {
			similarityCount++
		}

		// left index is at the start of the slice
		for li > 0 {
			li--
			if listB[li] == listA[i] {
				similarityCount++
			}
		}

		// right index is at the end of the slice
		for ri < len(listB)-1 {
			ri++
			if listB[ri] == listA[i] {
				similarityCount++
			}
		}

		similarityScore += listA[i] * similarityCount
	}

	return similarityScore
}

func main() {
	var result int
	data, err := readInputFile("input.txt")
	if err != nil {
		fmt.Printf("error reading the file: %s\n", err)
	}
	var listA, listB []int = []int{}, []int{}

	lines := strings.Split(data, "\n")

	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], "|")

		if len(data) < 2 {
			continue
		}
		numA, err := strconv.Atoi(strings.Trim(data[0], " "))
		if err != nil {
			fmt.Println("Error parsing numA: ", err)
			continue
		}
		numB, err := strconv.Atoi(strings.Trim(data[1], " "))
		if err != nil {
			fmt.Println("Error parsing numB: ", err)
			continue
		}

		listA = append(listA, int(numA))
		listB = append(listB, int(numB))
	}

	sort.Ints(listA)
	sort.Ints(listB)

	for index := range listA {
		result += absInt(listA[index] - listB[index])
	}

	fmt.Printf("result is %d", result)
	fmt.Printf("Similarity score: %d", calcSimilarity(listA, listB))

}
