package character

type Action int

const (
	Move Action = iota
	Hit
	Pick
	Eat
	Forge
)
