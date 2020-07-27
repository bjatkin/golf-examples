package main

import (
	"fantasyConsole/golf"
	"fmt"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)

	g.SetBG(golf.Col0)
	g.LoadSprs(spriteSheet)
	g.LoadMap(mapData)
	initGame()
	g.Run()
}

var player *entity
var debug string

func initGame() {
	player = newEntity(playerControlled,
		&hpComponent{health: 100},
		&transformComponent{x: 10, y: 10},
		&sprComponent{ani: [10]int{2, 3}, aniLen: 2, aniSpeed: 60, opt: golf.SOp{H: 2, TCol: golf.Col7}},
		&collideComponent{width: 8, height: 16, oldX: 10, oldY: 10},
	)

	//Wall
	newEntity(none,
		&transformComponent{x: 50, y: 150},
		&sprComponent{ani: [10]int{68}, aniLen: 1, opt: golf.SOp{H: 2, W: 4}},
		&collideComponent{width: 32, height: 16, oldX: 50, oldY: 150},
	)

	// for i := 0; i < 15; i++ {
	// 	addEnemy(float64(rand.Intn(192)), float64(rand.Intn(192)))
	// }

	// Player Movement
	allUpdateSystems[movePlayer] = toSystem(playerControlled, TypeTransformComponent, func(e *entity) {
		spr := sprComponents[e.id]
		tran := transformComponents[e.id]
		collide := collideComponents[e.id]
		speed := 1.5

		spr.ani = [10]int{2, 3}
		spr.aniLen = 2
		spr.aniSpeed = 60
		spr.opt.W = 1
		collide.width = 8
		if g.Btn(golf.WKey) {
			tran.y -= speed
			spr.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
			spr.aniLen = 8
			spr.aniSpeed = 10
			spr.opt.W = 2
			collide.width = 16
		}
		if g.Btn(golf.SKey) {
			tran.y += speed
			spr.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
			spr.aniLen = 8
			spr.aniSpeed = 10
			spr.opt.W = 2
			collide.width = 16
		}
		if g.Btn(golf.AKey) {
			tran.x -= speed
			spr.ani = [10]int{6, 8, 10, 12}
			spr.aniLen = 4
			spr.aniSpeed = 10
			spr.opt.FH = false
			spr.opt.W = 2
			collide.width = 16
		}
		if g.Btn(golf.DKey) {
			tran.x += speed
			spr.ani = [10]int{6, 8, 10, 12}
			spr.aniLen = 4
			spr.aniSpeed = 10
			spr.opt.FH = true
			spr.opt.W = 2
			collide.width = 16
		}
	})

	// Do AI
	allUpdateSystems[doAI] = toSystem(0, TypeAIComponent|TypeTransformComponent|TypeSprComponent, func(e *entity) {
		ai := aiComponents[e.id]
		tran := transformComponents[e.id]
		spr := sprComponents[e.id]

		speed := ai.speed
		targetTran := transformComponents[ai.target]
		dist := (((targetTran.x - tran.x) * (targetTran.x - tran.x)) + ((targetTran.y - tran.y) * (targetTran.y - tran.y)))
		if dist > ai.atkRange*ai.atkRange {
			return
		}

		if tran.x < targetTran.x {
			tran.x += speed
			spr.opt.FH = true
		}
		if tran.x > targetTran.x {
			tran.x -= speed
			spr.opt.FH = false
		}
		if tran.y < targetTran.y {
			tran.y += speed
		}
		if tran.y > targetTran.y {
			tran.y -= speed
		}
	})

	// AI attack
	allUpdateSystems[doAttack] = toSystem(enemy, TypeAIComponent|TypeTransformComponent, func(e *entity) {
		tran := transformComponents[e.id]
		ai := aiComponents[e.id]

		target := allEntities[ai.target]
		targetTran := transformComponents[ai.target]
		targetHP := hpComponents[ai.target]
		hDist := targetTran.x - tran.x
		vDist := targetTran.y - tran.y
		if hDist < 0 {
			hDist *= -1
		}
		if vDist < 0 {
			vDist *= -1
		}
		if hDist < 3 && vDist < 3 && target.hasComponent(TypeHPComponent) {
			targetHP.health--
		}
	})

	// Fill the list of collidable entities
	allUpdateSystems[startCollision] = toSystem(0, TypeTransformComponent|TypeCollideComponent, func(e *entity) {
		allCollidables[collidablePointer] = e.id
		collidablePointer++
	})

	// Calculate the collisions
	allUpdateSystems[doCollision] = toSystem(0, TypeTransformComponent|TypeCollideComponent, doPhysics)

	// Finish the collision detection
	allUpdateSystems[resolveCollision] = toSystem(0, TypeTransformComponent|TypeCollideComponent, func(e *entity) {
		tran := transformComponents[e.id]
		collide := collideComponents[e.id]

		tran.x = collide.oldX + collide.deltaX
		tran.y = collide.oldY + collide.deltaY
		collide.oldX = tran.x
		collide.oldY = tran.y
		collidablePointer = 0
	})

	// Draw all the sprites
	allDrawSystems[drawSprite] = toSystem(0, TypeTransformComponent|TypeSprComponent, func(e *entity) {
		spr := sprComponents[e.id]
		tran := transformComponents[e.id]

		spr.frame++
		if spr.aniSpeed == 0 {
			spr.aniSpeed = 1
		}
		if spr.frame%spr.aniSpeed == 0 {
			spr.aniFrame++
		}
		if spr.aniFrame >= spr.aniLen {
			spr.aniFrame = 0
		}
		frame := spr.ani[spr.aniFrame]
		opt := spr.opt
		if frame < 0 {
			opt.FH = true
			frame *= -1
		}
		g.Spr(frame, tran.x, tran.y, opt)
		if e.hasComponent(TypeCollideComponent) {
			collide := collideComponents[e.id]
			g.Rect(tran.x, tran.y, collide.width, collide.height, golf.Col1)
		}
	})

	// Draw HP hud element
	allDrawSystems[drawHP] = toSystem(playerControlled, TypeHPComponent, func(e *entity) {
		hp := hpComponents[e.id]

		if hp.health < 0 {
			g.TextL("You DED!", golf.TOp{Col: golf.Col3})
			return
		}
		g.TextL(fmt.Sprintf("HP: %d", hp.health), golf.TOp{Col: golf.Col3})
	})
}

