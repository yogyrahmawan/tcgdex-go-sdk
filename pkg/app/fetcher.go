package app

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/yogyrahmawan/tcgdex-go-sdk/pkg/model"
)

type Lister interface {
	ListCardTypes() ([]string, error)
	ListCardRetreatCosts() ([]int, error)
	ListCardRarities() ([]string, error)
	ListCardIllustrators() ([]string, error)
	ListCardCategories() ([]string, error)
	ListPokemonStages() ([]string, error)
	ListSuffixes() ([]string, error)
	ListVariants() ([]string, error)
}

type Fetcheable interface {
	FetchSingleCard(cardID string) (*model.Card, error)
	SearchCards(options model.CardQueryOptions) ([]model.CardBrief, error)
	GetSets(setID string) (*model.Set, error)
	SearchSets(options model.SetQueryOptions) ([]model.SetBrief, error)
	GetCardBySetAndLocalId(setID, localID string) (*model.Card, error)
	GetSingleSerie(serieID string) (*model.Serie, error)
	SearchSeries(options model.SerieQueryOptions) ([]model.SerieBrief, error)
	Lister
}

type fetcher struct {
	baseURL    string
	httpClient *http.Client
}

func NewFetcher(client *http.Client, httpClientTimeout time.Duration, baseURL string) Fetcheable {
	if client == nil {
		client = &http.Client{
			Timeout: httpClientTimeout,
		}
	}

	return &fetcher{
		baseURL:    baseURL,
		httpClient: client,
	}
}

