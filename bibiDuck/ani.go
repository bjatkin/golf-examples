package main

import (
	"fantasyConsole/golf"
)

type drawable interface {
	draw(float64, float64)
}

type spr struct {
	n int
	o golf.SOp
}

func (s *spr) draw(x, y float64) {
	g.Spr(s.n, x, y, s.o)
}

type ani struct {
	frames   []int
	frame    float64
	speed    float64
	o        golf.SOp
	onlyOnce bool
}

func (s *ani) draw(x, y float64) {
	g.Spr(s.frames[int(s.frame)], x, y, s.o)

	s.frame += s.speed
	if s.frame >= float64(len(s.frames)) {
		s.frame = float64(len(s.frames) - 1)
		if !s.onlyOnce {
			s.frame = 0
		}
	}
}

var allFethers = []*sprite{}

func newFether(x, y float64) *sprite {
	ret := &sprite{
		x: x,
		y: y,
		a: &ani{
			frames: []int{32, 64, 96, 128, 160, 128, 96, 64},
			speed:  0.1,
			o:      alpha,
		},
	}

	allFethers = append(allFethers, ret)
	ret.init()
	return ret
}

var allEggs = []*sprite{}

func newEgg(x, y float64) *sprite {
	ret := &sprite{
		x: x,
		y: y,
		a: &stateAni{
			states: []drawable{
				&ani{ // Shake
					frames: []int{193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 195, 197, 195, 197, 195, 197},
					speed:  0.1,
					o:      golf.SOp{W: 2, H: 2, TCol: golf.Col5},
				},
				&ani{ // Collected
					frames:   []int{193, 199, 75, 199, 75, 199, 75, 199, 75},
					speed:    0.1,
					o:        golf.SOp{W: 2, H: 2, TCol: golf.Col5},
					onlyOnce: true,
				},
			},
		},
	}

	allEggs = append(allEggs, ret)
	ret.init()
	return ret
}

type stateAni struct {
	states []drawable
	state  int
}

func (s *stateAni) draw(x, y float64) {
	if s.state >= len(s.states) {
		return
	}
	s.states[s.state].draw(x, y)
}

func newMob(x, y float64, states []drawable) *sprite {
	ret := &sprite{
		x: x,
		y: y,
		a: &stateAni{
			states: states,
		},
	}
	ret.init()
	return ret
}

var allSprites = []*sprite{}

type sprite struct {
	id   int
	x, y float64
	a    drawable
}

func (s *sprite) init() {
	s.id = len(allSprites)
	allSprites = append(allSprites, s)
}

func (s *sprite) delete() {
	allSprites[s.id] = allSprites[len(allSprites)-1]
	allSprites[s.id].id = s.id
	allSprites = allSprites[:len(allSprites)-1]
}

func drawSprites() {
	for _, sprite := range allSprites {
		sprite.a.draw(sprite.x, sprite.y)
	}
}
