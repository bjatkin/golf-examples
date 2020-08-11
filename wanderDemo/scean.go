package main

import "math"

type vec2 struct {
	x, y float64
}

var mainScean *scean

func exitScean(player vec2) (*scean, bool) {
	for s, pt := range mainScean.entrances {
		dx := math.Abs(player.x - pt[0].x)
		dy := math.Abs(player.y - pt[0].y)
		if dx < 8 && dy < 4 {
			return s, true
		}
	}
	return nil, false
}

type scean struct {
	mapXY     vec2
	mapWH     vec2
	entrances map[*scean][2]vec2
	poi       []*interaction
}

var rentHouse = scean{
	mapXY: vec2{52, 0},
	mapWH: vec2{5, 7},
}

var dogHouse = scean{
	mapXY: vec2{57, 0},
	mapWH: vec2{10, 13},
}

var bootHouse = scean{
	mapXY: vec2{67, 0},
	mapWH: vec2{11, 9},
}

var townScean = scean{
	mapXY: vec2{0, 0},
	mapWH: vec2{52, 37},
}

// We have to do this to prevent an infinite type checking loop
func initSceans() {
	// Init all the maps
	townScean.entrances = make(map[*scean][2]vec2)
	dogHouse.entrances = make(map[*scean][2]vec2)
	bootHouse.entrances = make(map[*scean][2]vec2)
	rentHouse.entrances = make(map[*scean][2]vec2)

	mainScean = &townScean
	linkSceans(&townScean, vec2{136, 96}, vec2{136, 102}, &dogHouse, vec2{16, 88}, vec2{16, 78})
	linkSceans(&townScean, vec2{296, 126}, vec2{295, 132}, &bootHouse, vec2{0, 55}, vec2{0, 44})
	linkSceans(&townScean, vec2{104, 94}, vec2{104, 102}, &rentHouse, vec2{0, 42}, vec2{12, 42})
}

func linkSceans(a *scean, exitA, enterA vec2, b *scean, exitB, enterB vec2) {
	a.entrances[b] = [2]vec2{exitA, enterA}
	b.entrances[a] = [2]vec2{exitB, enterB}
}

func (s *scean) draw() {
	g.Map(int(s.mapXY.x), int(s.mapXY.y), int(s.mapWH.x), int(s.mapWH.y), 0, 0)
}
