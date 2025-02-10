package app

type Lister interface {
	ListCardTypes() ([]string, error)
	ListCardRetreatCosts() ([]int, error)
	ListCardRarities() ([]string, error)
	ListCardIllustrators() ([]string, error)
	ListCardCategories() ([]string, error)
}
