package game

import (
	"wsst/cmd/client"
	"wsst/shared"
)

var score = shared.Score{
	Type:         "score",
	BlackClicks:  0,
	WhiteClicks:  0,
	BlackPercent: 0,
	WhitePercent: 0,
}

func updateWhiteScore(client *client.Client) {

}

func updateBlackScore(client *client.Client) {

}
