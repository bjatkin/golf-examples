package main

import (
	"fantasyConsole/golf"
	"math/rand"
)

func initPlayer() {
	player = newEntity(playerControlled,
		&hpComponent{health: 100, maxHealth: 100},
		&transformComponent{x: 192, y: 480},
		&aniComponent{ani: [10]int{2, 3}, aniLen: 2, aniSpeed: (1 / 60)},
		&cooldownComponent{reset1: 5, reset2: 120},
		&sprComponent{opt: golf.SOp{H: 2, TCol: golf.Col2}},
		&solidComponent{w: 8, h: 16},
		&bloodBankComponent{balance: 0},
	)

	// Player Movement
	allUpdateSystems[movePlayer] = toSystem(
		playerControlled,
		TypeSprComponent|TypeAniComponent|TypeTransformComponent|TypeSolidComponent|TypeCooldownComponent,
		func(e *entity) {
			spr := sprComponents[e.id]
			ani := aniComponents[e.id]
			tran := transformComponents[e.id]
			solid := solidComponents[e.id]

			ani.ani = [10]int{2, 3}
			ani.aniLen = 2
			ani.aniSpeed = 1.0 / 60.0
			spr.opt.W = 1
			solid.w = 8
			attack := playerAttack(e, ani, spr, tran, solid)
			playerMove(e, ani, spr, tran, solid, attack)

			// TEST CODE
			if g.Btn(golf.PKey) {
				pXY := transformComponents[player.id]
				for i := 0; i < 4; i++ {
					addBloodParticle(
						bloodParticles,
						rand.Float64()*16+(pXY.x-32),
						rand.Float64()*8+(pXY.y+12),
						rand.Float64()-0.5,
						rand.Float64()*5,
						float64(rand.Intn(10)+4),
					)
				}
			}
			// TEST CODE
		})
}

func playerAttack(e *entity, ani *aniComponent, spr *sprComponent, tran *transformComponent, solid *solidComponent) bool {
	cool := cooldownComponents[e.id]
	if g.Btn(golf.ZKey) {
		ani.ani = [10]int{4}
		ani.aniLen = 1
		spr.opt.W = 2
		solid.w = 16
		if cool.cooldown2 < 0 {
			cool.cooldown2 = cool.reset2
			newProjectile(tran.x+4, tran.y+4, -6, 0, smallLeft)
			newProjectile(tran.x+4, tran.y+4, 0, -6, smallUp)
			newProjectile(tran.x+4, tran.y+4, 6, 0, smallRight)
			newProjectile(tran.x+4, tran.y+4, 0, 6, smallDown)

			newProjectile(tran.x+4, tran.y+4, -6, -6, small135)
			newProjectile(tran.x+4, tran.y+4, 6, -6, small45)
			newProjectile(tran.x+4, tran.y+4, -6, 6, small225)
			newProjectile(tran.x+4, tran.y+4, 6, 6, small315)
		}
		return true
	}
	if g.Btn(golf.XKey) {
		ptypeV, ptypeH := -1, -1
		dx, dy := 0.0, 0.0
		if g.Btn(golf.LeftArrow) {
			ani.ani = [10]int{76}
			ani.aniLen = 1
			spr.opt.W = 2
			solid.w = 16
			ptypeH = bigLeft
			dx = -6
		}
		if g.Btn(golf.RightArrow) {
			ani.ani = [10]int{-76}
			ani.aniLen = 1
			spr.opt.W = 2
			solid.w = 16
			ptypeH = bigRight
			dx = 6
		}
		if g.Btn(golf.UpArrow) {
			ani.ani = [10]int{80}
			ani.aniLen = 1
			spr.opt.W = 2
			solid.w = 16
			ptypeV = bigUp
			dy = -6
		}
		if g.Btn(golf.DownArrow) {
			ani.ani = [10]int{78}
			ani.aniLen = 1
			spr.opt.W = 2
			solid.w = 16
			ptypeV = bigDown
			dy = 6
		}

		if cool.cooldown1 < 0 && (ptypeV != -1 || ptypeH != -1) {
			cool.cooldown1 = cool.reset1
			if ptypeV == -1 {
				newProjectile(tran.x+4, tran.y+4, dx, dy, ptypeH)
			}
			if ptypeH == -1 {
				newProjectile(tran.x+4, tran.y+4, dx, dy, ptypeV)
			}
			if ptypeH == bigLeft && ptypeV == bigUp {
				newProjectile(tran.x+4, tran.y+4, dx, dy, bigUpLeft)
			}
			if ptypeH == bigLeft && ptypeV == bigDown {
				newProjectile(tran.x+4, tran.y+4, dx, dy, bigDownLeft)
			}
			if ptypeH == bigRight && ptypeV == bigUp {
				newProjectile(tran.x+4, tran.y+4, dx, dy, bigUpRight)
			}
			if ptypeH == bigRight && ptypeV == bigDown {
				newProjectile(tran.x+4, tran.y+4, dx, dy, bigDownRight)
			}
		}

		return true
	}
	return false
}

func playerMove(e *entity, ani *aniComponent, spr *sprComponent, tran *transformComponent, solid *solidComponent, attack bool) {
	speed := 1.5
	if g.Btn(golf.UpArrow) && !attack {
		tran.y -= speed
		ani.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
		ani.aniLen = 8
		ani.aniSpeed = 1.0 / 10.0
		spr.opt.W = 2
		solid.w = 16
	}
	if g.Btn(golf.DownArrow) && !attack {
		tran.y += speed
		ani.ani = [10]int{14, 16, 18, 20, -14, -16, -18, -20}
		ani.aniLen = 8
		ani.aniSpeed = 1.0 / 10.0
		spr.opt.W = 2
		solid.w = 16
	}
	if g.Btn(golf.LeftArrow) && !attack {
		tran.x -= speed
		ani.ani = [10]int{6, 8, 10, 12}
		ani.aniLen = 4
		ani.aniSpeed = 1.0 / 10.0
		spr.opt.W = 2
		solid.w = 16
	}
	if g.Btn(golf.RightArrow) && !attack {
		tran.x += speed
		ani.ani = [10]int{-6, -8, -10, -12}
		ani.aniLen = 4
		ani.aniSpeed = 1.0 / 10.0
		spr.opt.W = 2
		solid.w = 16
	}
}
