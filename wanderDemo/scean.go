package main

import "math"

type vec2 struct {
	x, y float64
}

var mainScean *scean

func exitScean(player vec2) (*scean, bool) {
	for s, pt := range mainScean.entrances {
		dx := math.Abs(player.x - pt.x)
		dy := math.Abs(player.y - pt.y)
		if dx < 3 && dy < 3 {
			return s, true
		}
	}
	return nil, false
}

type scean struct {
	mapXY     vec2
	mapWH     vec2
	cameraXY  vec2
	entrances map[*scean]vec2
	poi       []*interaction
}

var rentHouse = scean{
	mapXY:    vec2{52, 0},
	mapWH:    vec2{9, 9},
	cameraXY: vec2{0, 0},
}

var dogHouse = scean{
	mapXY:    vec2{57, 0},
	mapWH:    vec2{9, 9},
	cameraXY: vec2{0, 0},
}

var bootHouse = scean{
	mapXY:    vec2{67, 0},
	mapWH:    vec2{10, 10},
	cameraXY: vec2{0, 0},
}

var townScean = scean{
	mapXY:    vec2{0, 0},
	mapWH:    vec2{36, 36},
	cameraXY: vec2{0, 0},
}

// We have to do this to prevent an infinite type checking loop
func initSceans() {
	mainScean = &townScean
	linkSceans(&townScean, vec2{0, 0}, &dogHouse, vec2{0, 0})
	linkSceans(&townScean, vec2{0, 0}, &dogHouse, vec2{0, 0})
}

func linkSceans(a *scean, enterA vec2, b *scean, enterB vec2) {
	a.entrances[b] = enterA
	a.entrances[a] = enterB
}

func (s *scean) draw() {
	g.Map(int(s.mapXY.x), int(s.mapXY.y), int(s.mapWH.x), int(s.mapWH.y), 0, 0)
}
