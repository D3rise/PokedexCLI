package pokedex

import (
	"github.com/D3rise/pokedexcli/internal/context"
	"sync"
)

type Pokemon struct {
	Name string `json:"name"`
}

type Pokedex struct {
	mu             sync.RWMutex
	caughtPokemons map[string]Pokemon
}

const PokedexContextKey context.ContextKey = "_PokedexContextKey"

func NewPokedex() *Pokedex {
	return &Pokedex{
		mu:             sync.RWMutex{},
		caughtPokemons: make(map[string]Pokemon),
	}
}

func (p *Pokedex) AddNewPokemon(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.caughtPokemons[name] = Pokemon{}
}

func (p *Pokedex) GetCaughtPokemons() map[string]Pokemon {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.caughtPokemons
}
