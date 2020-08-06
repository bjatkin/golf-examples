package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	// "github.com/bjatkin/golf-engine/golf"
	"fantasyConsole/golf"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)

	g.LoadSprs(spriteSheet)
	g.LoadMap(mapData)
	// g.LoadFlags(flagData)
	g.DrawMouse(1)

	// g.RAM[0x3601] = 0
	g.Run()
}

var clipped bool
var cameraX, cameraY int

func update() {
	if g.Btnp(golf.ZKey) {
		if clipped {
			g.RClip()
		} else {
			g.Clip(30, 50, 132, 71)
		}

		clipped = !clipped
	}

	if g.Btnp(golf.XKey) {
		cameraX, cameraY = 0, 0
	}
	if g.Btn(golf.UpArrow) {
		cameraY -= 5
	}
	if g.Btn(golf.DownArrow) {
		cameraY += 5
	}
	if g.Btn(golf.LeftArrow) {
		cameraX -= 5
	}
	if g.Btn(golf.RightArrow) {
		cameraX += 5
	}
	g.Camera(cameraX, cameraY)

	if g.Mbtnp(golf.LeftClick) {
		mx, my := g.Mouse()
		// left up arrow
		if mx > 8 && mx < 24 && my > 64 && my < 80 {
			palA++
			lUp = 12
			if palA > golf.Pal15 {
				palA = golf.Pal0
			}
		}
		// left down arrow
		if mx > 8 && mx < 24 && my > 88 && my < 104 {
			palA--
			lDown = 12
			if palA > golf.Pal15 {
				palA = golf.Pal15
			}
		}
		// right up arrow
		if mx > 168 && mx < 184 && my > 64 && my < 80 {
			palB++
			rUp = 12
			if palB > golf.Pal15 {
				palB = golf.Pal0
			}
		}
		// right down arrow
		if mx > 168 && mx < 184 && my > 88 && my < 104 {
			palB--
			rDown = 12
			if palB > golf.Pal15 {
				palB = golf.Pal15
			}
		}
	}
	g.PalA(palA)
	g.PalB(palB)

}

func draw() {
	g.Cls(golf.Col7)
	g.Map(0, 0, 24, 24, 0, 0)
	g.RectFill(0, 0, 192, 8, golf.Col6)
	g.RectFill(0, 184, 192, 8, golf.Col6)
	g.TextL("GoLF API Demo")
	g.TextR("FPS: " + strconv.Itoa(calcFPS()))
	g.Text(0, 185, "Press (<)(>)(^)(v), z or x")

	scale := math.Sin(float64(g.Frames()-255) / 30)
	flip := 1.0
	if scale < 0 {
		flip = -1.0
	}

	drawSlime()

	drawLogo(scale*flip, flip != -1)

	drawArrows()

	mx, my := g.Mouse()
	for i := 0; i < 3; i++ {
		addPartical(mx+rand.Intn(15)-7, my+rand.Intn(15)-7, golf.Col4)
	}
	drawParticals()
}

func drawLogo(scale float64, flip bool) {
	// Change to the internal sprite sheet
	width := 64.0 * scale
	logoBuff := 0x3647
	spriteBase := 0x3F48

	g.RAM[0x6F49] = byte(logoBuff >> 8)
	g.RAM[0x6F4A] = byte(logoBuff & 0b0000000011111111)
	fadeFrom := []golf.Col{golf.Col0, golf.Col3, golf.Col5, golf.Col7}
	fadeTo := []golf.Col{golf.Col0, golf.Col3, golf.Col5, golf.Col7}
	if flip {
		fadeTo = []golf.Col{golf.Col0, golf.Col1, golf.Col1, golf.Col1}
	}

	g.SSpr(152, 0, 64, 24, float64(96-width), 64.0, golf.SOp{TCol: golf.Col1, SW: scale * 2, SH: 2.0, FH: flip, PFrom: fadeFrom, PTo: fadeTo})

	// Change back to the main sprite sheet
	g.RAM[0x6F49] = byte(spriteBase >> 8)
	g.RAM[0x6F4A] = byte(spriteBase & 0b0000000011111111)
}

var fpsCounter int
var lastFPSCheck int64
var currentFPS int

func calcFPS() int {
	fpsCounter++
	sec := 1000000000
	now := time.Now().UnixNano()
	if lastFPSCheck+int64(sec) < now {
		currentFPS = fpsCounter
		fpsCounter = 0
		lastFPSCheck = time.Now().UnixNano()
		return currentFPS
	}
	return currentFPS
}

var palA, palB = golf.Pal0, golf.Pal1
var lUp, lDown, rUp, rDown int

func drawArrows() {
	s := 0
	if lUp > 0 {
		s = 2
	}
	g.Spr(s, 8, 64, golf.SOp{TCol: golf.Col7, W: 2, H: 2})
	s = 4
	if lDown > 0 {
		s = 6
	}
	g.Spr(s, 8, 88, golf.SOp{TCol: golf.Col7, W: 2, H: 2})
	s = 0
	if rUp > 0 {
		s = 2
	}
	g.Spr(s, 168, 64, golf.SOp{TCol: golf.Col7, W: 2, H: 2})
	s = 4
	if rDown > 0 {
		s = 6
	}
	lUp--
	lDown--
	rUp--
	rDown--
	g.Spr(s, 168, 88, golf.SOp{TCol: golf.Col7, W: 2, H: 2})
	g.Text(0, 82, fmt.Sprintf("Pal %d", (palA&0b01111111)))
	g.Text(156, 82, fmt.Sprintf("Pal %d", (palB&0b01111111)))
}

func drawSlime() {
	g.SSpr(112, 0, 136, 80, 10, 104, golf.SOp{TCol: golf.Col7})

	mx, my := g.Mouse()
	drawSlimeEye(55, 130, float64(mx+cameraX), float64(my+cameraY))
	drawSlimeEye(80, 130, float64(mx+cameraX), float64(my+cameraY))
}

func drawSlimeEye(cx, cy, lx, ly float64) {
	dist := math.Sqrt((lx-cx)*(lx-cx) + (ly-cy)*(ly-cy))
	dX, dY := lx-cx, ly-cy
	g.CircFill(cx, cy, 12.0, golf.Col1)
	delta := 1.0 / dist
	if dist < 3 {
		delta = 0
	}
	g.CircFill(cx+dX*delta, cy+dY*delta, 10.0, golf.Col7)
	delta = 8.0 / dist
	if dist < 3 {
		delta = 0
	}
	g.CircFill(cx+dX*delta, cy+dY*delta, 3.0, golf.Col0)
}

type partical struct {
	life int
	x, y int
	col  golf.Col
}

var particals = [50]partical{}
var pIndex = 0

func addPartical(x, y int, col golf.Col) {
	particals[pIndex] = partical{60, x, y, col}
	pIndex++
	if pIndex > len(particals)-1 {
		pIndex = 0
	}
}

func drawParticals() {
	for i := 0; i < len(particals); i++ {
		if particals[i].life > 0 {
			g.Pset(float64(particals[i].x), float64(particals[i].y), particals[i].col)
			particals[i].x += rand.Intn(3) - 1
			particals[i].y += rand.Intn(3) - 1
			particals[i].life--
		}
	}
}
