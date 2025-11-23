package model

type Zombie struct {
	HP     int
	Damage int
	Spawn  *Spawn
}

func (z *Zombie) MinusHP(damage int) {
	z.HP -= damage
}

func (z *Zombie) GetDamage() int {
	return z.Damage
}

func (z *Zombie) GetSpawn() (int, int) {
	return z.Spawn.X, z.Spawn.Y
}
