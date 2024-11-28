package model

type Battery struct {
	Type     string `json:"type"`
	Capacity int    `json:"capacity"`
	Charge   int    `json:"charge"`
}
