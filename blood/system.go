package main

// All the update systems
const (
	movePlayer = iota
	doAI
	doAttack
	startCollision
	doCollision
	resolveCollision
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
