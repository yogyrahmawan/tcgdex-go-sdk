package model

type Card struct {
	Illustrator    string         `json:"illustrator"`
	Category       string         `json:"category"`
	ID             string         `json:"id"`
	Image          string         `json:"image"`
	LocalID        string         `json:"localId"`
	Name           string         `json:"name"`
	Rarity         string         `json:"rarity"`
	Set            Set            `json:"set"`
	Variants       CardVariants   `json:"variants"`
	Hp             int            `json:"hp"`
	Types          []string       `json:"types"`
	EvolveFrom     string         `json:"evolveFrom"`
	Description    string         `json:"description"`
	Stage          string         `json:"stage"`
	Attacks        []CardAttack   `json:"attacks"`
	Weaknesses     []CardWeakness `json:"weaknesses"`
	Retreat        int            `json:"retreat"`
	RegulationMark string         `json:"regulationMark"`
	Legal          Legal          `json:"legal"`
	DexIds         []int          `json:"dexIDs"`
	Level          []string       `json:"level"`
	Suffix         []string       `json:"suffix"`
	Item           []CardItem     `json:"CardItem"`
	Abilities      []CardAbility  `json:"CardAbility"`
	Effect         string         `json:"effect"`
	TrainerType    string         `json:"trainerType"`
	EnergyType     string         `json:"energyType"`
	Resistances    []CardWeakness `json:"resistances"`
}

type CardAttack struct {
	Cost   []string `json:"cost"`
	Name   string   `json:"name"`
	Effect string   `json:"effect"`
	Damage int      `json:"damage,omitempty"`
}

type CardWeakness struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CardVariants struct {
	FirstEdition bool `json:"firstEdition"`
	Holo         bool `json:"holo"`
	Normal       bool `json:"normal"`
	Reverse      bool `json:"reverse"`
	WPromo       bool `json:"wPromo"`
}

type CardItem struct {
	Name   string `json:"name"`
	Effect string `json:"effect"`
}

type CardAbility struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Effect string `json:"effect"`
}

type CardQueryOptions struct {
	Id                     string `url:"id,omitempty"`
	LocalId                string `url:"localId,omitempty"`
	Name                   string `url:"name,omitempty"`
	PaginationPage         int    `url:"pagination:page,omitempty"`
	PaginationItemsPerPage int    `url:"pagination:itemsPerPage,omitempty"`
}

type CardBrief struct {
	ID      string `json:"id"`
	LocalID string `json:"localId"`
	Name    string `json:"name"`
	Image   string `json:"image"`
}
