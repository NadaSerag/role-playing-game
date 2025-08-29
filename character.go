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
	EquipWeapon(weapon Weapon)
	EquipArmor(armor Armor)
}

func (warrior Warrior) EquipWeapon(weapon Weapon) {
	if weapon.name == "Magic Staff" {
		fmt.Println("Warriors cannot equip magic staves.")
	} else {
		warrior.equippedWeapon = weapon.name
	}
}

func (warrior Warrior) EquipArmor(armor Armor) {
	if armor.name == "Magic Robe" {
		fmt.Println("Warriors cannot equip magic robes.")
	} else {
		warrior.equippedArmor = armor.name
	}
}

func (mage Mage) EquipWeapon(weapon Weapon) {
	if weapon.name == "Heavy Sword" {
		fmt.Println("Warriors cannot equip heavy armors.")
	} else {
		mage.equippedWeapon = weapon.name
	}
}

func (mage Mage) EquipArmor(armor Armor) {
	mage.equippedArmor = armor.name
}

func (archer Archer) EquipWeapon(weapon Weapon) {
	if weapon.name == "Heavy Sword" {
		fmt.Println("Warriors cannot equip heay shields.")
	} else {
		archer.equippedWeapon = weapon.name
	}
}

func (archer Archer) EquipArmor(armor Armor) {
	archer.equippedArmor = armor.name
}
