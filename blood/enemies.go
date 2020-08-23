package main

import (
	"fantasyConsole/golf"
	"math/rand"
)

var totalEnemyCount int
var maxEnemyCount = 20

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
		if g.Frames()%12000 == 0 && maxEnemyCount < 100 {
			maxEnemyCount++
		}
		if totalEnemyCount > maxEnemyCount {
			return
		}
		tran := transformComponents[e.id]
		if rand.Intn(100) == 1 {
			addZombie(tran.x+100, tran.y+float64(rand.Intn(31)-15))
			totalEnemyCount++
		}
		if rand.Intn(100) == 1 {
			addZombie(tran.x-100, tran.y+float64(rand.Intn(31)-15))
			totalEnemyCount++
		}
	})

	// Kill Enemies will < 0 hp
	allUpdateSystems[doEnemyDeath] = toSystem(enemy, TypeHPComponent, func(e *entity) {
		hp := hpComponents[e.id]
		if hp.health > 0 {
			return
		}

		for i, zomb := range allEnemies {
			if zomb.id == e.id {
				if e.hasComponent(TypeTransformComponent) {
					zXY := transformComponents[e.id]
					for i := 0; i < 20; i++ {
						addBloodParticle(
							bloodParticles,
							rand.Float64()*16+(zXY.x-8),
							rand.Float64()*8+(zXY.y+12),
							rand.Float64()-0.5,
							rand.Float64()*5,
							float64(rand.Intn(10)+4),
						)
					}
				}

				allEnemies[i] = allEnemies[len(allEnemies)-1]
				allEnemies = allEnemies[:len(allEnemies)-1]
				deleteEntity(e)
				totalEnemyCount--
				return
			}
		}
	})

	// Draw above the head hp
	allDrawSystems[drawMiniHP] = toSystem(none, TypeHPComponent|TypeTransformComponent, func(e *entity) {
		if e.hasFlag(playerControlled) {
			return
		}
		hp := hpComponents[e.id]
		hpLen := (float64(hp.health) / float64(hp.maxHealth)) * 8
		tran := transformComponents[e.id]
		g.Line(tran.x-3, tran.y-3, tran.x+hpLen, tran.y-3, golf.Col1)
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
		&hpComponent{health: 10, maxHealth: 10},
	)

	allEnemies = append(allEnemies, ret)
	return ret
}
