package main

import (
	"fantasyConsole/golf"
	"math/rand"
)

func initEnemies() {
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

		target := getEntity(ai.target)
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
}

var allEnemies = []*entity{}

func addZombie(x, y float64) *entity {
	ret := newEntity(enemy,
		&transformComponent{x: x, y: y},
		&aniComponent{ani: [10]int{24, 25, 26}, aniLen: 3, aniSpeed: 1.0 / 30.0},
		&sprComponent{opt: golf.SOp{H: 2, TCol: golf.Col2}},
		&aiComponent{atkRange: 150, target: player.id, speed: 0.25},
		&solidComponent{w: 8, h: 16},
		&hpComponent{health: 30, maxHealth: 30},
	)

	allEnemies = append(allEnemies, ret)
	return ret
}
