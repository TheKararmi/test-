package model

type Nos struct {
	HP     int
	Damage int
	Spawn  *Spawn
}

func (n *Nos) MinusHP(damage int) {
	n.HP -= damage
}

func (n *Nos) GetDamage() int {
	return n.Damage
}

func (n *Nos) GetSpawn() (int, int) {
	return n.Spawn.X, n.Spawn.Y
}
