package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	mod := float64(10-score) / 2
	
	return int(math.Floor( -mod))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rand.Intn(16)+3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	new:= Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		}
		new.Hitpoints =10+ Modifier(new.Constitution)
		return new
}
