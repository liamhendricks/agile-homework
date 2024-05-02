package swapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"agile-homework/src/clients/swapi/models"
)

// interact with swapi.dev api
type SwapiClient struct {
	baseUrl    string
	httpClient *http.Client
}

func NewSwapiClient() *SwapiClient {
	return &SwapiClient{
		baseUrl:    models.SwapiBaseUrl,
		httpClient: http.DefaultClient,
	}
}

type SwapiResponse[G models.Gettable] struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Results []G     `json:"results"`
}

// retrieves all characters from the swapi api by search term
func (c *SwapiClient) Characters(ctx context.Context, query string) (SwapiResponse[models.Character], error) {
	var swapiResponse SwapiResponse[models.Character]

	url := fmt.Sprintf("%s%s?search=%s", c.baseUrl, models.AllCharactersPath, query)

	res, err := c.getResponse(ctx, url)
	if err != nil {
		return swapiResponse, fmt.Errorf("error getting data: %w", err)
	}
	defer res.Body.Close()

	err = decodeResponse(res, &swapiResponse)
	if err != nil {
		return swapiResponse, fmt.Errorf("error decoding response: %w", err)
	}

	// return 404 if no records found
	if len(swapiResponse.Results) == 0 {
		return swapiResponse, SwapiNotFoundError{
			Code: http.StatusNotFound,
			Msg:  ErrNoData.Error(),
		}
	}

	*swapiResponse.Next = strings.Replace(
		*swapiResponse.Next,
		models.SwapiBaseUrl+models.AllCharactersPath,
		"",
		1)

	return swapiResponse, nil
}

// retrieves a single planet from the swapi api by id
func (c *SwapiClient) Planet(ctx context.Context, id string) (models.Planet, error) {
	var planetResponse models.Planet

	url := fmt.Sprintf("%s%s%s", c.baseUrl, models.AllPlanetsPath, id)
	res, err := c.getResponse(ctx, url)
	if err != nil {
		return planetResponse, fmt.Errorf("error getting data: %w", err)
	}
	defer res.Body.Close()

	err = decodeResponse(res, &planetResponse)
	if err != nil {
		return planetResponse, fmt.Errorf("error decoding response: %w", err)
	}

	return planetResponse, nil
}

// retrieves a single species from the swapi api by id
func (c *SwapiClient) Species(ctx context.Context, id string) (models.Species, error) {
	var speciesResponse models.Species

	url := fmt.Sprintf("%s%s%s", c.baseUrl, models.AllPlanetsPath, id)
	res, err := c.getResponse(ctx, url)
	if err != nil {
		return speciesResponse, fmt.Errorf("error getting data: %w", err)
	}
	defer res.Body.Close()

	err = decodeResponse(res, &speciesResponse)
	if err != nil {
		return speciesResponse, fmt.Errorf("error decoding response: %w", err)
	}

	return speciesResponse, nil
}

// retrieves a single starship from the swapi api by id
func (c *SwapiClient) Starship(ctx context.Context, id string) (models.Starship, error) {
	var starshipResponse models.Starship

	url := fmt.Sprintf("%s%s%s", c.baseUrl, models.AllPlanetsPath, id)
	res, err := c.getResponse(ctx, url)
	if err != nil {
		return starshipResponse, fmt.Errorf("error getting data: %w", err)
	}

	defer res.Body.Close()

	err = decodeResponse(res, &starshipResponse)
	if err != nil {
		return starshipResponse, fmt.Errorf("error decoding response: %w", err)
	}

	return starshipResponse, nil
}

func (c *SwapiClient) getResponse(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Close = true
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting data: %w", err)
	}

	return res, nil
}

func decodeResponse(res *http.Response, val interface{}) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading data: %w", err)
	}

	err = json.Unmarshal(body, val)
	if err != nil {
		return fmt.Errorf("error decoding data: %w", err)
	}

	return nil
}
