package main

import (
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id       int
	drawings []drawing
}

type drawing struct {
	numRed   int
	numGreen int
	numBlue  int
}

func parseGame(line string) game {
	gameRe := regexp.MustCompile(`Game (\d+):(.*)`)
	id := gameRe.FindStringSubmatch(line)[1]
	intId, _ := strconv.Atoi(id)

	rest := gameRe.FindStringSubmatch(line)[2]

	drawings := make([]drawing, 0)
	for _, drawingSubString := range strings.Split(rest, ";") {
		d := drawing{}

		redRe := regexp.MustCompile(`(\d+) red`)
		greenRe := regexp.MustCompile(`(\d+) green`)
		blueRe := regexp.MustCompile(`(\d+) blue`)

		if redRe.MatchString(drawingSubString) {
			numRed := redRe.FindStringSubmatch(drawingSubString)[1]
			d.numRed, _ = strconv.Atoi(numRed)
		}

		if greenRe.MatchString(drawingSubString) {
			numGreen := greenRe.FindStringSubmatch(drawingSubString)[1]
			d.numGreen, _ = strconv.Atoi(numGreen)
		}

		if blueRe.MatchString(drawingSubString) {
			numBlue := blueRe.FindStringSubmatch(drawingSubString)[1]
			d.numBlue, _ = strconv.Atoi(numBlue)
		}

		drawings = append(drawings, d)
	}

	return game{
		id:       intId,
		drawings: drawings,
	}
}

func (g game) fitsRestriction(r, gr, b int) bool {
	fits := true

	for _, drawing := range g.drawings {
		if drawing.numRed > r || drawing.numGreen > gr || drawing.numBlue > b {
			fits = false
			break
		}
	}

	return fits
}

func (g game) getMinNumCubes() (int, int, int) {
	minNumRed := 0
	minNumGreen := 0
	minNumBlue := 0

	for _, drawing := range g.drawings {
		if drawing.numRed > minNumRed {
			minNumRed = drawing.numRed
		}

		if drawing.numGreen > minNumGreen {
			minNumGreen = drawing.numGreen
		}

		if drawing.numBlue > minNumBlue {
			minNumBlue = drawing.numBlue
		}
	}

	return minNumRed, minNumGreen, minNumBlue
}
