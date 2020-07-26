package main

var allEntities = [100]entity{} // is 100 to many? not enough?
var entityPointer = 0

type entity struct {
	id    int
	flags flag
	sig   compType
}

func (e *entity) hasFlag(f flag) bool {
	return (e.flags & f) == f
}

func (e *entity) hasComponent(c compType) bool {
	return (e.sig & c) == c
}

func newEntity(flags flag, comps ...component) *entity {
	ret := &entity{flags: flags}
	ret.id = entityPointer
	var sig compType
	for _, comp := range comps {
		comp.setEnt(ret)
		comp.add()
		sig |= comp.ctype()
	}
	ret.sig = sig
	allEntities[entityPointer] = *ret
	entityPointer++
	return ret
}
