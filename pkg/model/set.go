package model

type Set struct {
	CardCount    CardCount    `json:"cardCount"`
	Cards        []Card       `json:"cards"`
	ID           string       `json:"id"`
	Legal        Legal        `json:"legal"`
	Logo         string       `json:"logo"`
	Name         string       `json:"name"`
	ReleaseDate  string       `json:"releaseDate"`
	Serie        Serie        `json:"serie"`
	Symbol       string       `json:"symbol"`
	TcgOnline    string       `json:"tcgOnline"`
	Abbreviation Abbreviation `json:"abbreviation"`
}

type CardCount struct {
	FirstEd  int `json:"firstEd"`
	Holo     int `json:"holo"`
	Normal   int `json:"normal"`
	Official int `json:"official"`
	Reverse  int `json:"reverse"`
	Total    int `json:"total"`
}

type Legal struct {
	Standard bool `json:"standard"`
	Expanded bool `json:"expanded"`
}

type Serie struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Abbreviation struct {
	Official string `json:"official"`
}
