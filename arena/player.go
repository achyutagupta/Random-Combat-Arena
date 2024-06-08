package arena

import (
	"fmt"
	"log"
)

// Player struct holds the attributes of a player.
type Player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}

// CreatePlayer prompts the user to input valid attributes for a player.
func CreatePlayer(playerName string) *Player {
	var health, strength, attack int

	for {
		fmt.Printf("Enter attributes for %s:\n", playerName)
		fmt.Print("Health (positive integer): ")
		if _, err := fmt.Scan(&health); err != nil || health <= 0 {
			log.Println("Invalid input. Please enter a positive integer for health.")
			continue
		}
		fmt.Print("Strength (positive integer): ")
		if _, err := fmt.Scan(&strength); err != nil || strength <= 0 {
			log.Println("Invalid input. Please enter a positive integer for strength.")
			continue
		}
		fmt.Print("Attack (positive integer): ")
		if _, err := fmt.Scan(&attack); err != nil || attack <= 0 {
			log.Println("Invalid input. Please enter a positive integer for attack.")
			continue
		}
		break
	}

	return NewPlayer(playerName, health, strength, attack)
}

// NewPlayer creates a new player with specified attributes.
func NewPlayer(name string, health, strength, attack int) *Player {
	return &Player{Name: name, Health: health, Strength: strength, Attack: attack}
}

// ReduceHealth reduces the player's health by the specified damage.
func (p *Player) ReduceHealth(damage int) {
	p.Health = max(0, p.Health-damage)
}

// IsAlive checks if the player is still alive.
func (p *Player) IsAlive() bool {
	return p.Health > 0
}

// Returns max of two integer values.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
