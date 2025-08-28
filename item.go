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
