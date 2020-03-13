package character

import "alterhero/world"

type Cost struct {
	Object world.Object
	Number int
}

type Weapon struct {
	Distance   int
	Strongness int
	Cost []Cost
}

var Fist = Weapon{
	Distance:   1,
	Strongness: 1,
}
var Bullet = Weapon{
	Distance:   2,
	Strongness: 3,
	Cost: []Cost{
		{
			Object: world.Rock,
			Number: 1,
		},
	},
}
var Stick = Weapon{
	Distance:   1,
	Strongness: 2,
	Cost: []Cost{
		{
			Object: world.Wood,
			Number: 1,
		},
	},
}
var Axe = Weapon{
	Distance:   1,
	Strongness: 3,
	Cost: []Cost{
		{
			Object: world.Rock,
			Number: 1,
		},
		{
			Object: world.Wood,
			Number: 1,
		},
	},
}
var Arch = Weapon{
	Distance:   3,
	Strongness: 3,
	Cost: []Cost{
		{
			Object: world.Wood,
			Number: 2,
		},
	},
}
var Sword = Weapon{
	Distance:   1,
	Strongness: 4,
	Cost: []Cost{
		{
			Object: world.Stone,
			Number: 3,
		},
		{
			Object: world.Fire,
			Number: 1,
		},
	},
}