func (f *fetcher) FetchSingleCard(cardID string) (*model.Card, error) {
	url, err := url.Parse(f.baseURL + "/cards/" + cardID)
	if err != nil {
		return nil, fmt.Errorf("parse fetch single card: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("get single card: %w", err)
	}

	var card model.Card
	if err := decodeJSONResponse(httpResp, &card); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return &card, nil
}

func (f *fetcher) SearchCards(options model.CardQueryOptions) ([]model.CardBrief, error) {
	queryStrings, err := query.Values(options)
	if err != nil {
		return nil, fmt.Errorf("values: %w", err)
	}

	url, err := url.Parse(f.baseURL + "/cards?" + queryStrings.Encode())
	if err != nil {
		return nil, fmt.Errorf("parse search cards: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("search cards: %w", err)
	}

	var cardBriefs []model.CardBrief
	if err := decodeJSONResponse(httpResp, &cardBriefs); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return cardBriefs, nil
}

func (f *fetcher) GetSets(setID string) (*model.Set, error) {
	url, err := url.Parse(f.baseURL + "/sets/" + setID)
	if err != nil {
		return nil, fmt.Errorf("parse fetch get sets: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("get sets: %w", err)
	}

	var set model.Set
	if err := decodeJSONResponse(httpResp, &set); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return &set, nil
}

func (f *fetcher) SearchSets(options model.SetQueryOptions) ([]model.SetBrief, error) {
	queryStrings, err := query.Values(options)
	if err != nil {
		return nil, fmt.Errorf("values: %w", err)
	}

	url, err := url.Parse(f.baseURL + "/sets?" + queryStrings.Encode())
	if err != nil {
		return nil, fmt.Errorf("parse search sets: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("search sets: %w", err)
	}

	var setBriefs []model.SetBrief
	if err := decodeJSONResponse(httpResp, &setBriefs); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return setBriefs, nil
}

func (f *fetcher) GetCardBySetAndLocalId(setID, localID string) (*model.Card, error) {
	url, err := url.Parse(f.baseURL + "/sets/" + setID + "/" + localID)
	if err != nil {
		return nil, fmt.Errorf("parse fetch get card by set and localId: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("get card by set and localId: %w", err)
	}

	var card model.Card
	if err := decodeJSONResponse(httpResp, &card); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return &card, nil
}

func (f *fetcher) GetSingleSerie(serieID string) (*model.Serie, error) {
	url, err := url.Parse(f.baseURL + "/series/" + serieID)
	if err != nil {
		return nil, fmt.Errorf("parse get single serie: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("get single serie: %w", err)
	}

	var serie model.Serie
	if err := decodeJSONResponse(httpResp, &serie); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return &serie, nil
}

func (f *fetcher) SearchSeries(options model.SerieQueryOptions) ([]model.SerieBrief, error) {
	queryStrings, err := query.Values(options)
	if err != nil {
		return nil, fmt.Errorf("values: %w", err)
	}

	url, err := url.Parse(f.baseURL + "/series?" + queryStrings.Encode())
	if err != nil {
		return nil, fmt.Errorf("parse search series: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("search series: %w", err)
	}

	var serieBriefs []model.SerieBrief
	if err := decodeJSONResponse(httpResp, &serieBriefs); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return serieBriefs, nil
}

func (f *fetcher) ListCardTypes() ([]string, error) {
	url, err := url.Parse(f.baseURL + "/types")
	if err != nil {
		return nil, fmt.Errorf("parse list card types: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("list card types: %w", err)
	}

	var cardTypes []string
	if err := decodeJSONResponse(httpResp, &cardTypes); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return cardTypes, nil
}

func (f *fetcher) ListCardRetreatCosts() ([]int, error) {
	url, err := url.Parse(f.baseURL + "/retreats")
	if err != nil {
		return nil, fmt.Errorf("parse list card retreat costs: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("list card retreat costs: %w", err)
	}

	var cardRetreatCosts []int
	if err := decodeJSONResponse(httpResp, &cardRetreatCosts); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return cardRetreatCosts, nil
}

func (f *fetcher) ListCardRarities() ([]string, error) {
	url, err := url.Parse(f.baseURL + "/rarities")
	if err != nil {
		return nil, fmt.Errorf("parse list card rarities: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("list card rarities: %w", err)
	}

	var cardRarities []string
	if err := decodeJSONResponse(httpResp, &cardRarities); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return cardRarities, nil
}

func (f *fetcher) ListCardIllustrators() ([]string, error) {
	url, err := url.Parse(f.baseURL + "/illustrators")
	if err != nil {
		return nil, fmt.Errorf("parse list card illustrators: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("list card illustrators: %w", err)
	}

	var cardIllustrators []string
	if err := decodeJSONResponse(httpResp, &cardIllustrators); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return cardIllustrators, nil
}

func (f *fetcher) ListCardCategories() ([]string, error) {
	url, err := url.Parse(f.baseURL + "/categories")
	if err != nil {
		return nil, fmt.Errorf("parse list card categories: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("list card categories: %w", err)
	}

	var cardCategories []string
	if err := decodeJSONResponse(httpResp, &cardCategories); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return cardCategories, nil
}

func (f *fetcher) ListPokemonStages() ([]string, error) {
	url, err := url.Parse(f.baseURL + "/stages")
	if err != nil {
		return nil, fmt.Errorf("parse list pokemon stages: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("list pokemon stages: %w", err)
	}

	var pokemonStages []string
	if err := decodeJSONResponse(httpResp, &pokemonStages); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return pokemonStages, nil
}

func (f *fetcher) ListSuffixes() ([]string, error) {
	url, err := url.Parse(f.baseURL + "/suffixes")
	if err != nil {
		return nil, fmt.Errorf("parse list pokemon suffixes: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("list pokemon suffixes: %w", err)
	}

	var suffixes []string
	if err := decodeJSONResponse(httpResp, &suffixes); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return suffixes, nil
}

func (f *fetcher) ListVariants() ([]string, error) {
	url, err := url.Parse(f.baseURL + "/variants")
	if err != nil {
		return nil, fmt.Errorf("parse list variants: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("list variants: %w", err)
	}

	var variants []string
	if err := decodeJSONResponse(httpResp, &variants); err != nil {
		return nil, fmt.Errorf("decode json response: %w", err)
	}

	return variants, nil
}
