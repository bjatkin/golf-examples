package main

import (
	"fantasyConsole/golf"
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

	// Do Iframes
	allUpdateSystems[doIFrames] = toSystem(none, TypeHPComponent, func(e *entity) {
		hpComponents[e.id].iFrames--
	})

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
		})

	// Draw Player hud element
	allDrawSystems[drawHUD] = toSystem(playerControlled, TypeHPComponent|TypeBloodBankComponent|TypeCooldownComponent, func(e *entity) {
		hp := hpComponents[e.id]
		pBank := bloodBankComponents[e.id]
		cool := cooldownComponents[e.id]

		if hp.health < 0 {
			g.Update = deathScreenUpdate
			g.Draw = deathScreenDraw
			return
		}

		if pBank.balance > 1000 {
			g.Update = winScreenUpdate
			g.Draw = winScreenDraw
			return
		}

		// Draw the HP
		g.RectFill(0, 160, 40, 32, golf.Col0, true) // BG
		hpPercent := 1 - (float64(hp.health) / float64(hp.maxHealth))
		g.RectFill(0, 160+(32*hpPercent), 40, 32, golf.Col1, true) // HP level
		g.Spr(384, 0, 160, golf.SOp{TCol: golf.Col2, W: 5, H: 4, Fixed: true})

		// Draw the blood bank
		g.RectFill(151, 160, 40, 32, golf.Col0, true) // BG
		bloodPercent := 1 - (float64(pBank.balance) / 1000.0)
		g.RectFill(152, 160+(32*bloodPercent), 40, 32, golf.Col2, true) // blood bank level
		g.Spr(384, 151, 160, golf.SOp{TCol: golf.Col2, W: 5, H: 4, Fixed: true, FH: true})

		// Cool Down Rect
		attack1Percent := 1 - (float64(cool.cooldown1) / float64(cool.reset1))
		attack2Percent := 1 - (float64(cool.cooldown2) / float64(cool.reset2))
		g.RectFill(70, 168, 56, 24, golf.Col0, true)                     // BG
		g.RectFill(76, 168, 18-(18*attack1Percent), 24, golf.Col2, true) // attack 1
		g.RectFill(98, 168, 18-(18*attack2Percent), 24, golf.Col2, true) // attack 2
		g.Spr(422, 70, 168, golf.SOp{TCol: golf.Col2, W: 7, H: 3, Fixed: true})
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
