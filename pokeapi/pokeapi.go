package pokeapi

import (
	"errors"
	"fmt"
	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/pokeapi/responses"
	"net/http"
	"time"

	"github.com/D3rise/pokedexcli/internal/cache"
	"github.com/D3rise/pokedexcli/internal/requests"
)

const (
	defaultPokeApiEndpoint = "https://pokeapi.co/api/v2/"
	locationAreaEndpoint   = "location-area/"
	pokemonEndpoint        = "pokemon/"
)

const PokeApiContextKey context.ContextKey = "_PokeApiContextKey"

type PokeAPI struct {
	apiEndpoint string
	cache       *cache.Cache
}

func NewPokeAPI(apiEndpoint string) *PokeAPI {
	var endpoint string
	if apiEndpoint == "" {
		endpoint = defaultPokeApiEndpoint
	} else {
		endpoint = apiEndpoint
	}

	c := cache.NewCache(10 * time.Second)

	return &PokeAPI{
		apiEndpoint: endpoint,
		cache:       c,
	}
}

func (p *PokeAPI) GetPokemonInfo(idOrName string) (responses.PokemonInfoResponse, error) {
	url := p.apiEndpoint + pokemonEndpoint + idOrName
	return getUsingCacheOrRequest[responses.PokemonInfoResponse](url, p.cache)
}

func (p *PokeAPI) GetLocationAreaInfo(areaName string) (responses.LocationAreaInfoResponse, error) {
	url := p.apiEndpoint + locationAreaEndpoint + areaName
	return getUsingCacheOrRequest[responses.LocationAreaInfoResponse](url, p.cache)
}

func (p *PokeAPI) GetLocationAreaList(limit int, offset int) (responses.LocationAreaListResponse, error) {
	url := p.apiEndpoint + locationAreaEndpoint + fmt.Sprintf("?limit=%d&offset=%d", limit, offset)
	return getUsingCacheOrRequest[responses.LocationAreaListResponse](url, p.cache)
}

func getUsingCacheOrRequest[T any](url string, cache *cache.Cache) (T, error) {
	var res *http.Response
	var body []byte
	var err error

	if cached, ok := cache.Get(url); ok {
		body = cached
	} else {
		res, body, err = requests.Get(url)
		if err != nil {
			return *new(T), err
		}

		if res.StatusCode == http.StatusNotFound {
			return *new(T), errors.New("not found")
		}

		cache.Add(url, body)
	}

	result, err := requests.UnmarshalBody[T](body)
	if err != nil {
		return *new(T), err
	}

	return result, nil
}