func update() {
	g.PalA(12)
	g.PalB(1)
	runUpdateSystems()
}

func draw() {
	g.Cls()
	runDrawSystems()
	a := collideComponents[allCollidables[0]]
	b := collideComponents[allCollidables[1]]
	x1, y1 := a.oldX+a.deltaX, a.oldY+a.deltaY
	w1, h1 := a.width, a.height
	x2, y2 := b.oldX+b.deltaX, b.oldY+b.deltaY
	w2, h2 := b.width, b.height
	debug += fmt.Sprintf("\nx1 %d, y1 %d, w1 %d, h1 %d\nx2 %d, y2 %d, w1 %d, h1 %d", int(x1), int(y1), int(w1), int(h1), int(x2), int(y2), int(w2), int(h2))
	var xO1, yO1, xO2, yO2, xO3, yO3 bool
	if x1 >= x2 && x1 <= x2+h2 {
		xO1 = true
	}
	if x1+h1 >= x2 && x1+h1 <= x2+h2 {
		xO2 = true
	}
	if x1 >= x2 && x1+h1 <= x2+h2 {
		xO3 = true
	}
	if y1 >= y2 && y1 <= y2+h2 {
		yO1 = true
	}
	if y1+h1 >= y2 && y1+h1 <= y2+h2 {
		yO2 = true
	}
	if y1 >= y2 && y1+h1 <= y2+h2 {
		yO3 = true
	}
	debug += fmt.Sprintf("\nxO1: %v, yO1: %v,\nxO2: %v, yO2: %v\nxO3: %v, yO3: %v", xO1, yO1, xO2, yO2, xO3, yO3)

	g.TextR(debug, golf.TOp{Col: golf.Col3})
	debug = ""
}

func addEnemy(x, y float64) *entity {
	// zombie
	return newEntity(enemy,
		&transformComponent{x: x, y: y},
		&sprComponent{ani: [10]int{24, 25, 26}, aniLen: 3, aniSpeed: 30, opt: golf.SOp{H: 2, TCol: golf.Col7}},
		&aiComponent{atkRange: 150, target: player.id, speed: 0.25},
		&collideComponent{width: 8, height: 16, oldX: x, oldY: y},
	)
}
