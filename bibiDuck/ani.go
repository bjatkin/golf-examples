package main

import "fantasyConsole/golf"

type aniSprite struct {
	id     int
	frames []int
	frame  float64
	speed  float64
	x, y   float64
	o      golf.SOp
}

func (s *aniSprite) draw() {
	g.Spr(s.frames[int(s.frame)], s.x, s.y, s.o)
}

func (s *aniSprite) tick() {
	s.frame += s.speed
	if s.frame > float64(len(s.frames)) {
		s.frame = 0
	}
}

var allAni = []*aniSprite{}

func drawAniSprites() {
	for _, ani := range allAni {
		ani.draw()
		ani.tick()
	}
}

func newFether(x, y float64) *aniSprite {
	ret := &aniSprite{
		id:     len(allAni),
		frames: []int{32, 64, 96, 128, 160, 128, 96, 64},
		speed:  0.1,
		x:      x,
		y:      y,
		o:      alpha,
	}
	allAni = append(allAni, ret)
	return ret
}

func newEgg(x, y float64) *aniSprite {
	ret := &aniSprite{
		id:     len(allAni),
		frames: []int{193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 195, 197, 195, 197, 195, 197},
		speed:  0.1,
		x:      x,
		y:      y,
		o:      golf.SOp{W: 2, H: 2, TCol: golf.Col5},
	}
	allAni = append(allAni, ret)
	return ret
}

type multiAniSprite struct {
	id     int
	states []aniSprite
	state  int
}

func (ms *multiAniSprite) setState(state int) {
	ms.state = state
	allAni[ms.id] = &ms.states[state]
}
