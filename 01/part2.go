package main

import (
	"regexp"
	"strconv"
)

func calculateCalibrationValuePartTwo(line string) int {
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
	lastMatch := numberFromName(r.FindAllString(line, -1)[len(r.FindAllString(line, -1))-1])
	// convert string to int
	lastMatchInt, err := strconv.Atoi(lastMatch)
	if err != nil {
		panic(err)
	}

	return firstMatchInt*10 + lastMatchInt
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
