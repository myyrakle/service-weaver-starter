package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ServiceWeaver/weaver"
)

// Shuffler component.
type Shuffler interface {
	Shuffle(context.Context, string) (string, error)
}

// Implementation of the Shuffler component.
type shuffler struct {
	weaver.Implements[Shuffler]
}

// Randomly shuffles string
func (self *shuffler) Shuffle(_ context.Context, s string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	runeSlice := []rune(s)

	for i := len(runeSlice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}

	logger := self.Logger()
	logger.Info(fmt.Sprintf("%s -> %s", s, string(runeSlice)))

	return string(runeSlice), nil
}
