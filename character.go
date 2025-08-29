package main

import "fmt"

type Character struct {
	name           string
	health         int
	strength       int
	defense        int
	magic          int
	equippedWeapon string
	equippedArmor  string
}

//creating characters

func NewCharacter(name string) Character {
	newCharacter := Character{
		name:     name,
		health:   0,
		strength: 0,
		defense:  0,
		magic:    0,
	}

	return newCharacter
}

// Warrior`: High `strength` and `defense`, low `magic`
type Warrior struct {
	Character
}

func NewWarrior(name string) Warrior {
	return Warrior{
		Character: Character{
			name:     name,
			health:   10,
			strength: 15,
			defense:  10,
			magic:    5,
		},
	}
}

type Mage struct {
	Character
}

// `Mage`: High `magic`, low `strength`
func NewMage(name string) Mage {
	return Mage{
		Character: Character{
			name:     name,
			health:   10,
			strength: 5,
			defense:  10,
			magic:    15,
		},
	}
}

type Archer struct {
	Character
}

// `Archer`: Moderate `strength` and `health`, low `magic`
func NewArcher(name string) Archer {
	return Archer{
		Character: Character{
			name:     name,
			health:   10,
			strength: 10,
			defense:  7,
			magic:    5,
		},
	}
}

type CharacterInterface interface {
	EquipWeapon(weapon item)
	EquipArmor(armor item)
}

func (warrior Warrior) EquipWeapon(weapon item) {
	if weapon.name == "Magic Staff" {
		fmt.Println("Warriors cannot equip magic staves.")
	} else {
		warrior.equippedWeapon = weapon.name
	}
}

func (warrior Warrior) EquipArmor(armor item) {
	if armor.name == "Magic Robe" {
		fmt.Println("Warriors cannot equip magic robes.")
	} else {
		warrior.equippedArmor = armor.name
	}
}
