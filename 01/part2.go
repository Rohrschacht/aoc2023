package main

import (
	"log"
	"regexp"
	"strconv"
	"unicode"
)

func calculateCalibrationValuePartTwo(line string) int {
	log.Printf("line: %s", line)

	// find first digit in string
	r := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	// find first match
	firstMatch := numberFromName(r.FindString(line))
	// convert string to int
	firstMatchInt, err := strconv.Atoi(firstMatch)
	if err != nil {
		panic(err)
	}

	// find last match
	substrings := findAllSubstrings(line)
	lastMatch := numberFromName(substrings[len(substrings)-1])
	// convert string to int
	lastMatchInt, err := strconv.Atoi(lastMatch)
	if err != nil {
		panic(err)
	}

	log.Printf("line: %s, firstMatch: %s, lastMatch: %s", line, firstMatch, lastMatch)

	return firstMatchInt*10 + lastMatchInt
}

func findAllSubstrings(line string) []string {
	var substrings []string

	numberStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := 0; i < len(line); i++ {
		if unicode.IsNumber(rune(line[i])) {
			substrings = append(substrings, line[i:i+1])
		}

		for _, numberString := range numberStrings {
			if i <= len(line)-len(numberString) && line[i:i+len(numberString)] == numberString {
				substrings = append(substrings, numberString)
			}
		}
	}

	log.Printf("substrings: %v", substrings)

	return substrings
}

func numberFromName(digit string) string {
	if len(digit) == 1 {
		return digit
	}

	lut := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	return lut[digit]
}
