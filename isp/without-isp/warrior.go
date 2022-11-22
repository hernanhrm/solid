package main

type Warrior interface {
	Kick()
	Punch()
	Transform()
}

type Warriors []Warrior

func executeWithoutIPS(warriors Warriors) {
	for _, warrior := range warriors {
		warrior.Kick()
		warrior.Punch()
		warrior.Transform()
	}
}
