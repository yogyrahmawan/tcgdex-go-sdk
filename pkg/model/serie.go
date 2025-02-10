package model

type Serie struct {
	ID   string     `json:"id"`
	Logo string     `json:"logo"`
	Name string     `json:"name"`
	Sets []SetBrief `json:"sets"`
}
