package Handling

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"wsst/Modeling"
)

var (
	// RequestBodyError
	// Тело запроса не удается распознать
	RequestBodyError = errors.New("unable to recognize JSON common request")
)

// CommonBody
// Структура "общего" JSON запроса
// все, что не соответствует структуре -- не обрабатывается
type CommonBody struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// HandleMessage
// Распознает тип сообщения, вызывает обработчики
// сообщений, описанные ниже
// Принимает тело запроса в вите CommonBody. Это необходимо для
// сообщения с клиентом
func HandleMessage(message []byte) error {
	body := CommonBody{}
	result := json.Unmarshal(message, &body)

	if result != nil {
		return RequestBodyError
	}

	if !IsFirstRequest() {
		if ProcessRequestTime() != nil {
			return RequestDeltaTimeError
		}
	}

	switch body.Type {
	case "init":
		{
			var init, deserializeError = DeserializeInitMessage(&message)
			log.Printf("Executed: (error: %s) %s \n", deserializeError, init)
			break
		}
	case "score":
		{
			var score, deserializeError = DeserializeScoreMessage(&message)
			log.Printf("Executed: (error: %s) %s \n", deserializeError, score)
			break
		}
	case "battery_status":
		{
			var battery, deserializeError = DeserializeBatteryStatusMessage(&message)
			log.Printf("Executed: (error: %s) %s \n", deserializeError, battery)
			break
		}
	default:
		{
			log.Printf("%s", body)
		}
	}

	return nil
}

// DeserializeBatteryStatusMessage
// Распознает структуру тела запроса как сообщение
// о "зарядке" батареи игрока (см. ТЗ)
// Возвращает структуру десериазированного запроса
// и состояние о выполненной операции
func DeserializeBatteryStatusMessage(message *[]byte) (Modeling.Battery, error) {
	var request = Modeling.Battery{}
	var reader = bytes.NewBufferString(string(*message))
	var jsonDecoderError = json.NewDecoder(reader).Decode(&request)

	if jsonDecoderError != nil {
		return Modeling.Battery{}, jsonDecoderError
	}

	return request, nil
}

// DeserializeScoreMessage
// Распознает структуру тела запроса как сообщение
// о статистике клиента (игрока)
// Возвращает структуру десериазированного запроса
// и состояние о выполненной операции
func DeserializeScoreMessage(message *[]byte) (Modeling.Score, error) {
	var request = Modeling.Score{}
	var reader = bytes.NewBufferString(string(*message))
	var jsonDecoderError = json.NewDecoder(reader).Decode(&request)

	if jsonDecoderError != nil {
		return Modeling.Score{}, jsonDecoderError
	}

	return request, nil
}

// DeserializeInitMessage
// Распознает структуру тела запроса как сообщение
// о начале игры.
// Возвращает структуру десериазированного запроса
// и состояние о выполненной операции
func DeserializeInitMessage(message *[]byte) (Modeling.Init, error) {
	var request = Modeling.Init{}
	var reader = bytes.NewBufferString(string(*message))
	var jsonDecoderError = json.NewDecoder(reader).Decode(&request)

	if jsonDecoderError != nil {
		return Modeling.Init{}, jsonDecoderError
	}

	return request, nil
}
