package arena_test

import (
	"MagicalArena/arena"
	"testing"
)

// TestGame simulates a game and checks that it ends correctly.
func TestGame(t *testing.T) {
	playerA := arena.NewPlayer("Player A", 50, 5, 10)
	playerB := arena.NewPlayer("Player B", 100, 10, 5)
	game := arena.NewGame(playerA, playerB)
	game.Start()

	if !(playerA.Health == 0 || playerB.Health == 0) {
		t.Errorf("Game did not end correctly. Player A health: %d, Player B health: %d", playerA.Health, playerB.Health)
	}
}

// TestPlayerAttributes checks that the player attributes are set correctly.
func TestPlayerAttributes(t *testing.T) {
	player := arena.NewPlayer("Player", 50, 5, 10)
	if player.Name != "Player" {
		t.Errorf("Expected Player, got %s", player.Name)
	}
	if player.Health != 50 {
		t.Errorf("Expected 50, got %d", player.Health)
	}
	if player.Strength != 5 {
		t.Errorf("Expected 5, got %d", player.Strength)
	}
	if player.Attack != 10 {
		t.Errorf("Expected 10, got %d", player.Attack)
	}
}

// TestDetermineFirstAttacker checks that the first attacker is determined correctly based on health.
func TestDetermineFirstAttacker(t *testing.T) {
	playerA := arena.NewPlayer("Player A", 50, 5, 10)
	playerB := arena.NewPlayer("Player B", 100, 10, 5)
	game := arena.NewGame(playerA, playerB)
	firstAttacker, _ := game.DetermineFirstAttacker()
	if firstAttacker != playerA {
		t.Errorf("Expected Player A to be the first attacker, got %s", firstAttacker.Name)
	}
}

// TestPlayerIsAlive checks that player's alive status is correctly determined.
func TestPlayerIsAlive(t *testing.T) {
	player := arena.NewPlayer("Player", 50, 5, 10)
	if !player.IsAlive() {
		t.Errorf("Expected Player to be alive, got dead")
	}
	player.ReduceHealth(50)
	if player.IsAlive() {
		t.Errorf("Expected Player to be dead, got alive")
	}
}
