package pokedexapi

import (
	"errors"
	"fmt"
	"github.com/D3rise/pokedexcli/internal/context"
	"github.com/D3rise/pokedexcli/pokedexapi/responses"
	"net/http"
	"time"

	"github.com/D3rise/pokedexcli/internal/cache"
	"github.com/D3rise/pokedexcli/internal/requests"
)

const (
	defaultPokedexApiEndpoint = "https://pokeapi.co/api/v2/"
	locationAreaEndpoint      = "location-area/"
	pokemonEndpoint           = "pokemon/"
)

const PokedexApiContextKey context.ContextKey = "_PokedexApiContextKey"

type PokedexAPI struct {
	apiEndpoint string
	cache       *cache.Cache
}

func NewPokedexAPI(apiEndpoint string) *PokedexAPI {
	var endpoint string
	if apiEndpoint == "" {
		endpoint = defaultPokedexApiEndpoint
	} else {
		endpoint = apiEndpoint
	}

	c := cache.NewCache(10 * time.Second)

	return &PokedexAPI{
		apiEndpoint: endpoint,
		cache:       c,
	}
}

func (p *PokedexAPI) GetPokemonInfo(idOrName string) (responses.PokemonInfoResponse, error) {
	url := p.apiEndpoint + pokemonEndpoint + idOrName
	return getUsingCacheOrRequest[responses.PokemonInfoResponse](url, p.cache)
}

func (p *PokedexAPI) GetLocationAreaInfo(areaName string) (responses.LocationAreaInfoResponse, error) {
	url := p.apiEndpoint + locationAreaEndpoint + areaName
	return getUsingCacheOrRequest[responses.LocationAreaInfoResponse](url, p.cache)
}

func (p *PokedexAPI) GetLocationAreaList(limit int, offset int) (responses.LocationAreaListResponse, error) {
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
