package main

type Warrior interface {
	Kick()
	Punch()
}

type SuperSaiyan interface {
	Transform()
}

type Warriors []Warrior

func executeWithIPS(warriors Warriors) {
	for _, warrior := range warriors {
		warrior.Kick()
		warrior.Punch()

		// For each warrior, we check if it is a SuperSaiyan
		if superSaiyan, ok := warrior.(SuperSaiyan); ok {
			superSaiyan.Transform()
		}
	}
}
