package main

import (
	"math"
	"strconv"

	"github.com/bjatkin/golf-engine/golf"
)

var nearGopher bool

func initHUD() {
	shinyFether = &ani{
		frames: []int{32, 32, 32, 32, 32, 32, 32, 32, 192, 224, 256, 288},
		speed:  0.1,
		o:      golf.SOp{TCol: golf.Col5, Fixed: true},
	}

	shinyEgg = &ani{
		frames: []int{193, 193, 193, 193, 193, 193, 193, 193, 257, 259, 261, 263},
		speed:  0.1,
		o:      golf.SOp{W: 2, H: 2, TCol: golf.Col5, Fixed: true},
	}
}

func drawHUD() {
	fether := 95.0
	eggs := 140.0

	// Outline for the HUD
	g.RectFill(0, 172, 192, 20, golf.Col7, true)
	g.Rect(0, 172, 192, 20, golf.Col0, true)

	// HP Section
	hearts := "HP "
	for i := 0; i < 3; i++ {
		if i < duck.hp {
			hearts += "<3"
		} else {
			hearts += "<4"
		}
	}
	g.Text(8, 177, hearts, golf.TOp{SW: 2, SH: 2, Fixed: true})

	// Fether count
	g.RectFill(fether-5, 174, 38, 16, golf.Col6, true)
	shinyFether.draw(fether, 179)
	g.Spr(16, fether+8, 179, golf.SOp{TCol: golf.Col5, Fixed: true})
	g.Text(fether+16, 180, strconv.Itoa(collectedFethers), golf.TOp{Fixed: true})

	// Egg count
	g.RectFill(eggs-5, 174, 44, 16, golf.Col6, true)
	shinyEgg.draw(eggs, 174)
	g.Spr(16, eggs+14, 179, golf.SOp{TCol: golf.Col5, Fixed: true})
	g.Text(eggs+22, 180, strconv.Itoa(collectedEggs), golf.TOp{Fixed: true})
}

// Draw an indicator that joe gopher is close enough to talk to
func drawSpeachBubble() {
	// the location of joe gopher
	x, y := 387.0, 125.0

	dx := math.Abs(duck.x - x)
	dy := math.Abs(duck.y - y)
	if dx > 30 || dy > 30 {
		nearGopher = false
		return
	}
	nearGopher = true

	// Draw the interaction bubble
	g.Spr(73, x, y, golf.SOp{W: 2, H: 2, TCol: golf.Col5})

	// Draw a bouncing button
	yy := y + 3
	if (g.Frames()/30)%2 == 0 {
		yy = y + 4
	}
	g.Text(x+4, y+4, "(x)", golf.TOp{Col: golf.Col4})
	g.Text(x+6, y+4, "(x)", golf.TOp{Col: golf.Col4})
	g.Text(x+5, y+4, "(x)", golf.TOp{Col: golf.Col4})
	g.Text(x+4, yy, "(x)")
	g.Text(x+6, yy, "(x)")
	g.Text(x+5, yy, "(x)")
}
