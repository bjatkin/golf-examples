package main

import "github.com/bjatkin/golf-engine/golf"

// convo is an interaction between the player and another character
type convo struct {
	portrait   []int
	lines      []string
	line       int
	height     int
	goalHeight int
	running    bool
}

// draw the conversation box and the text
func (c *convo) draw() {
	g.RectFill(0, float64(c.height), 192, 34, golf.Col7, true)
	g.Rect(0, float64(c.height), 192, 34, golf.Col0, true)
	g.Spr(c.portrait[c.line], 1, float64(c.height)+2, golf.SOp{W: 4, H: 4, TCol: golf.Col5, Fixed: true})
	if c.height < c.goalHeight {
		c.height++
	}
	if c.height > c.goalHeight {
		c.height--
	}
	g.Text(35, float64(c.height+4), c.lines[c.line], golf.TOp{Fixed: true})
}

// get the next line of the conversation
func (c *convo) next() {
	if !c.running {
		c.running = true
		c.goalHeight = 158
		c.line = 0
	}
	if c.height != c.goalHeight {
		c.height = c.goalHeight
		return
	}
	c.line++
	if c.line >= len(c.lines) {
		c.line--
		c.goalHeight = 192
		c.running = false
	}
}

// create the conversation with joe gopher
func initConvo() {
	gopherConvo = &convo{
		portrait: []int{65, 69, 65},
		lines: []string{
			"BIBI DUCK: Hey, do you\nknow how to get out of\nthis place!?",
			"JOE GOPHER: Why would you\nwant to leave?\nThis place is great!",
			"BIBI DUCK: Oh...",
		},
		height:     192,
		goalHeight: 192,
	}
}
