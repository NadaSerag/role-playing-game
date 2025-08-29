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

	switch characterChoice {
	case "a":
		{
			NewWarrior(characterName)
		}
	case "b":
		{
			NewMage(characterName)
		}
	case "c":
		{
			NewArcher(characterName)
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
	

}
