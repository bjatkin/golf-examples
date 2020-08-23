package main

import "fantasyConsole/golf"

var winFadeCount = 0

func winScreenUpdate() {
	g.PalA(0)
	g.PalB(1)
	winFadeCount++
}

var winFadeArray = [][]golf.Col{
	[]golf.Col{golf.Col0},
	[]golf.Col{golf.Col4},
	[]golf.Col{golf.Col1},
	[]golf.Col{golf.Col5},
	[]golf.Col{golf.Col2},
	[]golf.Col{golf.Col6},
	[]golf.Col{golf.Col3},
	[]golf.Col{golf.Col7},
}

func winScreenDraw() {
	g.Cls(golf.Col0)
	g.Camera(0, 0)
	fIndex := winFadeCount / 30
	if fIndex > 7 {
		fIndex = 7
	}
	g.Spr(148, 48, 60, golf.SOp{
		W: 12, H: 5,
		PFrom: []golf.Col{golf.Col3},
		PTo:   winFadeArray[fIndex]})
}
