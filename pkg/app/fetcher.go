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
		return nil, fmt.Errorf("parse fetch single card: %w", err)
	}

	httpResp, err := f.httpClient.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("get single card: %w", err)
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
