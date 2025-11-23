package model

import "math/rand"

type Spawn struct {
	X int
	Y int
}

func SpawnGenerage() *Spawn {
	return &Spawn{
		X: rand.Int(),
		Y: rand.Int(),
	}
}
