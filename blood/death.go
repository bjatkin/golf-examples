package main

import (
	"fantasyConsole/golf"
)

var deathFadeCount = 0

func deathScreenUpdate() {
	g.PalA(0)
	g.PalB(1)
	deathFadeCount++
}

var deathFadeArray = [][]golf.Col{
	[]golf.Col{golf.Col0},
	[]golf.Col{golf.Col4},
	[]golf.Col{golf.Col1},
	[]golf.Col{golf.Col5},
	[]golf.Col{golf.Col2},
	[]golf.Col{golf.Col6},
	[]golf.Col{golf.Col3},
	[]golf.Col{golf.Col7},
}

func deathScreenDraw() {
	g.Cls(golf.Col0)
	g.Camera(0, 0)
	fIndex := deathFadeCount / 30
	if fIndex > 7 {
		fIndex = 7
	}
	g.Spr(141, 68, 60, golf.SOp{
		W: 7, H: 5,
		PFrom: []golf.Col{golf.Col3},
		PTo:   deathFadeArray[fIndex]})
}
