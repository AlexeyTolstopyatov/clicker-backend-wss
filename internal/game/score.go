package game

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
	"time"
	"wsst/internal/client"
	"wsst/model"
)

// score
// Общая структура счета команд.
// Содержит счеты команд и доли (в процентах).
var score = model.Score{
	Type:         "score",
	BlackClicks:  0,
	WhiteClicks:  0,
	BlackPercent: 0,
	WhitePercent: 0,
}

// UpdateScore
// Запускает счетчик обновлений счета команд,
// проверяет состояние каждые 2 секунды.
func UpdateScore(clientPtr *client.Client) {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		switch clientPtr.InitPtr.Team {
		case "white":
			updateWhiteScore(clientPtr.InitPtr)
			break

		case "black":
			updateBlackScore(clientPtr.InitPtr)
			break
		}
		s := score
		scoreJson, err := json.Marshal(s)
		if err != nil {
			log.Error(err.Error())
			return
		}
		goroutineMutex.Lock()
		_ = clientPtr.ConnectPtr.WriteMessage(websocket.TextMessage, scoreJson)
		goroutineMutex.Unlock()

	}
}

// updateWhiteScore
// Обновляет счет для команды белых.
// initPtr - указатель на структуру "init"-запроса
// откуда будут взяты данные о кликах.
func updateWhiteScore(initPtr *model.Init) {
	sum := score.WhiteClicks + score.BlackClicks

	score.WhiteClicks = initPtr.Clicks
	if sum == 0 {
		return
	}
	score.WhitePercent = float64(score.WhiteClicks/sum) * 100
}

// updateBlackScore
// Обновляет счет для команды черных.
// initPtr - указатель на структуру "init"-запроса
// откуда будут взяты данные о кликах.
func updateBlackScore(initPtr *model.Init) {
	sum := score.WhiteClicks + score.BlackClicks

	score.BlackClicks = initPtr.Clicks
	if sum == 0 {
		return
	}
	score.WhitePercent = float64(score.BlackClicks/sum) * 100
}
