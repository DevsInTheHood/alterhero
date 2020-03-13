package character

type Inventory interface {
	storeWood() error
	currentWoods() int
	storeStone() error
	currentStones() int
	changeWeapon(weapon Weapon)
	currentWeapon() Weapon
}

type Equipment struct {
	Woods  int
	Stones int
	Weapon Weapon
}
