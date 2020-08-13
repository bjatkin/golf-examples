package main

import (
	"fantasyConsole/golf"
	"fmt"
	"math/rand"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(update, draw)

	g.LoadSprs(spriteSheet)
	g.LoadMap(mapData)

	initGame()
	g.Run()
}

var player *entity
var debug string
var totalEnemyCount int

func initGame() {
	player = newEntity(playerControlled,
		&hpComponent{health: 100, maxHealth: 100},
		&transformComponent{x: 192, y: 480},
		&sprComponent{ani: [10]int{2, 3}, aniLen: 2, aniSpeed: 60, opt: golf.SOp{H: 2, TCol: golf.Col2}},
		&solidComponent{w: 8, h: 16},
	)

	// Collision walls
	newEntity(none,
		&collisionMeshComponent{0, 115, 153, 170, 112},
	)
	newEntity(none,
		&collisionMeshComponent{0, 115, 265, 69, 70},
	)
	newEntity(none,
		&collisionMeshComponent{0, 216, 265, 69, 151},
	)
	newEntity(none,
		&collisionMeshComponent{0, 208, 400, 8, 16},
	)
	newEntity(none,
		&collisionMeshComponent{0, 115, 400, 69, 16},
	)
	newEntity(none,
		&collisionMeshComponent{0, 115, 335, 50, 65},
	)

	// Player Movement
	allUpdateSystems[movePlayer] = toSystem(playerControlled, TypeTransformComponent, func(e *entity) {
		spr := sprComponents[e.id]
		tran := transformComponents[e.id]
		speed := 1.5

		spr.ani = [10]int{2, 3}
		spr.aniLen = 2
		spr.aniSpeed = 60
		spr.opt.W = 1
		if g.Btn(golf.WKey) {
			tran.y -= speed
			spr.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
			spr.aniLen = 8
			spr.aniSpeed = 10
			spr.opt.W = 2
		}
		if g.Btn(golf.SKey) {
			tran.y += speed
			spr.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
			spr.aniLen = 8
			spr.aniSpeed = 10
			spr.opt.W = 2
		}
		if g.Btn(golf.AKey) {
			tran.x -= speed
			spr.ani = [10]int{6, 8, 10, 12}
			spr.aniLen = 4
			spr.aniSpeed = 10
			spr.opt.FH = false
			spr.opt.W = 2
		}
		if g.Btn(golf.DKey) {
			tran.x += speed
			spr.ani = [10]int{6, 8, 10, 12}
			spr.aniLen = 4
			spr.aniSpeed = 10
			spr.opt.FH = true
			spr.opt.W = 2
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
		if dist > ai.atkRange*ai.atkRange || dist < 9 {
			return
		}

		dx, dy := 0.0, 0.0
		rMove := rand.Intn(12)
		if rMove == 1 {
			tran.x += speed
			spr.opt.FH = true
		}
		if rMove == 2 {
			tran.x -= speed
			spr.opt.FH = false
		}
		if rMove == 3 {
			tran.y += speed
		}
		if rMove == 4 {
			tran.y -= speed
		}

		if tran.x < targetTran.x {
			dx = speed
		}
		if tran.x > targetTran.x {
			dx = -speed
		}
		if tran.y < targetTran.y {
			dy = speed
		}
		if tran.y > targetTran.y {
			dy = -speed
		}

		if dx != 0 && dy != 0 {
			if rand.Intn(2) == 1 {
				dx = 0
			} else {
				dy = 0
			}
		}

		tran.x += dx
		tran.y += dy
		if dx > 0 {
			spr.opt.FH = true
		}
		if dx < 0 {
			spr.opt.FH = false
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
		if hDist < 3 && vDist < 3 && target.hasComponent(TypeHPComponent) && targetHP.iFrames <= 0 {
			targetHP.health -= 10
			targetHP.iFrames = 60

			// Add knockback
			targetTran.x += hDist * 2
			targetTran.y += vDist * 2
		}
	})

	// Do Iframes
	allUpdateSystems[doIFrames] = toSystem(none, TypeHPComponent, func(e *entity) {
		hpComponents[e.id].iFrames--
	})

	// Add Enemies
	allUpdateSystems[addEnemies] = toSystem(playerControlled, TypeTransformComponent, func(e *entity) {
		if totalEnemyCount > 20 {
			return
		}
		tran := transformComponents[e.id]
		if rand.Intn(500) == 1 {
			addZombie(tran.x+100, tran.y+float64(rand.Intn(31)-15))
			totalEnemyCount++
		}
		if rand.Intn(500) == 1 {
			addZombie(tran.x-100, tran.y+float64(rand.Intn(31)-15))
			totalEnemyCount++
		}
	})

	// Resolve collisons with collision meshes
	allUpdateSystems[doCollision] = toSystem(none, TypeTransformComponent|TypeSolidComponent, func(e *entity) {
		t := transformComponents[e.id]
		s := solidComponents[e.id]
		for _, e := range allEntities {
			if e.hasComponent(TypeCollisionMeshComponent) {
				m := collisionMeshComponents[e.id]

				if t.x < m.x+m.w && t.x+s.w > m.x &&
					t.y < m.y+m.h && t.y+s.h > m.y {
					//TODO there was a collision here that needs to be handled
					debug += "collision!\n"
				}
			}
		}
	})

	// Draw all the sprites
	allDrawSystems[drawSprite] = toSystem(none, TypeTransformComponent|TypeSprComponent, func(e *entity) {
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
	})

	// Draw HP hud element
	allDrawSystems[drawHP] = toSystem(playerControlled, TypeHPComponent|TypeTransformComponent, func(e *entity) {
		hp := hpComponents[e.id]

		if hp.health < 0 {
			g.TextL("You DED!", whiteTxt)
			return
		}
		g.RectFill(0, 0, 192, 8, golf.Col0, true)
		g.TextL("HP:", whiteTxt)
		maxLen := 50.0
		g.RectFill(20, 1, (float64(hp.health)/float64(hp.maxHealth))*maxLen, 6, golf.Col1, true)

		pos := transformComponents[e.id]
		g.TextR(fmt.Sprintf("X: %.0f, Y: %.0f", pos.y, pos.y), whiteTxt)
	})

	// Draw Collision Meshes for debugging
	allDrawSystems[tmp] = toSystem(none, TypeCollisionMeshComponent, func(e *entity) {
		mesh := collisionMeshComponents[e.id]
		g.Rect(mesh.x, mesh.y, mesh.w, mesh.h, golf.Col1)
	})
}

func update() {
	g.PalA(15)
	g.PalB(1)
	runUpdateSystems()
}

var whiteTxt = golf.TOp{Col: golf.Col3}

func draw() {
	g.Cls(golf.Col0)
	g.Map(40, 20, 40, 70, 0, 0)

	pos := transformComponents[player.id]
	g.Camera(int(pos.x)-92, int(pos.y)-88)

	runDrawSystems()

	g.TextR(debug, whiteTxt)
	debug = ""
}

func addZombie(x, y float64) *entity {
	return newEntity(enemy,
		&transformComponent{x: x, y: y},
		&sprComponent{ani: [10]int{24, 25, 26}, aniLen: 3, aniSpeed: 30, opt: golf.SOp{H: 2, TCol: golf.Col2}},
		&aiComponent{atkRange: 150, target: player.id, speed: 0.25},
		&solidComponent{w: 8, h: 16},
	)
}
