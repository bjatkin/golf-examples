package main

import "math"

var shinyFether *ani
var shinyEgg *ani
var collectedFethers int
var collectedEggs int

// Swap all the egg and fether tiles for pickups
func initPickups() {
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			tile := g.Mget(x, y)
			if tile == 32 { // Fether Tile
				g.Mset(x, y, 0)
				newFether(float64(x*8), float64(y*8))
			}
			if tile == 193 { // Egg Tile
				g.Mset(x, y, 0)
				g.Mset(x+1, y, 0)
				g.Mset(x, y+1, 0)
				g.Mset(x+1, y+1, 0)
				newEgg(float64(x*8), float64(y*8))
			}
		}
	}
}

// Check to see if any pickups were collected by the player
func checkPickups() {
	remove := -1
	for i, fether := range allFethers {
		dx := math.Abs(fether.x - duck.x)
		dy := math.Abs(fether.y - duck.y)
		if dx < 16 && dy < 16 {
			collectedFethers++
			fether.delete()

			// Add partical confetti
			for p := 0; p < 20; p++ {
				addParticle(fether.x, fether.y+10)
			}
			remove = i
			break
		}
	}
	if remove != -1 {
		allFethers[remove] = allFethers[len(allFethers)-1]
		allFethers = allFethers[:len(allFethers)-1]
	}

	remove = -1
	for i, egg := range allEggs {
		dx := math.Abs(egg.x - duck.x)
		dy := math.Abs(egg.y - duck.y)
		if dx < 16 && dy < 16 {
			collectedEggs++

			// Collected egg animation
			egg.a.(*stateAni).state++
			remove = i
			break
		}
	}
	if remove != -1 {
		allEggs[remove] = allEggs[len(allEggs)-1]
		allEggs = allEggs[:len(allEggs)-1]
	}
}
