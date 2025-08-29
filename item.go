package main

type item struct {
	name string
}

type Weapon struct {
	item
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

func NewArmor(name string) Armor {
	var def int
	switch name {
	case "Plate Armor":
		{
			def = 5
		}
	case "Leather Armor":
		{
			def = 3
		}
	case "Magic Robe":
		{
			def = 4
		}
	}

	newArmor := Armor{
		item: item{
			name: name,
		},
		defenseBonus: def,
	}

	return newArmor
}
