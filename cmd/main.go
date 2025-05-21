package main

import "fmt"

func main() {
	// Particle optimisation algorithm
	solution := ParticleSwarmOptimization()

	fmt.Printf("The answer to the solution is %f \n", *solution)
}
