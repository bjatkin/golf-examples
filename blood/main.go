package main

import (
	"fantasyConsole/golf"
	"fmt"
	"math/rand"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)

	g.BG(golf.Col0)
	g.LoadSprs(spriteSheet)
	g.LoadMap(mapData)
	initGame()
	g.Run()
}

var player int
var debug string

func initGame() {
	player = addEntity(entity{
		flags:  playerControlled,
		sig:    pos | spr | hp | colidable,
		hp:     100,
		pos:    posComp{10, 10},
		spr:    sprComp{ani: [10]int{2, 3}, aniLen: 2, aniSpeed: 60, opt: golf.SprOpts{Height: 2, Transparent: golf.Col7}},
		colide: colidableComp{width: 8, height: 16, oldX: 10, oldY: 10},
	})

	//Wall
	addEntity(entity{
		sig:    pos | spr | colidable,
		pos:    posComp{50, 150},
		spr:    sprComp{ani: [10]int{68}, aniLen: 1, opt: golf.SprOpts{Height: 2, Width: 4}},
		colide: colidableComp{width: 32, height: 16, oldX: 50, oldY: 150},
	})

	// for i := 0; i < 15; i++ {
	addEnemy(rand.Intn(192), rand.Intn(192))
	addEnemy(rand.Intn(192), rand.Intn(192))
	// }

	allUpdateSystems[movePlayer] = toSystem(playerControlled, pos, func(e *entity) {
		e.spr.ani = [10]int{2, 3}
		e.spr.aniLen = 2
		e.spr.aniSpeed = 60
		e.spr.opt.Width = 1
		e.colide.width = 8
		if g.Btn(golf.WKey) {
			e.pos.y--
			e.spr.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
			e.spr.aniLen = 8
			e.spr.aniSpeed = 10
			e.spr.opt.Width = 2
			e.colide.width = 16
		}
		if g.Btn(golf.SKey) {
			e.pos.y++
			e.spr.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
			e.spr.aniLen = 8
			e.spr.aniSpeed = 10
			e.spr.opt.Width = 2
			e.colide.width = 16
		}
		if g.Btn(golf.AKey) {
			e.pos.x--
			e.spr.ani = [10]int{6, 8, 10, 12}
			e.spr.aniLen = 4
			e.spr.aniSpeed = 10
			e.spr.opt.FlipH = false
			e.spr.opt.Width = 2
			e.colide.width = 16
		}
		if g.Btn(golf.DKey) {
			e.pos.x++
			e.spr.ani = [10]int{6, 8, 10, 12}
			e.spr.aniLen = 4
			e.spr.aniSpeed = 10
			e.spr.opt.FlipH = true
			e.spr.opt.Width = 2
			e.colide.width = 16
		}
	})

	allUpdateSystems[doAI] = toSystem(0, ai|pos, func(e *entity) {
		t := allEntities[e.ai.target]
		dist := (((t.pos.x - e.pos.x) * (t.pos.x - e.pos.x)) + ((t.pos.y - e.pos.y) * (t.pos.y - e.pos.y)))
		if dist > e.ai.atkRange*e.ai.atkRange {
			return
		}
		e.spr.ani = [10]int{24, 25}
		e.spr.aniLen = 2
		if g.Frames()%5 != 0 { //Dont want to move too fast!
			return
		}

		e.spr.ani = [10]int{24, 25, 26}
		e.spr.aniLen = 3
		if e.pos.x < t.pos.x {
			e.pos.x++
			e.spr.opt.FlipH = true
			debug += "ai moving right\n"
		}
		if e.pos.x > t.pos.x {
			e.pos.x--
			e.spr.opt.FlipH = false
			debug += "ai moving left\n"
		}
		if e.pos.y < t.pos.y {
			e.pos.y++
			debug += "ai moving up\n"
		}
		if e.pos.y > t.pos.y {
			e.pos.y--
			debug += "ai moving down\n"
		}
	})

	allUpdateSystems[doAttack] = toSystem(enemy, ai|pos, func(e *entity) {
		t := &allEntities[e.ai.target]
		hDist := t.pos.x - e.pos.x
		vDist := t.pos.y - e.pos.y
		if hDist < 0 {
			hDist *= -1
		}
		if vDist < 0 {
			vDist *= -1
		}
		if hDist < 3 && vDist < 3 {
			t.hp--
		}
	})

	allUpdateSystems[physicsRound0] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound1] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound2] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound3] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound4] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound5] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound6] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound7] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound8] = toSystem(0, pos|colidable, runPhysicsStep)
	allUpdateSystems[physicsRound9] = toSystem(0, pos|colidable, runPhysicsStep)

	allUpdateSystems[resolvePhysics] = toSystem(0, colidable|pos, func(e *entity) {
		e.pos.x = e.colide.oldX + int(e.colide.deltaX)
		e.pos.y = e.colide.oldY + int(e.colide.deltaY)
		e.colide.deltaX = 0
		e.colide.deltaY = 0
		e.colide.oldX = e.pos.x
		e.colide.oldY = e.pos.y
	})

	allDrawSystems[drawSprite] = toSystem(0, pos|spr, func(e *entity) {
		e.spr.frame++
		if e.spr.aniSpeed == 0 {
			e.spr.aniSpeed = 1
		}
		if e.spr.frame%e.spr.aniSpeed == 0 {
			e.spr.aniFrame++
		}
		if e.spr.aniFrame >= e.spr.aniLen {
			e.spr.aniFrame = 0
		}
		frame := e.spr.ani[e.spr.aniFrame]
		opt := e.spr.opt
		if frame < 0 {
			opt.FlipH = true
			frame *= -1
		}
		g.Spr(frame, e.pos.x, e.pos.y, opt)
		if e.hasComponent(colidable) {
			g.Rect(e.pos.x, e.pos.y, e.colide.width, e.colide.height, golf.Col1)
		}
	})

	allDrawSystems[drawHP] = toSystem(playerControlled, hp, func(e *entity) {
		if e.hp < 0 {
			g.TextL("You DED!", golf.TextOpts{Col: golf.Col3})
			return
		}
		g.TextL(fmt.Sprintf("HP: %d", e.hp), golf.TextOpts{Col: golf.Col3})
	})
}

func update() {
	runUpdateSystems()
}

func draw() {
	g.Cls()
	runDrawSystems()
	g.TextR(debug, golf.TextOpts{Col: golf.Col3})
	debug = ""
}

func addEnemy(x, y int) int {
	// zombie
	id := addEntity(entity{
		flags:  enemy,
		sig:    pos | spr | ai | colidable,
		pos:    posComp{x, y},
		spr:    sprComp{ani: [10]int{24, 25, 26}, aniLen: 3, aniSpeed: 30, opt: golf.SprOpts{Height: 2, Transparent: golf.Col7}},
		ai:     aiComp{atkRange: 150, target: player},
		colide: colidableComp{width: 8, height: 16, oldX: x, oldY: y},
	})
	return id
}
