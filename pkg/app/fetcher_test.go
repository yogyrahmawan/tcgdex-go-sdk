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

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
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

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
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

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
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

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
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

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
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

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
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

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	cards, err := f.SearchCards(model.CardQueryOptions{
		Name:                   "pikachu",
		PaginationPage:         1,
		PaginationItemsPerPage: 2,
	})
	assert.NoError(t, err)
	assert.Len(t, cards, 2)
}

func TestGetSetsFound(t *testing.T) {
	r, err := recorder.New("fixtures/get_sets_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	sets, err := f.GetSets("swsh1")
	assert.NoError(t, err)
	assert.Equal(t, "swsh1", sets.ID)
	assert.Equal(t, 216, sets.CardCount.Total)
	assert.Equal(t, "Sword & Shield", sets.Name)
}

func TestGetSetsNotFound(t *testing.T) {
	r, err := recorder.New("fixtures/get_sets_not_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	sets, err := f.GetSets("notfound")
	assert.Error(t, err)
	assert.Nil(t, sets)
}

func TestSearchSetsByNameNotFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_sets_by_name_not_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	sets, err := f.SearchSets(model.SetQueryOptions{
		Name: "pokemon",
	})
	assert.NoError(t, err)
	assert.Len(t, sets, 0)
}

func TestSearchSetsByNameFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_sets_by_name_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	sets, err := f.SearchCards(model.CardQueryOptions{
		Name: "Jungle",
	})
	assert.NoError(t, err)
	assert.Len(t, sets, 1)
}

func TestSearchSetsByIdFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_sets_by_id_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	sets, err := f.SearchSets(model.SetQueryOptions{
		Id: "base2",
	})
	assert.NoError(t, err)
	assert.Len(t, sets, 1)
}

func TestSearchSetsUsingPagination(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_sets_using_pagination")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	cards, err := f.SearchCards(model.CardQueryOptions{
		PaginationPage:         1,
		PaginationItemsPerPage: 2,
	})
	assert.NoError(t, err)
	assert.Len(t, cards, 2)
}

func TestFetchSingleCardBySetAndLocalIdOK(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_single_card_by_setid_and_local_idstatus_ok")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	card, err := f.GetCardBySetAndLocalId("swsh3", "136")
	assert.NoError(t, err)
	assert.Equal(t, "swsh3-136", card.ID)
	assert.Equal(t, "Pokemon", card.Category)
}

func TestFetchSingleCardBySetAndLocalIdNotFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_single_card_by_setid_and_local_id_status_not_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	card, err := f.GetCardBySetAndLocalId("swsh3111", "32")
	assert.Nil(t, card)
	assert.Error(t, err)
}

func TestGetSerieFound(t *testing.T) {
	r, err := recorder.New("fixtures/get_serie_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	serie, err := f.GetSingleSerie("swsh")
	assert.NoError(t, err)
	assert.Equal(t, "swsh", serie.ID)
	assert.Len(t, serie.Sets, 19)
}

func TestGetSerieNotFound(t *testing.T) {
	r, err := recorder.New("fixtures/get_serie_not_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	serie, err := f.GetSingleSerie("notfound")
	assert.Error(t, err)
	assert.Nil(t, serie)
}
