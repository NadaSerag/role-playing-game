package main

import (
	"fmt"
	"strconv"
)

type Character struct {
	name           string
	health         int
	strength       int
	defense        int
	magic          int
	equippedWeapon Weapon
	equippedArmor  Armor
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
	AddWeaponPoints(weapon Weapon)
	AddArmorPoints(archer Armor)
	EquipWeapon(weapon Weapon)
	EquipArmor(armor Armor)
	PrintStats(characterPersonality string)
}

func (character *Character) AddWeaponPoints(weapon Weapon) {
	switch weapon.name {
	case "Heavy Sword":
		{
			character.strength += 5
		}
	case "Magic Staff":
		{
			character.strength += 3
		}
	case "Bow":
		{
			character.strength += 4
		}
	}
}

func (character *Character) AddArmorPoints(armor Armor) {
	switch armor.name {
	case "Plate Armor":
		{
			character.defense += 5
		}
	case "Leather Armor":
		{
			character.defense += 3
		}
	case "Magic Robe":
		{
			character.defense += 2
		}
	}
}

func (warrior *Warrior) EquipWeapon(weapon Weapon) {
	if weapon.name == "Magic Staff" {
		fmt.Println("Warriors cannot equip magic staves.")
	} else {
		warrior.equippedWeapon = weapon
		warrior.AddWeaponPoints(weapon)
		fmt.Println(weapon.name, " equipped!")
	}
}

func (warrior *Warrior) EquipArmor(armor Armor) {
	if armor.name == "Magic Robe" {
		fmt.Println("Warriors cannot equip magic robes.")
	} else {
		warrior.equippedArmor = armor
		warrior.AddArmorPoints(armor)
		fmt.Println(armor.name, " equipped!")
	}
}

func (mage *Mage) EquipWeapon(weapon Weapon) {
	if weapon.name == "Heavy Sword" {
		fmt.Println("Warriors cannot equip heavy armors.")
	} else {
		mage.equippedWeapon = weapon
		mage.AddWeaponPoints(weapon)
		fmt.Println(weapon.name, " equipped!")
	}
}

func (mage *Mage) EquipArmor(armor Armor) {
	mage.equippedArmor = armor
	mage.AddArmorPoints(armor)
	fmt.Println(armor.name, " equipped!")
}

func (archer *Archer) EquipWeapon(weapon Weapon) {
	if weapon.name == "Heavy Sword" {
		fmt.Println("Warriors cannot equip heay shields.")
	} else {
		archer.equippedWeapon = weapon
		archer.AddWeaponPoints(weapon)
		fmt.Println(weapon.name, " equipped!")
	}
}

func (archer *Archer) EquipArmor(armor Armor) {
	archer.equippedArmor = armor
	archer.AddArmorPoints(armor)
	fmt.Println(armor.name, " equipped!")
}

// function to print stats, common for all dervied structs of Character
func (character Character) PrintStats(characterPersonality string) {
	fmt.Println("Character: ", character.name, " (", characterPersonality, ") ")
	fmt.Println("Health: ", character.health)
	fmt.Println("Strength: ", character.strength)
	fmt.Println("Defense: ", character.defense)
	fmt.Println("Magic: ", character.magic)

	damagePointsStr := strconv.Itoa(character.equippedWeapon.GetPoints())
	defensePointsStr := strconv.Itoa(character.equippedArmor.GetPoints())

	fmt.Println("Equipped Weapon: ", character.equippedWeapon.GetName(), "(+", damagePointsStr, "damage)")
	fmt.Println("Equipped Armor: ", character.equippedArmor.GetName(), "(+", defensePointsStr, "defense)")

}
