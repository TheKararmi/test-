package model

type Torgoves struct {
	HP     int
	Damage int
	Spawn  *Spawn
}

func (t *Torgoves) MinusHP(damage int) {
	t.HP -= damage
}

func (t *Torgoves) GetDamage() int {
	return t.Damage
}

func (t *Torgoves) GetSpawn() (int, int) {
	return t.Spawn.X, t.Spawn.Y
}
