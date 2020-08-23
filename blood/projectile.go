package main

import (
	"math"

	"github.com/bjatkin/golf-engine/golf"
)

func initProjectileSystem() {
	allUpdateSystems[doProjectile] = toSystem(
		projectile,
		TypeTransformComponent|TypeTravelComponent,
		func(e *entity) {
			t := transformComponents[e.id]
			delta := travelComponents[e.id]
			t.x += delta.dx
			t.y += delta.dy
			if int(t.x+8) < cameraX || int(t.y+8) < cameraY ||
				int(t.x) > cameraX+192 || int(t.y) > cameraY+192 {
				deleteEntity(e)
				return
			}
			for _, zomb := range allEnemies {
				z := getEntity(zomb.id)
				zt := transformComponents[z.id]
				hp := hpComponents[z.id]
				dx, dy := (zt.x+4)-t.x, (zt.y+8)-t.y
				if math.Abs(dx) < 4 && math.Abs(dy) < 8 {
					hp.health--
					if dx > 0 {
						zt.x += 2
					}
					if dy > 0 {
						zt.y += 2
					}
					if dx < 0 {
						zt.x -= 2
					}
					if dy < 0 {
						zt.y -= 2
					}
					deleteEntity(e)
					return
				}
			}
		})

	allUpdateSystems[tickCooldown] = toSystem(
		none,
		TypeCooldownComponent,
		func(e *entity) {
			cool := cooldownComponents[e.id]
			cool.cooldown1--
			cool.cooldown2--
		},
	)
}

const (
	bigUp = iota
	bigDown
	bigLeft
	bigRight
	bigUpLeft
	bigUpRight
	bigDownLeft
	bigDownRight
	smallUp
	smallDown
	smallLeft
	smallRight
	small135
	small45
	small225
	small315
	small150
	small120
	small60
	small30
	small210
	small240
	small300
	small330
)

func newProjectile(x, y, dx, dy float64, ptype int) *entity {
	n := 82
	o := golf.SOp{TCol: golf.Col2}
	switch ptype {
	case bigUp:
		o = golf.SOp{TCol: golf.Col2, FV: true}
	case bigLeft:
		o = golf.SOp{TCol: golf.Col2, FH: true}
		n = 114
	case bigRight:
		n = 114
	case bigUpLeft:
		o = golf.SOp{TCol: golf.Col2, FV: true, FH: true}
		n = 83
	case bigUpRight:
		o = golf.SOp{TCol: golf.Col2, FV: true}
		n = 83
	case bigDownLeft:
		o = golf.SOp{TCol: golf.Col2, FH: true}
		n = 83
	case bigDownRight:
		n = 83
	case smallUp:
		n = 84
	case smallDown:
		o = golf.SOp{TCol: golf.Col2, FV: true}
		n = 84
	case smallLeft:
		n = 116
	case smallRight:
		o = golf.SOp{TCol: golf.Col2, FH: true}
		n = 116
	case small135:
		n = 115
	case small45:
		o = golf.SOp{TCol: golf.Col2, FV: true}
		n = 115
	case small225:
		o = golf.SOp{TCol: golf.Col2, FH: true}
		n = 115
	case small315:
		o = golf.SOp{TCol: golf.Col2, FH: true, FV: true}
		n = 115
	}
	return newEntity(projectile,
		&sprComponent{n: n, opt: o},
		&transformComponent{x: x, y: y},
		&travelComponent{dx: dx, dy: dy},
	)
}
