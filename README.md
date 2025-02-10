# TCGdex Go SDK

The TCGdex Go SDK provides a convenient access with the Open Source TCGdex API.

_The full API/SDK documentation is available at [API Documentation - TCGdex](https://www.tcgdex.dev)_

## Getting Started 
### How to install 
run the following command:
```
go get github.com/yogyrahmawan/tcgdex-go-sdk
```

### Example 
```
package main

import (
	"net/http"
	"time"

	"github.com/yogyrahmawan/tcgdex-go-sdk/pkg/model"
	"github.com/yogyrahmawan/tcgdex-go-sdk/pkg/sdk"
)

const (
	tcgdexEnBaseURL = "https://api.tcgdex.net/v2/en"
)

func main() {
	fetcher := sdk.NewFetcher(http.DefaultClient, 5*time.Second, tcgdexEnBaseURL)

	// get single card
	fetcher.FetchSingleCard("swsh3-136")

	// search cards by card query options
	fetcher.SearchCards(model.CardQueryOptions{Name: "pikachu"})

	// get sets
	fetcher.GetSets("swsh1")

	// search sets
	fetcher.SearchSets(model.SetQueryOptions{
		Name: "Jungle",
	})

	// get card by set and local id
	fetcher.GetCardBySetAndLocalId("swsh3", "136")

	// get single serie
	fetcher.GetSingleSerie("swsh")

	// search series
	fetcher.SearchSeries(model.SerieQueryOptions{
		Name: "Neo",
	})

	// list card types
	fetcher.ListCardTypes()

	// list card retreat types
	fetcher.ListCardRetreatCosts()

	// list card rarities
	fetcher.ListCardRarities()

	// list card illustrators
	fetcher.ListCardIllustrators()

	// list card categories
	fetcher.ListCardCategories()

	// list pokemon stages
	fetcher.ListPokemonStages()

	// list suffixes
	fetcher.ListSuffixes()

	// list variants
	fetcher.ListVariants()
}

```

## Contributing 
* Fork
* Commit
* Open Pull Request

## License 
This project is licensed under the MIT License. A copy of the license is available at [LICENSE](https://github.com/yogyrahmawan/tcgdex-go-sdk/blob/main/LICENSE)
