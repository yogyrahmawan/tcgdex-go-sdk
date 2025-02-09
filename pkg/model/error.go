package model

import "fmt"

type TcgdexHttpError struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
}

func (e TcgdexHttpError) String() string {
	return fmt.Sprintf("%s:%s:%d:%s:%s", e.Type, e.Title, e.Status, e.Endpoint, e.Method)
}
