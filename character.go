package main

type Character struct {
	name     string
	health   int
	strength int
	defense  int
	magic    int
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
// `Mage`: High `magic`, low `strength`

type Warrior struct {
	Character
}

func NewWarrior(name string) Warrior {
	return Warrior{
		Character: Character{
			name:     name,
			health:   120,
			strength: 100,
			defense:  100,
			magic:    10,
		},
	}
}

type Mage struct {
	Character
}

func NewMage(name string) Mage {
	return Mage{
		Character: Character{
			name:     name,
			health:   120,
			strength: 20,
			defense:  80,
			magic:    100,
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
			health:   90,
			strength: 90,
			defense:  70,
			magic:    10,
		},
	}
}
