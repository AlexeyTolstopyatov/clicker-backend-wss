package Handling

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"math/rand"
	"wsst/Modeling"
)

var (
	clicksCount = 0
	clientTeam  = "white"
)

var ClientBattery = Modeling.Battery{
	Type:     "battery_status",
	Capacity: 100,
	Charge:   100,
}

var ClientScore = Modeling.Score{
	Type:         "score",
	BlackClicks:  0,
	BlackPercent: 0,
	WhiteClicks:  0,
	WhitePercent: 0,
}

// Set
// Распознает тип сообщения, вызывает обработчики
// сообщений, описанные ниже, подсчитывает статистику
//
// Отправляет:
//
//	Score{}		 		JSON сообщение
//	Battery{}			JSON сообщение
//
// Принимает:
//
//	ClientClick string	Текст
//	Context		*Conn	Указатель на обработчик событий
func Set(context *websocket.Conn) {
	clicksCount += 1
	if ClientBattery.Charge > 0 {
		ClientBattery.Charge -= 10
	}

	switch clientTeam {
	case "black":
		ClientScore.BlackClicks += clicksCount
		break
	case "white":
		ClientScore.WhiteClicks += clicksCount
		break
	}
	// c = (b / (b+w)) * 100
	ClientScore.BlackPercent =
		float64(ClientScore.BlackClicks/(ClientScore.BlackClicks+ClientScore.WhiteClicks)) * 100

	// c = (w / (b+w)) * 100
	ClientScore.WhitePercent =
		float64(ClientScore.WhiteClicks/(ClientScore.BlackClicks+ClientScore.WhiteClicks)) * 100

	batteryJson, _ := json.Marshal(ClientBattery)
	scoreJson, _ := json.Marshal(ClientScore)

	_ = context.WriteMessage(websocket.TextMessage, batteryJson)
	_ = context.WriteMessage(websocket.TextMessage, scoreJson)
}

// GetClientTeam
// Получает команду за которую играет клиент
func GetClientTeam() string {
	return clientTeam
}

// SetClientTeam
// Устанавливает команду за которую играет клиент
func SetClientTeam() {
	if rand.Int()%2 == 0 {
		clientTeam = "white"
	} else {
		clientTeam = "black"
	}
}

// SetInitMessage
// Возвращает структуру запроса инициализации
func SetInitMessage(team string) Modeling.Init {
	return Modeling.Init{
		Type:            "init",
		BatteryCapacity: 100,
		BatteryCharge:   100,
		Clicks:          0,
		Team:            team,
	}
}
