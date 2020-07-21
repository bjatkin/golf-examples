package main

import (
	"strconv"
)

var allCollidables = [len(allEntities)]*entity{}
var collidablePointer = 0

func doPhysics(e *entity) {
	e.collide.deltaX = e.pos.x - e.collide.oldX
	e.collide.deltaY = e.pos.y - e.collide.oldY
	for i := 0; i < collidablePointer; i++ {
		check := allCollidables[i]
		if check == e {
			continue
		}
		x1, y1 := check.collide.oldX+check.collide.deltaX, check.collide.oldY+check.collide.deltaY
		w1, h1 := check.collide.width, check.collide.height
		x2, y2 := e.collide.oldX+e.collide.deltaX, e.collide.oldY+e.collide.deltaY
		w2, h2 := e.collide.width, e.collide.height
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
