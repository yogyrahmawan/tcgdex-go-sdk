package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
