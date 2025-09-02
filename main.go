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
			warrior := NewWarrior(characterName)
			player = &warrior

			//why not player = &NewWarrior(characterName) ??
			//&NewWarrior(...) tries to take the address of a function call result, which Go does not allow directly.
			//Go requires that & be applied to a variable or composite literal, not the return value of a function.

			//modifying the string characterChoice to be the character type to display it later in PrintStats
			characterChoice = "Warrior"
		}
	case "b":
		{
			mage := NewMage(characterName)
			player = &mage
			characterChoice = "Mage"
		}
	case "c":
		{
			archer := NewArcher(characterName)
			player = &archer
			characterChoice = "Archer"
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
	armorChoice := getInput("Choose an armor: (a) Plate Armor (Damage Bonus: 5), (b) Leather Armor (Damage Bonus: 3), (c) Magic Robe (Damage Bonus: 4)", reader)
	var chosenArmor Armor
	switch armorChoice {
	case "a":
		{
			chosenArmor = NewArmor("Plate Armor")
		}
	case "b":
		{
			chosenArmor = NewArmor("Leather Armor")
		}
	case "c":
		{
			chosenArmor = NewArmor("Magic Robe")
		}
	}
	player.EquipArmor(chosenArmor)


	player.PrintStats(characterChoice)

}