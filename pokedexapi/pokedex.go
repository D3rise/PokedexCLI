package pokedexapi

import (
	"fmt"
	"github.com/D3rise/pokedexcli/pokedexapi/responses"
	"time"

	"github.com/D3rise/pokedexcli/internal/pokecache"
	"github.com/D3rise/pokedexcli/internal/requests"
)

const DEFAULT_POKEDEX_API_ENDPOINT = "https://pokeapi.co/api/v2/"
const LOCATION_AREA_ENDPOINT = "location-area/"

type PokedexAPI struct {
	apiEndpoint string
	cache       *pokecache.Cache
}

func NewPokedexAPI(apiEndpoint string) *PokedexAPI {
	var endpoint string
	if apiEndpoint == "" {
		endpoint = DEFAULT_POKEDEX_API_ENDPOINT
	} else {
		endpoint = apiEndpoint
	}

	cache := pokecache.NewCache(10 * time.Second)

	return &PokedexAPI{
		apiEndpoint: endpoint,
		cache:       cache,
	}
}

func (p *PokedexAPI) GetLocationAreaInfo(areaName string) (responses.LocationAreaInfoResponse, error) {
	url := p.apiEndpoint + LOCATION_AREA_ENDPOINT + areaName
	return getUsingCacheOrRequest[responses.LocationAreaInfoResponse](url, (*p).cache)
}

func (p *PokedexAPI) GetLocationAreaList(limit int, offset int) (responses.LocationAreaListResponse, error) {
	url := p.apiEndpoint + LOCATION_AREA_ENDPOINT + fmt.Sprintf("?limit=%d&offset=%d", limit, offset)
	return getUsingCacheOrRequest[responses.LocationAreaListResponse](url, (*p).cache)
}

func getUsingCacheOrRequest[T any](url string, cache *pokecache.Cache) (T, error) {
	var body []byte
	var err error

	if cached, ok := (*cache).Get(url); ok {
		body = cached
	} else {
		body, err = requests.Get(url)
		if err != nil {
			return *new(T), err
		}

		(*cache).Add(url, body)
	}

	result, err := requests.UnmarshalBody[T](body)
	if err != nil {
		return *new(T), err
	}

	return result, nil
}
