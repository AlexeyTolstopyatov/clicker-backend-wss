package game

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"sync"
	"time"
	"wsst/internal/client"
)

var (
	// goroutineMutex
	// Message sending controller
	goroutineMutex sync.Mutex
)

// Click
// Registers "click" from client's request,
// If "battery-charge" not 0%.
// Else aborting iteration
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
// Starts counter (update timer) of battery status
// Returns (sends to client) the JSON message about
// battery status
// Does it every 2 seconds
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
// Sends the "init" message
// "Init" message -- Standard Kit of client's characteristics
// (see /model/Init.go)
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
// Determines the client's team
// (pseudo-randomize)
// If selected milliseconds divided by 2
// with no reminder then client plays for "white" team.
// Else client plays for "black" team
//
// Returns: "white" or "black"
func getTeam() string {
	if time.Now().UnixMilli()%2 == 0 {
		return "white"
	} else {
		return "black"
	}
}
