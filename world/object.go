package world

type Object rune

const (
	/* STATIC OBJECTS */
	EmptySpace  Object = ' '
	Wood        Object = '|'
	DoubleWood  Object = '!'
	Tree        Object = '@'
	Stone       Object = 'o'
	DoubleStone Object = 'ø'
	Rock        Object = 'O'
	Water       Object = '˜'
	Fire        Object = 'M'
	/* MOVING OBJECTS */
	Hero   Object = '&'
	Animal Object = 'x'
	Dragon Object = '§'
)
