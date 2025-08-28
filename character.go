package main

type character struct {
	name     string
	health   int
	strength int
	defense  int
	magic    int
}

//creating characters

func NewCharacter(name string) character {
	newCharacter := character{
		name:     name,
		health:   0,
		strength: 0,
		defense:  0,
		magic:    0,
	}

	return newCharacter
}
