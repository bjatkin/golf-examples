package main

import (
	"fantasyConsole/golf"
	"math"
)

type interactable interface {
	drawInteractable(vec2)
	resetInteraction()
	interact(vec2) bool
	next() bool
	line() string
}

type interaction struct {
	location    vec2
	lines       []int
	currentLine int
	cache       string
}

var mainInteraction interactable

func (i *interaction) drawInteractable(player vec2) {
	dx := math.Abs(i.location.x - player.x)
	dy := math.Abs(i.location.y - player.y)
	if dx > 30 || dy > 30 {
		return
	}

	o := golf.SOp{TCol: golf.Col6, W: 2, H: 2}
	x := i.location.x
	y := i.location.y
	switch (g.Frames() / 20) % 4 {
	case 0:
		g.Spr(131, x, y-22, o)
	case 1:
		g.Spr(133, x, y-23, o)
	case 2:
		g.Spr(135, x, y-22, o)
	case 3:
		g.Spr(133, x, y-23, o)
	}
	return
}

func (i *interaction) interact(iPoint vec2) bool {
	dx := math.Abs(i.location.x - iPoint.x)
	dy := math.Abs(i.location.y - iPoint.y)
	if dx < 16 && dy < 16 {
		return true
	}
	return false
}

func (i *interaction) resetInteraction() {
	i.cache = ""
	i.currentLine = 0
}

func (i *interaction) next() bool {
	i.cache = ""
	i.currentLine++
	return i.currentLine < len(i.lines)
}

func (i *interaction) line() string {
	if i.cache == "" {
		line := allLines[i.lines[i.currentLine]]
		i.cache = fillTemplate(line)
	}
	return i.cache
}

var guardInter = interaction{
	lines: []int{0, 1, 2, 3, 4, 5},
}

var oldManInter = interaction{
	lines: []int{6, 7, 33, 8, 34, 9},
}

var bootLadyInter = interaction{
	lines: []int{11},
}

var wellManInter = interaction{
	lines: []int{12, 37, 13},
}

var talkingWellInter = interaction{
	location: vec2{337, 60},
	lines:    []int{14, 15},
}

var fishKidInter = interaction{
	lines: []int{16, 17},
}

var rentGuyInter = interaction{
	lines: []int{18},
}

var talkingDogInter = interaction{
	lines: []int{21, 22, 23, 35},
}

var countingKidInter = interaction{
	lines: []int{24, 25},
}

var townSignInter = interaction{
	location: vec2{129, 190},
	lines:    []int{31, 32},
}

var storeSignInter = interaction{
	location: vec2{104, 190},
	lines:    []int{38, 39, 40, 41},
}

func updateInteraction() {
	if g.Btnp(golf.ZKey) {
		if mainInteraction.next() {
			return
		}

		g.Update = update
		g.Draw = draw

		if mainInteraction == &oldMan {
			storyEventHandler.sendEvent(talkedToOldGuy)
		}
		if mainInteraction == &talkingDog {
			storyEventHandler.sendEvent(talkedToDog)
		}
		if mainInteraction == &wellMan {
			storyEventHandler.sendEvent(talkedToWellGuy)
		}
		if mainInteraction == &countingKid && g.Frames()/30 >= 253 {
			storyEventHandler.sendEvent(talkToCountKidAfterFinished)
		}
	}
}

func drawInteraction() {
	g.RectFill(0, 150, 192, 42, golf.Col0, true)
	g.RectFill(1, 151, 190, 40, golf.Col7, true)
	g.Rect(2, 152, 188, 38, golf.Col0, true)
	g.Text(4, 158, mainInteraction.line(), golf.TOp{Fixed: true})
}
