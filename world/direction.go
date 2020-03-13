package world

type Direction [2]int

var West = Direction{-1, 0}
var NorthWest = Direction{-1, 1}
var North = Direction{0, 1}
var NortEast = Direction{1, 1}
var East = Direction{1, 0}
var SouthEast = Direction{1, -1}
var South = Direction{0, -1}
var SouthWest = Direction{-1, -1}
