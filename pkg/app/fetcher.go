package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/yogyrahmawan/tcgdex-go-sdk/pkg/model"
)

type Fetcheable interface {
	FetchSingleCard(cardID string) (*model.Card, error)
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
	url, err := url.Parse(f.baseURL + "/" + cardID)
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
