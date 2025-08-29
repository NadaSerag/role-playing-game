package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, readerPtr *bufio.Reader) string {
	fmt.Println(prompt)
	input, _ := readerPtr.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

// to run type in terminal (powershell): go run calc.go
func main() {

	reader := bufio.NewReader(os.Stdin)

	characterChoice := getInput("Choose the character you want to create: (a) Warrior, (b) Mage, (c) Archer", reader)
	if characterChoice != "a" && characterChoice != "b" && characterChoice != "c" {
		fmt.Println("Invalid choice")
		return
	}
	characterName := getInput("Name your character: ", reader)

	var player CharacterInterface

	switch characterChoice {
	case "a":
		{
			player = NewWarrior(characterName)
		}
	case "b":
		{
			player = NewMage(characterName)
		}
	case "c":
		{
			player = NewArcher(characterName)
		}
	}
	// **Weapons**:
	//     - "Heavy Sword" (Damage Bonus: 5)
	//     - "Magic Staff" (Damage Bonus: 3)
	//     - "Bow" (Damage Bonus: 4)
	//   - **Armor**:
	//     - "Plate Armor" (Defense Bonus: 5)
	//     - "Leather Armor" (Defense Bonus: 3)
	//     - "Magic Robe" (Defense Bonus: 2)
	// - The user chooses "Heavy Sword" and "Plate Armor."

	weaponChoice := getInput("Choose a weapon: (a) Heavy Sword (Damage Bonus: 5), (b) Magic Staff (Damage Bonus: 3), (c) Bow (Damage Bonus: 4)", reader)

	var chosenWeapon Weapon
	switch weaponChoice {
	case "a":
		{
			chosenWeapon = NewWeapon("Heavy Sword")
		}
	case "b":
		{
			chosenWeapon = NewWeapon("Magic Staff")
		}
	case "c":
		{
			chosenWeapon = NewWeapon("Bow")
		}
	}

	player.EquipWeapon(chosenWeapon)
	fmt.Println(chosenWeapon.name, " equipped!")

	armorChoice := getInput("Choose an armor: (a) Plate Armor (Damage Bonus: 5), (b) Leather Armor (Damage Bonus: 3), (c) Magic Robe (Damage Bonus: 4)", reader)

	var chosenArmor Armor

	switch armorChoice {
	case "a":
		{
			chosenArmor = NewArmor("Heavy Sword")
		}
	case "b":
		{
			chosenArmor = NewArmor("Magic Staff")
		}
	case "c":
		{
			chosenArmor = NewArmor("Bow")
		}
	}
	player.EquipArmor(chosenArmor)
	fmt.Println(chosenArmor.name, " equipped!")

}
