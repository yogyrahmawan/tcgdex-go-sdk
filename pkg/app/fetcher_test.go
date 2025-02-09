package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yogyrahmawan/tcgdex-go-sdk/pkg/model"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestFetchSingleCardOK(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_single_card_status_ok")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en/cards")
	card, err := f.FetchSingleCard("swsh3-136")
	assert.NoError(t, err)
	assert.Equal(t, "swsh3-136", card.ID)
	assert.Equal(t, "Pokemon", card.Category)
}

func TestFetchSingleCardNotFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_single_card_status_not_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en/cards")
	card, err := f.FetchSingleCard("swsh3")
	assert.Nil(t, card)
	assert.Error(t, err)
}

func TestSearchCardsByNameNotFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_cards_by_name_not_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en/cards")
	cards, err := f.SearchCards(model.CardQueryOptions{
		Name: "pokemon",
	})
	assert.NoError(t, err)
	assert.Len(t, cards, 0)
}

func TestSearchCardsByNameFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_cards_by_name_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en/cards")
	cards, err := f.SearchCards(model.CardQueryOptions{
		Name: "pikachu",
	})
	assert.NoError(t, err)
	assert.Len(t, cards, 179)
}

func TestSearchCardsByIdFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_cards_by_id_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en/cards")
	cards, err := f.SearchCards(model.CardQueryOptions{
		Id: "xyp-XY124",
	})
	assert.NoError(t, err)
	assert.Len(t, cards, 1)
}

func TestSearchCardsByLocalIdFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_cards_by_localid_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en/cards")
	cards, err := f.SearchCards(model.CardQueryOptions{
		LocalId: "XY124",
	})
	assert.NoError(t, err)
	assert.Len(t, cards, 1)
}

func TestSearchCardsUsingPagination(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_cards_using_pagination")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en/cards")
	cards, err := f.SearchCards(model.CardQueryOptions{
		Name:                   "pikachu",
		PaginationPage:         1,
		PaginationItemsPerPage: 2,
	})
	assert.NoError(t, err)
	assert.Len(t, cards, 2)
}
