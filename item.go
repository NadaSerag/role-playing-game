package main

type item struct {
	name string
}

type Weapon struct {
	item        // using composition here
	damageBonus int
}

type Armor struct {
	item
	defenseBonus int
}

func NewWeapon(name string) Weapon {
	var def int
	switch name {
	case "Heavy Sword":
		{
			def = 5
		}
	case "Magic Staff":
		{
			def = 3
		}
	case "Bow":
		{
			def = 4
		}
	}

	newWeapon := Weapon{
		item: item{
			name: name,
		},
		damageBonus: def,
	}

	return newWeapon
}
