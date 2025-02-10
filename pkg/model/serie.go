package model

type Serie struct {
	ID   string     `json:"id"`
	Logo string     `json:"logo"`
	Name string     `json:"name"`
	Sets []SetBrief `json:"sets"`
}

type SerieQueryOptions struct {
	Id                     string `url:"id,omitempty"`
	Name                   string `url:"name,omitempty"`
	PaginationPage         int    `url:"pagination:page,omitempty"`
	PaginationItemsPerPage int    `url:"pagination:itemsPerPage,omitempty"`
}

type SerieBrief struct {
	ID   string `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}
