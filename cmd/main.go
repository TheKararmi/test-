package main

import (
	"fmt"
	"minecraft/pkg/model"
)

type NPC interface {
	MinusHP(damage int)
	GetDamage() int
	GetSpawn() (x int, y int)
}
type inventory struct {
	GoldenApple  int
	Pixare       int
	BlockOfGrass int
}

func main() {
	z := &model.Zombie{
		HP:     80,
		Damage: 10,
		Spawn:  model.SpawnGenerage(),
	}

	c := &model.Criper{
		HP:     100,
		Damage: 50,
		Spawn:  model.SpawnGenerage(),
	}

	n := &model.Nos{
		HP:     50,
		Damage: 0,
		Spawn:  model.SpawnGenerage(),
	}

	t := &model.Torgoves{
		HP:     50,
		Damage: 0,
		Spawn:  model.SpawnGenerage(),
	}

	npcs := []NPC{z, c, n, t}

	Boom(50, npcs)

	for _, v := range npcs {
		fmt.Println(v)
	}
}

func Boom(boomDamage int, npcs []NPC) {
	for _, v := range npcs {
		v.MinusHP(boomDamage)
	}
}
