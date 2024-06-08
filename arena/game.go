package arena

import "fmt"

// Game struct hold players and die used in the game.
type Game struct {
	PlayerA *Player
	PlayerB *Player
	Die     *Die
}

// Initialises a New Game.
func NewGame(playerA, playerB *Player) *Game {
	return &Game{
		PlayerA: playerA,
		PlayerB: playerB,
		Die:     NewDie(6),
	}
}

// Start begins the game and continues until one player reaches zero health.
func (g *Game) Start() {
	fmt.Printf("Starting the game between %s and %s\n", g.PlayerA.Name, g.PlayerB.Name)
	for g.PlayerA.IsAlive() && g.PlayerB.IsAlive() {
		g.PlayTurn()
	}
	if g.PlayerA.IsAlive() {
		fmt.Printf("Game over! Winner: %s\n", g.PlayerA.Name)
	} else {
		fmt.Printf("Game over! Winner: %s\n", g.PlayerB.Name)
	}
}


// PlayTurn executes a single turn in the game where each player attacks once.
func (g *Game) PlayTurn() {
	attacker, defender := g.determineFirstAttacker()
	g.attack(attacker, defender)
	if defender.IsAlive() {
		g.attack(defender, attacker)
	}
}

// Determines which player attacks first based on their health.
func (g *Game) determineFirstAttacker() (*Player, *Player) {
	if g.PlayerA.Health < g.PlayerB.Health {
		return g.PlayerA, g.PlayerB
	}
	return g.PlayerB, g.PlayerA
}


// Simulates an attack from the attacker to the defender.
func (g *Game) attack(attacker, defender *Player) {
	attackRoll := g.Die.Roll()
	defendRoll := g.Die.Roll()

	attackDamage := attacker.Attack * attackRoll
	defendDamage := defender.Strength * defendRoll

	actualDamage := max(0, attackDamage-defendDamage)
	defender.ReduceHealth(actualDamage)

	fmt.Printf("%s attacks %s for %d damage (Attack Damage %d - Defended Damage %d). %s's health is now %d\n",
		attacker.Name, defender.Name, actualDamage, attackDamage, defendDamage, defender.Name, defender.Health)
}
