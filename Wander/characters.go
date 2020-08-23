package main

import "github.com/bjatkin/golf-engine/golf"

type collidable interface {
	collide(vec2) bool
}

type character struct {
	*interaction
	n   int
	pos vec2
	o   golf.SOp
}

func (c *character) collide(player vec2) bool {
	w, h := float64(c.o.W*8), float64(c.o.H*8)

	// player is on the left or right
	if player.x+14 <= c.pos.x || player.x >= c.pos.x+w {
		return false
	}

	// player is above or below
	if player.y+16 <= c.pos.y || player.y >= c.pos.y+h {
		return false
	}

	return true
}

func (c *character) drawInteractable(player vec2) {
	c.location = c.pos
	c.interaction.drawInteractable(player)
	g.Spr(c.n, c.pos.x, c.pos.y, c.o)
}

var guard1 = character{
	interaction: &guardInter,
	n:           22,
	pos:         vec2{192, 30},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2, FH: true},
}

var guard2 = character{
	interaction: &guardInter,
	n:           22,
	pos:         vec2{208, 30},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2},
}

var fishKid = character{
	interaction: &fishKidInter,
	n:           18,
	pos:         vec2{248, 182},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2},
}

var wellMan = character{
	interaction: &wellManInter,
	n:           30,
	pos:         vec2{352, 30},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2},
}

var talkingDog = character{
	interaction: &talkingDogInter,
	n:           20,
	pos:         vec2{38, 44},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2},
}

var countingKid = character{
	interaction: &countingKidInter,
	n:           26,
	pos:         vec2{152, 144},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2},
}

var oldMan = character{
	interaction: &oldManInter,
	n:           28,
	pos:         vec2{74, 40},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2},
}

var bootLady = character{
	interaction: &bootLadyInter,
	n:           24,
	pos:         vec2{24, 22},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2},
}

var rentGuy = character{
	interaction: &rentGuyInter,
	n:           30, // TODO swap this for the new graphics when they get created
	pos:         vec2{24, 22},
	o:           golf.SOp{TCol: golf.Col6, W: 2, H: 2},
}

func initCharacterEvents() {
	storyEventHandler.onEvent(talkedToDog, func() { rentGuy.lines = []int{19, 36, 20} })
	storyEventHandler.onEvent(twoFiftySixSecondsPassed, func() { countingKid.lines = []int{26, 27, 28, 29} })
	storyEventHandler.onEvent(talkToCountKidAfterFinished, func() { countingKid.lines = []int{30} })
	storyEventHandler.onEvent(talkedToOldGuy, func() { oldMan.lines = []int{6, 10} })
}
