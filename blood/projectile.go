package main

func initProjectileSystem() {
	allUpdateSystems[doProjectile] = toSystem(
		projectile,
		TypeTransformComponent|TypeTravelComponent,
		func(e *entity) {
			t := transformComponents[e.id]
			delta := travelComponents[e.id]
			t.x += delta.dx
			t.y += delta.dy
			if int(t.x+8) < cameraX || int(t.y+8) < cameraY ||
				int(t.x) > cameraX+192 || int(t.y) > cameraY+192 {
				deleteEntity(e)
			}
		})
}

func newProjectile(x, y, dx, dy float64, big bool) *entity {
	return newEntity(projectile,
		&sprComponent{n: 82},
		&transformComponent{x: x, y: y},
		&travelComponent{dx: dx, dy: dy},
	)
}
