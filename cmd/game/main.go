package game

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"time"
	"wsst/cmd/client"
)

// Click
// Registers click request
// counts statistics
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
	_ = state.ConnectPtr.WriteMessage(websocket.TextMessage, messageBytes)
}

func UpdateBattery(client *client.Client) {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		if client.BatteryPtr.Charge < 100 {
			client.BatteryPtr.Charge += 2
			client.InitPtr.BatteryCharge += 2
		}
		status, _ := json.Marshal(client.BatteryPtr)
		_ = client.ConnectPtr.WriteMessage(websocket.TextMessage, status)
	}
}

func SetInitMessage(clientPtr *client.Client) {
	clientPtr.InitPtr.Type = "init"
	clientPtr.InitPtr.Team = getTeam()
	clientPtr.InitPtr.BatteryCharge = 100
	clientPtr.InitPtr.BatteryCapacity = 100

	clientPtr.BatteryPtr.Type = "battery_status"
	clientPtr.BatteryPtr.Charge = 100
	clientPtr.BatteryPtr.Capacity = 100

	initBytes, _ := json.Marshal(clientPtr.InitPtr)
	_ = clientPtr.ConnectPtr.WriteMessage(websocket.TextMessage, initBytes)
}

func getTeam() string {
	if time.Now().UnixMilli()%2 == 0 {
		return "white"
	} else {
		return "black"
	}
}
