package main

import "math"

type vec2 struct {
	x, y float64
}
type scean struct {
	mapXY     vec2
	mapWH     vec2
	start     []vec2
	exit      []vec2
	nextScean []*scean
	nextStart []int
}

var rentHouse = scean{
	mapXY:     vec2{52, 0},
	mapWH:     vec2{9, 9},
	start:     []vec2{vec2{20, 20}},
	exit:      []vec2{vec2{25, 25}},
	nextScean: []*scean{&mainScean},
	nextStart: []int{0},
}

var dogHouse = scean{
	mapXY:     vec2{57, 0},
	mapWH:     vec2{9, 9},
	start:     []vec2{vec2{20, 20}},
	exit:      []vec2{vec2{25, 25}},
	nextScean: []*scean{&mainScean},
	nextStart: []int{1},
}

var bootHouse = scean{
	mapXY:     vec2{67, 0},
	mapWH:     vec2{10, 10},
	start:     []vec2{vec2{20, 20}},
	exit:      []vec2{vec2{25, 25}},
	nextScean: []*scean{&mainScean},
	nextStart: []int{2},
}

var mainScean = scean{
	mapXY: vec2{0, 0},
	mapWH: vec2{36, 36},
	start: []vec2{vec2{50, 50}},
	exit: []vec2{
		vec2{10, 10},  // Rent house
		vec2{30, 10},  // Dog house
		vec2{100, 50}, // Boot house
	},
}

// We have to do this to prevent an infinite type checking loop
func initSceans() {
	mainScean.nextScean = []*scean{&rentHouse, &dogHouse, &bootHouse}
	mainScean.nextStart = []int{0, 0, 0}
}

func (s *scean) draw() {
	g.Map(int(s.mapXY.x), int(s.mapXY.y), int(s.mapWH.x), int(s.mapWH.y), 0, 0)
}

func (s *scean) leave(x, y float64) (int, bool) {
	for i, e := range s.exit {
		dx := math.Abs(e.x - x)
		dy := math.Abs(e.y - y)
		if dx < 3 && dy < 3 {
			return i, true
		}
	}
	return 0, false
}
