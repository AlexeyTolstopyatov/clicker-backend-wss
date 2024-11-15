package Modeling

type Score struct {
	Type         string  `json:"type"`
	BlackClicks  int     `json:"black_clicks"`
	BlackPercent float64 `json:"black_percent"`
	WhiteClicks  int     `json:"white_clicks"`
	WhitePercent float64 `json:"white_percent"`
}
