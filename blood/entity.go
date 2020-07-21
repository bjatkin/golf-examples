package main

import (
	"fantasyConsole/golf"
)

var allEntities = [100]entity{} // is 100 to many? not enough?
var entityPointer = 0

type entity struct {
	id      int
	flags   flag
	sig     component
	spr     sprComp
	pos     posComp
	text    string
	ai      aiComp
	hp      int
	collide collidableComp
}

func (e *entity) hasFlag(f flag) bool {
	return (e.flags & f) == f
}

func (e *entity) hasComponent(c component) bool {
	return (e.sig & c) == c
}

func addEntity(e entity) int {
	allEntities[entityPointer] = e
	e.id = entityPointer
	entityPointer++
	return entityPointer - 1
}

type component int

const (
	pos     = component(1)
	spr     = component(2)
	text    = component(4)
	ai      = component(8)
	hp      = component(16)
	collide = component(32)
)

type posComp struct {
	x, y float64
}

type sprComp struct {
	ani      [10]int
	aniLen   int
	aniSpeed int
	opt      golf.SOp
	aniFrame int
	frame    int
}

type aiComp struct {
	atkRange float64
	target   int
}

type collidableComp struct {
	oldX, oldY     float64
	deltaX, deltaY float64
	width, height  float64
}

type flag int

const (
	playerControlled = flag(1)
	enemy            = flag(2)
)
