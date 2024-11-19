package game

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"sync"
	"time"
	"wsst/cmd/client"
)

var (
	// goroutineMutex
	// Управляет состоянием отправки сообщений клиенту
	goroutineMutex sync.Mutex
)

// Click
// Регистрирует "клик" от клиента,
// если "батарея" не разряжена.
// Если батарея разряжена, досрочно завершает работу
func Click(state *client.Client) {
	if state.BatteryPtr.Charge == 0 {
		message := "Wait!"
		messageBytes := bytes.NewBufferString(message)
		_ = state.ConnectPtr.WriteMessage(websocket.TextMessage, messageBytes.Bytes())
		return
	}
	state.InitPtr.Clicks += 1
	state.BatteryPtr.Charge -= 10

	messageBytes, _ := json.Marshal(state.BatteryPtr)

	goroutineMutex.Lock()
	_ = state.ConnectPtr.WriteMessage(websocket.TextMessage, messageBytes)
	goroutineMutex.Unlock()

}

// UpdateBattery
// Запускает счетчик обновлений состояния
// батареи клиента. Возвращает сообщение о
// состоянии батареи клиенту каждые 2 секунды.
func UpdateBattery(client *client.Client) {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		if client.BatteryPtr.Charge < 100 {
			client.BatteryPtr.Charge += 2
			client.InitPtr.BatteryCharge += 2
		}
		status, _ := json.Marshal(*client.BatteryPtr)

		goroutineMutex.Lock()
		_ = client.ConnectPtr.WriteMessage(websocket.TextMessage, status)
		goroutineMutex.Unlock()
	}
}

// SetInitMessage
// Авторизует пользователя (клиента) на сервере.
// Создает и отправляет сообщение
// инициализации "init"
func SetInitMessage(clientPtr *client.Client) {
	clientPtr.InitPtr.Type = "init"
	clientPtr.InitPtr.Team = getTeam()
	clientPtr.InitPtr.BatteryCharge = 100
	clientPtr.InitPtr.BatteryCapacity = 100

	clientPtr.BatteryPtr.Type = "battery_status"
	clientPtr.BatteryPtr.Charge = 100
	clientPtr.BatteryPtr.Capacity = 100

	initBytes, _ := json.Marshal(clientPtr.InitPtr)

	goroutineMutex.Lock()
	_ = clientPtr.ConnectPtr.WriteMessage(websocket.TextMessage, initBytes)
	goroutineMutex.Unlock()
}

// getTeam
// Около-случайно определяет
// команду клиента
func getTeam() string {
	if time.Now().UnixMilli()%2 == 0 {
		return "white"
	} else {
		return "black"
	}
}
