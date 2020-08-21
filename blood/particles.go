package main

type blood struct {
	x, y      float64
	height    float64
	collected bool
	dead      bool
}

// Should each particle be an entity or should the
// particle system itself be the entity?
// I think the latter...

func addBloodParticle(x, y, height float64) {
	// should add a particle to the list of particles
	// this will likely get called whenver something (enemy/ player)
	// dies
}

func initParticleSystem() {
	// TODO add the particle update system
	// TODO add the particle draw system
}
