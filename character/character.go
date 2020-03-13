package character

import "alterhero/world"

type Actor interface {
	act(field world.Field, equipment Equipment) (Action, world.Direction)
}

type Character struct {
	Life int
	Position [2]int
	Actor Actor
}
