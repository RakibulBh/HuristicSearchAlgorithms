package main

import "fmt"

func main() {
	// Particle optimisation algorithm
	// solution := ParticleSwarmOptimization()
	// fmt.Printf("The answer to the solution is %f \n", *solution)

	//RHMC algorithm
	// distances_easy := [][]int{
	// 	{0, 10, 15, 20},
	// 	{10, 0, 35, 25},
	// 	{15, 35, 0, 30},
	// 	{20, 25, 30, 0},
	// }
	distancesHard := [][]int{
		{0, 12, 18, 25, 30, 22, 15, 28, 20}, // A to A, B, C, D, E, F, G, H, I
		{12, 0, 27, 19, 16, 24, 32, 21, 26}, // B to A, B, C, D, E, F, G, H, I
		{18, 27, 0, 23, 29, 17, 20, 33, 14}, // C to A, B, C, D, E, F, G, H, I
		{25, 19, 23, 0, 15, 28, 31, 16, 22}, // D to A, B, C, D, E, F, G, H, I
		{30, 16, 29, 15, 0, 20, 24, 18, 27}, // E to A, B, C, D, E, F, G, H, I
		{22, 24, 17, 28, 20, 0, 19, 25, 30}, // F to A, B, C, D, E, F, G, H, I
		{15, 32, 20, 31, 24, 19, 0, 26, 21}, // G to A, B, C, D, E, F, G, H, I
		{28, 21, 33, 16, 18, 25, 26, 0, 19}, // H to A, B, C, D, E, F, G, H, I
		{20, 26, 14, 22, 27, 30, 21, 19, 0}, // I to A, B, C, D, E, F, G, H, I
	}
	distance, solution := RHMC(distancesHard, 100000)
	fmt.Printf("The solution found was %d with a total diatnce of %s \n", distance, solution)
}
