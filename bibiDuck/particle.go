package main

import (
	"math/rand"

	"github.com/bjatkin/golf-engine/golf"
)

type particle struct {
	x, y float64
	life int
}

var allParticles = [100]*particle{}
var pPointer = 0

// fill the particle system so we don't have null pointer errors
func initParticles() {
	for i := 0; i < len(allParticles); i++ {
		allParticles[i] = &particle{0, 0, 0}
	}
}

func addParticle(x, y float64) {
	// Just reuse old particles, don't create new ones
	// should make things run faster by preventing heap
	// allocations
	allParticles[pPointer].x = x
	allParticles[pPointer].y = y
	allParticles[pPointer].life = 20
	pPointer++
	if pPointer >= len(allParticles) {
		pPointer = 0
	}
}

// Draws all the living particles
func drawParticles() {
	for _, p := range allParticles {
		if p.life <= 0 {
			continue
		}

		p.x += float64(rand.Intn(3) - 1)
		p.y += float64(rand.Intn(5) - 3)
		p.life--

		g.Pset(p.x-float64(cameraX), p.y, golf.Col7)
	}
}
