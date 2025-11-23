package model

type Criper struct {
	HP     int
	Damage int
	Spawn  *Spawn
}

// func (c *Criper) MinusHP(damage int) {
// 	c.HP -= damage
// }

// func (c *Criper) GetDamage() int {
// 	return c.Damage
// }

// func (c *Criper) GetSpawn() (int, int) {
// 	return c.Spawn.X, c.Spawn.Y
// }
