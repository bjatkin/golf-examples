package main

import (
	"fantasyConsole/golf"
)

type compType int

// all the transform types
const (
	TypeTransformComponent = compType(1)
	TypeSprComponent       = compType(2)
	TypeTextComponent      = compType(4)
	TypeAIComponent        = compType(8)
	TypeHPComponent        = compType(16)
	TypeCollideComponent   = compType(32)
)

type component interface {
	setEnt(*entity)
	ctype() compType
	add()
}

var transformComponents = map[int]*transformComponent{}

type transformComponent struct {
	id   int
	x, y float64
}

func (c *transformComponent) setEnt(ent *entity) {
	c.id = ent.id
}
func (c *transformComponent) ctype() compType {
	return TypeTransformComponent
}
func (c *transformComponent) add() {
	transformComponents[c.id] = c
}

var sprComponents = map[int]*sprComponent{}

type sprComponent struct {
	id       int
	ani      [10]int
	aniLen   int
	aniSpeed int
	opt      golf.SOp
	aniFrame int
	frame    int
}

func (c *sprComponent) setEnt(ent *entity) {
	c.id = ent.id
}
func (c *sprComponent) ctype() compType {
	return TypeSprComponent
}
func (c *sprComponent) add() {
	sprComponents[c.id] = c
}

var aiComponents = map[int]*aiComponent{}

type aiComponent struct {
	id       int
	atkRange float64
	target   int
	speed    float64
}

func (c *aiComponent) setEnt(ent *entity) {
	c.id = ent.id
}
func (c *aiComponent) ctype() compType {
	return TypeAIComponent
}
func (c *aiComponent) add() {
	aiComponents[c.id] = c
}

var collideComponents = map[int]*collideComponent{}

type collideComponent struct {
	id             int
	oldX, oldY     float64
	deltaX, deltaY float64
	width, height  float64
}

func (c *collideComponent) setEnt(ent *entity) {
	c.id = ent.id
}
func (c *collideComponent) ctype() compType {
	return TypeCollideComponent
}
func (c *collideComponent) add() {
	collideComponents[c.id] = c
}

var hpComponents = map[int]*hpComponent{}

type hpComponent struct {
	id     int
	health int
}

func (c *hpComponent) setEnt(ent *entity) {
	c.id = ent.id
}
func (c *hpComponent) ctype() compType {
	return TypeHPComponent
}
func (c *hpComponent) add() {
	hpComponents[c.id] = c
}

var textComponents = map[int]*textComponent{}

type textComponent struct {
	id   int
	text string
}

func (c *textComponent) setEnt(ent *entity) {
	c.id = ent.id
}
func (c *textComponent) ctype() compType {
	return TypeHPComponent
}
func (c *textComponent) add() {
	textComponents[c.id] = c
}

type nilComponent struct{}

func (c *nilComponent) setEnt(ent *entity) {}
func (c *nilComponent) ctype() compType    { return 0 }
func (c *nilComponent) add()               {}