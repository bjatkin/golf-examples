package main

import "github.com/bjatkin/golf-engine/golf"

// conversation with joe gopher
var gopherConvo *convo

func initGopher() {
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if g.Mget(x, y) == 201 { // The gopher sprite

				// erase the gopher tiles
				g.Mset(x, y, 0)
				g.Mset(x+1, y, 0)
				g.Mset(x, y+1, 0)
				g.Mset(x+1, y+1, 0)

				// replace it the gopher character
				newMob(float64(x*8), float64(y*8), []drawable{
					&ani{ //Waiting
						frames: []int{201, 203},
						speed:  0.05,
						o:      golf.SOp{W: 2, H: 2},
					},
					&ani{ //Talking
						frames: []int{137, 201},
						speed:  0.05,
						o:      golf.SOp{W: 2, H: 2},
					},
				})
				break
			}
		}
	}
}
