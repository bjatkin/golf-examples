package main

import (
	"strconv"
)

var allCollidables = [len(allEntities)]int{}
var collidablePointer = 0

func doPhysics(e *entity) {
	collide := collideComponents[e.id]
	trans := transformComponents[e.id]

	collide.deltaX = trans.x - collide.oldX
	collide.deltaY = trans.y - collide.oldY
	for i := 0; i < collidablePointer; i++ {
		cid := allCollidables[i]
		if cid == e.id {
			continue
		}
		check := collideComponents[cid]
		x1, y1 := check.oldX+check.deltaX, check.oldY+check.deltaY
		w1, h1 := check.width, check.height
		x2, y2 := collide.oldX+collide.deltaX, collide.oldY+collide.deltaY
		w2, h2 := collide.width, collide.height
		if didCollide(x1, y1, w1, h1, x2, y2, w2, h2) {
			debug = "collide " + strconv.Itoa(i) + "/" + strconv.Itoa(collidablePointer-1)
		}
	}
}

func didCollide(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	var xOver, yOver bool
	if x1 >= x2 && x1 <= x2+h2 {
		xOver = true
	}
	if x1+h1 >= x2 && x1+h1 <= x2+h2 {
		xOver = true
	}
	if y1 >= y2 && y1 <= y2+h2 {
		yOver = true
	}
	if y1+h1 >= y2 && y1+h1 <= y2+h2 {
		yOver = true
	}
	return xOver && yOver
}
