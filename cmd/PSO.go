package main

import (
	"math"
	"math/rand/v2"
)

type Particle struct {
	Position float64
	Velocity float64
	PBest    float64
}

func ParticleSwarmOptimization() *float64 {
	// config
	inertia := 0.9
	cognitivePull := 2.0
	socialPull := 2.0
	iterations := 1000

	var Gbest float64

	particles := GenerateParticles(&Gbest)

	// run through iterations
	for range iterations {
		for _, particle := range particles {
			UpdateVandX(particle, inertia, cognitivePull, socialPull, iterations, &Gbest)
		}
	}

	return &Gbest
}

func UpdateVandX(particle *Particle, inertia, cognitivePull, socialPull float64, iterations int, Gbest *float64) {
	// Update the velocity
	particle.Velocity = (inertia * particle.Velocity) +
		(cognitivePull*rand.Float64())*(particle.PBest-particle.Position) +
		(socialPull*rand.Float64())*(*Gbest-particle.Position)

	particle.Position = particle.Position + particle.Velocity

	// update personal best
	if PCOFitness(particle.PBest) > PCOFitness(particle.Position) {
		particle.PBest = particle.Position
	}

	// Update gbest accordingly
	if PCOFitness(particle.PBest) < PCOFitness(*Gbest) {
		*Gbest = particle.PBest
	}
}

func GenerateParticles(Gbest *float64) []*Particle {
	particles := make([]*Particle, 0)

	// Generate 100 random particles
	for i := 0; i < 100; i++ {
		// random num between x and y = rand.IntN(x - y) - x
		randPos := float64(rand.IntN(10+10) - 10)
		randVel := float64(rand.IntN(1+1) - 1)

		// Create new random particle
		particle := &Particle{
			Position: randPos,
			Velocity: randVel,
			PBest:    randPos,
		}

		particles = append(particles, particle)

		// Update gbest accordingly
		if i == 0 {
			*Gbest = particle.PBest
			continue
		}

		fitness := PCOFitness(particle.PBest)
		if PCOFitness(*Gbest) > fitness {
			*Gbest = particle.Position
		}
	}

	return particles
}

func PCOFitness(solution float64) float64 {
	return math.Pow(solution, 2)
}
