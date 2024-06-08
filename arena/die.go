package arena

import (
	"math/rand"
	"time"
)

// Die struct specifies no. of sides.
type Die struct {
	sides int
	rng   *rand.Rand
}

// Creates a New Die.
func NewDie(sides int) *Die {
	return &Die{
		sides: sides,
		rng:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Rolls a die by returning a random integer between 1 and no. of sides.
func (d *Die) Roll() int {
	return d.rng.Intn(d.sides) + 1
}
