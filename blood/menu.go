package main

import "fantasyConsole/golf"

var menuY, yTarget = 177.0, 377.0
var beta, tSpeed = 0.0, 0.01
var move = false
var menuFadeArray = [][]golf.Col{
	[]golf.Col{golf.Col3, golf.Col0},
	[]golf.Col{golf.Col7, golf.Col0},
	[]golf.Col{golf.Col6, golf.Col0},
	[]golf.Col{golf.Col5, golf.Col0},
	[]golf.Col{golf.Col4, golf.Col0},
	[]golf.Col{golf.Col0, golf.Col4},
	[]golf.Col{golf.Col0, golf.Col5},
	[]golf.Col{golf.Col0, golf.Col6},
	[]golf.Col{golf.Col0, golf.Col7},
	[]golf.Col{golf.Col0, golf.Col3},
}

func menuUpdate() {
	g.PalA(15)
	g.PalB(0)
	if g.Btnp(golf.ZKey) {
		move = true
	}
	if move {
		beta += tSpeed
	}
	if beta > 1.1 {
		g.Update = gameUpdate
		g.Draw = gameDraw
	}
}

func menuDraw() {
	g.Cls(golf.Col0)
	fIndex := int(lerp(0, 9, beta))
	g.Map(40, 20, 40, 70, 0, 0,
		golf.SOp{
			PFrom: []golf.Col{golf.Col0, golf.Col3},
			PTo:   menuFadeArray[fIndex]})

	cy := int(lerp(menuY, yTarget, beta))
	g.Camera(100, cy)

	if !move {
		g.RectFill(120, 290, 160, 14, golf.Col0)
		g.Text(130, 295, "press the z key to start", whiteTxt)
	}
}
