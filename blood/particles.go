package main

import (
	"fantasyConsole/golf"
	"math"
)

type blood struct {
	x, y      float64
	dx, dy    float64
	height    float64
	collected bool
	live      bool
}

func addBloodParticle(e *entity, x, y, dx, dy, height float64) {
	if !e.hasComponent(TypeParticleComponent) {
		return
	}

	pList := particleComponents[e.id]
	pList.particles[pList.pid].x = x
	pList.particles[pList.pid].y = y
	pList.particles[pList.pid].dx = dx
	pList.particles[pList.pid].dy = dy
	pList.particles[pList.pid].collected = false
	pList.particles[pList.pid].live = true
	pList.particles[pList.pid].height = height

	pList.pid++
	if pList.pid > len(pList.particles)-1 {
		pList.pid = 0
	}
}

var bloodParticles *entity

func initParticleSystem() {
	bloodParticles = newEntity(none,
		&particleComponent{},
	)
	p := particleComponents[bloodParticles.id]
	for i := 0; i < len(p.particles); i++ {
		p.particles[i] = &blood{}
	}

	allUpdateSystems[doParticles] = toSystem(none, TypeParticleComponent, func(e *entity) {
		pList := particleComponents[e.id]
		pXY := transformComponents[player.id]
		pBank := bloodBankComponents[player.id]
		for _, p := range pList.particles {
			if !p.live {
				continue
			}
			dx, dy := math.Abs((pXY.x+4)-p.x), math.Abs((pXY.y+16)-p.y)
			if dx < 15 && dy < 15 && !p.collected {
				p.collected = true
				pBank.balance++
			}
			p.height += p.dy
			p.dy -= 0.2
			if p.collected {
				p.dy += 0.4
			}
			if p.height > 0 && p.height < 192 {
				p.x += p.dx
			}
			if p.height < 0 {
				p.dy = 0
				p.height = 0
			}
			if p.height > 192 {
				p.live = false
			}
		}
	})

	allDrawSystems[drawParticles] = toSystem(none, TypeParticleComponent, func(e *entity) {
		pList := particleComponents[e.id]
		for _, p := range pList.particles {
			if !p.live {
				continue
			}
			g.Pset(p.x-float64(cameraX), p.y-p.height-float64(cameraY), golf.Col1)
		}
	})
}
