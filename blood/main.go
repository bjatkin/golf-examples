package main

import (
	"math"

	"github.com/bjatkin/golf-engine/golf"
)

var g *golf.Engine

func main() {
	g = golf.NewEngine(menuUpdate, menuDraw)

	g.LoadSprs(spriteSheet)
	g.LoadMap(mapData)

	initGame()
	g.Run()
}

var player *entity
var debug string

func initGame() {
	initPlayer()
	initEnemies()
	initParticleSystem()
	initProjectileSystem()

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

	// Resolve collisons with collision meshes
	allUpdateSystems[doCollision] = toSystem(none, TypeTransformComponent|TypeSolidComponent, func(e *entity) {
		t := transformComponents[e.id]
		s := solidComponents[e.id]
		for _, e := range allEntities {
			if e.hasComponent(TypeCollisionMeshComponent) {
				m := collisionMeshComponents[e.id]

				if t.x < m.x+m.w && t.x+s.w > m.x &&
					t.y < m.y+m.h && t.y+s.h > m.y {
					xOff, yOff := 0.0, 0.0
					if t.x <= m.x { // shift to the left
						xOff = m.x - (t.x + s.w)
					}
					if t.x >= m.x { // shift to the right
						xOff = (m.x + m.w) - t.x
					}
					if t.y <= m.y { // shift up
						yOff = m.y - (t.y + s.h)
					}
					if t.y >= m.y { // shift down
						yOff = (m.y + m.h) - t.y
					}
					if math.Abs(xOff) < math.Abs(yOff) {
						t.x += xOff
					} else {
						t.y += yOff
					}
				}
			}
		}
	})

	// Do all the nessisary logic to step the animation
	allUpdateSystems[doSprAni] = toSystem(none, TypeAniComponent, func(e *entity) {
		ani := aniComponents[e.id]

		ani.aniFrame += ani.aniSpeed
		if int(ani.aniFrame) >= ani.aniLen {
			ani.aniFrame = 0
		}

		//update the spr component if there is one
		if e.hasComponent(TypeSprComponent) {
			spr := sprComponents[e.id]
			n := ani.ani[int(ani.aniFrame)]
			spr.opt.FH = false
			if n < 0 {
				spr.opt.FH = true
				n *= -1
			}
			spr.n = n
		}
	})

	// Draw all the sprites
	allDrawSystems[drawSprite] = toSystem(none, TypeTransformComponent|TypeSprComponent, func(e *entity) {
		spr := sprComponents[e.id]
		tran := transformComponents[e.id]

		g.Spr(spr.n, tran.x, tran.y, spr.opt)
	})

}

func gameUpdate() {
	runUpdateSystems()
}

var whiteTxt = golf.TOp{Col: golf.Col3}
var cameraX, cameraY = 0, 0

func gameDraw() {
	g.Cls(golf.Col0)
	g.Map(40, 20, 40, 70, 0, 0)

	pos := transformComponents[player.id]
	cameraX, cameraY = int(pos.x)-92, int(pos.y)-88
	g.Camera(cameraX, cameraY)

	runDrawSystems()

	g.TextR(debug, whiteTxt)
	debug = ""
}

func lerp(a, b, beta float64) float64 {
	return a + (b-a)*beta
}
