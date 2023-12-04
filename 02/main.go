package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read input.txt line by line
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	games := make([]game, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, parseGame(line))
	}

	filteredGames := make([]game, 0)
	for _, game := range games {
		if game.fitsRestriction(12, 13, 14) {
			filteredGames = append(filteredGames, game)
		}
	}

	sumIds := 0
	for _, game := range filteredGames {
		sumIds += game.id
	}
	fmt.Println("sum of ids:")
	fmt.Println(sumIds)

	powerSum := 0
	for _, game := range games {
		minNumRed, minNumGreen, minNumBlue := game.getMinNumCubes()
		powerSum += minNumRed * minNumGreen * minNumBlue
	}
	fmt.Println("power sum:")
	fmt.Println(powerSum)
}
