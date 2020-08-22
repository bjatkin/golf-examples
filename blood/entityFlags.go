package main

type flag int

const (
	none             = flag(0)
	playerControlled = flag(1)
	enemy            = flag(2)
	projectile       = flag(4)
)
