package world

type Path []Direction

type Proximity interface {
	containsWoods() (bool, Path)
	containsTrees() (bool, Path)
	containsStones() (bool, Path)
	containsRocks() (bool, Path)
	containsAnimals() (bool, Path)
	containsDragons() (bool, Path)
}

type Field struct {
	Area [][]rune
	Proximity Proximity
}
