package main

var physicsSteps = 9.0

func runPhysicsStep(e *entity) {
	dirX := float64(e.pos.x-e.colide.oldX) / physicsSteps
	dirY := float64(e.pos.y-e.colide.oldY) / physicsSteps
	Colide, xColide, yColide := false, false, false
	for i := 0; i < entityPointer; i++ {
		if i == e.id {
			continue
		}
		o := allEntities[i]
		if !o.hasComponent(colidable) {
			continue
		}
		x1 := float64(e.colide.oldX) + e.colide.deltaX + dirX
		y1 := float64(e.colide.oldY) + e.colide.deltaY + dirY
		w1, h1 := float64(e.colide.width), float64(e.colide.height)
		x2 := float64(o.colide.oldX) + o.colide.deltaX
		y2 := float64(o.colide.oldY) + o.colide.deltaY
		w2, h2 := float64(o.colide.width), float64(o.colide.height)
		if !Colide && didCollide(x1, y1, w1, h1, x2, y2, w2, h2) {
			Colide = true
		}
		if !xColide && didCollide(x1, y1-dirY, w1, h1, x2, y2, w2, h2) {
			xColide = true
		}
		if !yColide && didCollide(x1-dirX, y1, w1, h1, x2, y2, w2, h2) {
			yColide = true
		}
	}
	if !xColide {
		e.colide.deltaX += dirX
	}
	if !yColide {
		e.colide.deltaY += dirY
	}
}

func didCollide(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	xOver, yOver := false, false
	if x1 > x2 && x1 < x2+w2 {
		xOver = true
	}
	if y1 > y2 && y1 < y2+h2 {
		yOver = true
	}
	if x1+w1 > x2 && x1+w1 < x2+w2 {
		xOver = true
	}
	if y1+h1 > y2 && y1+h1 < y2+h2 {
		yOver = true
	}
	if x1 > x2 && x1+w1 < x2+w2 {
		xOver = true
	}
	if y1 > y2 && y1+h1 < y2+h2 {
		yOver = true
	}
	if x2 > x1 && x2+w2 < x1+w1 {
		xOver = true
	}
	if y2 > y1 && y2+h2 < y1+h1 {
		yOver = true
	}
	//We need these checks because the floats
	//are converted from integers and so might
	//actually be exactly equal
	if y1 == y2 || y1 == y2+h2 {
		yOver = true
	}
	if x1 == x2 || x1 == x2+w2 {
		xOver = true
	}
	return xOver && yOver
}
