package sdk

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
	sets, err := f.SearchSets(model.SetQueryOptions{
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

func TestSearchSetsByIdNotFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_series_by_id_not_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	series, err := f.SearchSeries(model.SerieQueryOptions{
		Id: "abc",
	})
	assert.NoError(t, err)
	assert.Len(t, series, 0)
}

func TestSearchSeriesByIdFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_series_by_id_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	series, err := f.SearchSeries(model.SerieQueryOptions{
		Id: "neo",
	})
	assert.NoError(t, err)
	assert.Len(t, series, 1)
}

func TestSearchSeriesByNameNotFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_series_by_name_not_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	series, err := f.SearchSeries(model.SerieQueryOptions{
		Name: "notfound",
	})
	assert.NoError(t, err)
	assert.Len(t, series, 0)
}

func TestSearchSeriesByNameFound(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_series_by_name_found")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	series, err := f.SearchSeries(model.SerieQueryOptions{
		Name: "Neo",
	})
	assert.NoError(t, err)
	assert.Len(t, series, 2)
}

func TestSearchSeriesUsingPagination(t *testing.T) {
	r, err := recorder.New("fixtures/fetch_search_series_using_pagination")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	series, err := f.SearchSeries(model.SerieQueryOptions{
		PaginationPage:         1,
		PaginationItemsPerPage: 2,
	})
	assert.NoError(t, err)
	assert.Len(t, series, 2)
}

func TestListCardTypes(t *testing.T) {
	r, err := recorder.New("fixtures/list_card_types")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	cardTypes, err := f.ListCardTypes()
	assert.NoError(t, err)
	assert.Len(t, cardTypes, 11)
}

func TestListCardRetreatCosts(t *testing.T) {
	r, err := recorder.New("fixtures/list_card_retreat_costs")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	cardTypes, err := f.ListCardRetreatCosts()
	assert.NoError(t, err)
	assert.Len(t, cardTypes, 5)
}

func TestListCardRarities(t *testing.T) {
	r, err := recorder.New("fixtures/list_card_rarities")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	cardRarities, err := f.ListCardRarities()
	assert.NoError(t, err)
	assert.Len(t, cardRarities, 35)
}

func TestListCardIllustrators(t *testing.T) {
	r, err := recorder.New("fixtures/list_card_illustrators")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	cardRarities, err := f.ListCardIllustrators()
	assert.NoError(t, err)
	assert.Len(t, cardRarities, 299)
}

func TestListCardCategories(t *testing.T) {
	r, err := recorder.New("fixtures/list_card_categories")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	cardCategories, err := f.ListCardCategories()
	assert.NoError(t, err)
	assert.Len(t, cardCategories, 3)
}

func TestListPokemonStages(t *testing.T) {
	r, err := recorder.New("fixtures/list_pokemon_stages")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	pokemonStages, err := f.ListPokemonStages()
	assert.NoError(t, err)
	assert.Len(t, pokemonStages, 10)
}

func TestSuffixes(t *testing.T) {
	r, err := recorder.New("fixtures/list_suffixes")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	pokemonStages, err := f.ListSuffixes()
	assert.NoError(t, err)
	assert.Len(t, pokemonStages, 7)
}

func TestListVariants(t *testing.T) {
	r, err := recorder.New("fixtures/list_variants")
	assert.NoError(t, err)
	defer func() {
		err := r.Stop()
		assert.NoError(t, err)
	}()

	f := NewFetcher(r.GetDefaultClient(), 5*time.Second, "https://api.tcgdex.net/v2/en")
	variants, err := f.ListVariants()
	assert.NoError(t, err)
	assert.Len(t, variants, 5)
}
