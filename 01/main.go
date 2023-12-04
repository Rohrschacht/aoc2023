package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func calculateCalibrationValue(line string) int {
	// find first digit in string
	r := regexp.MustCompile(`\d`)
	// find first match
	firstMatch := r.FindString(line)
	// convert string to int
	firstMatchInt, err := strconv.Atoi(firstMatch)
	if err != nil {
		panic(err)
	}

	// find last match
	lastMatch := r.FindAllString(line, -1)[len(r.FindAllString(line, -1))-1]
	// convert string to int
	lastMatchInt, err := strconv.Atoi(lastMatch)
	if err != nil {
		panic(err)
	}

	return firstMatchInt*10 + lastMatchInt
}

func main() {
	// read input.txt line by line
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//sum += calculateCalibrationValue(line)
		sum += calculateCalibrationValuePartTwo(line)
	}

	fmt.Println(sum)
}
