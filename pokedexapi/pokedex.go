package pokedexapi

import (
	"fmt"

	"github.com/D3rise/pokedexcli/pokedexapi/responses"
	"github.com/D3rise/pokedexcli/requests"
)

const DEFAULT_POKEDEX_API_ENDPOINT = "https://pokeapi.co/api/v2/"
const LOCATION_AREA_ENDPOINT = "location-area/"

type PokedexAPI struct {
	apiEndpoint string
}

func NewPokedexAPI(apiEndpoint string) *PokedexAPI {
	if apiEndpoint == "" {
		return &PokedexAPI{
			apiEndpoint: DEFAULT_POKEDEX_API_ENDPOINT,
		}
	} else {
		return &PokedexAPI{
			apiEndpoint: apiEndpoint,
		}
	}
}

func (p PokedexAPI) GetLocationAreaList(limit int, offset int) (responses.LocationAreaListResponse, error) {
	res, err := requests.GetAndUnmarshal[responses.LocationAreaListResponse](p.apiEndpoint + LOCATION_AREA_ENDPOINT + fmt.Sprintf("?limit=%d&offset=%d", limit, offset))
	if err != nil {
		return *new(responses.LocationAreaListResponse), err
	}

	return res, nil
}
