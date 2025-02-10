package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/yogyrahmawan/tcgdex-go-sdk/pkg/model"
)

type Fetcheable interface {
	FetchSingleCard(cardID string) (*model.Card, error)
	SearchCards(options model.CardQueryOptions) ([]model.CardBrief, error)
	GetSets(setID string) (*model.Set, error)
	SearchSets(options model.SetQueryOptions) ([]model.SetBrief, error)
	GetCardBySetAndLocalId(setID, localID string) (*model.Card, error)
	GetSingleSerie(serieID string) (*model.Serie, error)
	SearchSeries(options model.SerieQueryOptions) ([]model.SerieBrief, error)
	ListCardTypes() ([]string, error)
	ListCardRetreatCosts() ([]int, error)
	ListCardRarities() ([]string, error)
	ListCardIllustrators() ([]string, error)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode fetch single card error response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var card model.Card
	if err = json.NewDecoder(httpResp.Body).Decode(&card); err != nil {
		return nil, fmt.Errorf("decode fetch single card: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode fetch search cards error response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var cardBriefs []model.CardBrief
	if err = json.NewDecoder(httpResp.Body).Decode(&cardBriefs); err != nil {
		return nil, fmt.Errorf("decode fetch search cards: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode get sets error response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var set model.Set
	if err = json.NewDecoder(httpResp.Body).Decode(&set); err != nil {
		return nil, fmt.Errorf("decode get sets: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode search sets error response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var setBriefs []model.SetBrief
	if err = json.NewDecoder(httpResp.Body).Decode(&setBriefs); err != nil {
		return nil, fmt.Errorf("decode search sets: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode get card by set and localId error response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var card model.Card
	if err = json.NewDecoder(httpResp.Body).Decode(&card); err != nil {
		return nil, fmt.Errorf("decode get card by set and localId: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode get single serie error response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var serie model.Serie
	if err = json.NewDecoder(httpResp.Body).Decode(&serie); err != nil {
		return nil, fmt.Errorf("decode get single serie: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode search series error response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var serieBriefs []model.SerieBrief
	if err = json.NewDecoder(httpResp.Body).Decode(&serieBriefs); err != nil {
		return nil, fmt.Errorf("decode search series: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode list card types response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var cardTypes []string
	if err = json.NewDecoder(httpResp.Body).Decode(&cardTypes); err != nil {
		return nil, fmt.Errorf("decode list card types: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode list card retreat costs response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var cardRetreatCosts []int
	if err = json.NewDecoder(httpResp.Body).Decode(&cardRetreatCosts); err != nil {
		return nil, fmt.Errorf("decode list card retreat costs: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode list card rarities response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var cardRarities []string
	if err = json.NewDecoder(httpResp.Body).Decode(&cardRarities); err != nil {
		return nil, fmt.Errorf("decode list card rarities: %w", err)
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

	if httpResp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError

		if err = json.NewDecoder(httpResp.Body).Decode(&httpResp); err != nil {
			return nil, fmt.Errorf("decode list card illustrators response: %w", err)
		}

		return nil, errors.New(httpErr.String())
	}

	var cardIllustrators []string
	if err = json.NewDecoder(httpResp.Body).Decode(&cardIllustrators); err != nil {
		return nil, fmt.Errorf("decode list card illustrators: %w", err)
	}

	return cardIllustrators, nil
}
