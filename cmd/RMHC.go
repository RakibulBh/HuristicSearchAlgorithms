package main

import (
	"math/rand/v2"
)

func RHMC(distances [][]int, iterations int) (int, []string) {
	table := make(map[string]int)
	for i := 0; i < len(distances); i++ {
		table[string(rune('A'+i))] = i
	}

	// Generate a random solution
	currSolution := GenerateRandomSolution(distances)
	currentMin := fitness(currSolution, distances, table)
	currentSolution := make([]string, len(*currSolution))
	copy(currentSolution, *currSolution)

	for range iterations {
		fitness := fitness(currSolution, distances, table)
		if fitness < currentMin {
			currentMin = fitness
			currentSolution = make([]string, len(*currSolution))
			copy(currentSolution, *currSolution)
		}
		currSolution = SmallChangeOperator(currSolution)
	}

	return currentMin, currentSolution
}

func fitness(tour *[]string, distances [][]int, table map[string]int) int {

	totalDistance := 0

	for i := 0; i < len(*tour)-1; i++ {
		totalDistance += distances[table[(*tour)[i]]][table[(*tour)[i+1]]]
	}

	return totalDistance + distances[table[(*tour)[len(*tour)-1]]][table[(*tour)[0]]]
}

func GenerateRandomSolution(distances [][]int) *[]string {
	if len(distances) != len((distances[0])) {
		return nil
	}

	// dictionary for mapping cities with distances
	table := make(map[int]string)
	for i := 0; i < len(distances); i++ {
		table[i] = string(rune('A' + i))
	}

	tour := make([]string, len(distances))

	for i := range tour {
		tour[i] = table[i]
	}

	rand.Shuffle(len(tour), func(i, j int) {
		tour[i], tour[j] = tour[j], tour[i]
	})

	return &tour
}

func SmallChangeOperator(tour *[]string) *[]string {
	first := rand.IntN(len(*tour))
	second := rand.IntN(len(*tour))
	for first == second {
		second = rand.IntN(len(*tour))
	}

	temp := (*tour)[first]
	(*tour)[first] = (*tour)[second]
	(*tour)[second] = temp

	return tour
}
