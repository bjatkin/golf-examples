package main

// All the update systems
const (
	movePlayer = iota
	doAI
	doAttack
	physicsRound0
	physicsRound1
	physicsRound2
	physicsRound3
	physicsRound4
	physicsRound5
	physicsRound6
	physicsRound7
	physicsRound8
	physicsRound9
	resolvePhysics
	updateSystemLen
)

// All the draw systems
const (
	drawSprite = iota
	drawHP
	drawSystemLen
)

var allUpdateSystems = [updateSystemLen]system{}
var allDrawSystems = [drawSystemLen]system{}

type system func(*entity)

func toSystem(f flag, c component, sys func(*entity)) system {
	return func(e *entity) {
		if (e.flags&f) == f && (e.sig&c) == c {
			sys(e)
		}
	}
}

func runUpdateSystems() {
	for _, sys := range allUpdateSystems {
		for ent := 0; ent < entityPointer; ent++ {
			sys(&allEntities[ent])
		}
	}
}

func runDrawSystems() {
	for _, sys := range allDrawSystems {
		for ent := 0; ent < entityPointer; ent++ {
			sys(&allEntities[ent])
		}
	}
}
