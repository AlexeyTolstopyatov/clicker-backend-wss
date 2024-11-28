package client

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
	"math/rand"
	"sync"
	"wsst/model"
)

// Client
// Client's structure :3
type Client struct {
	ConnectPtr *websocket.Conn
	InitPtr    *model.Init
	BatteryPtr *model.Battery
}

var (
	// ActiveClientTable (ACT)
	// Global Active Clients table
	ActiveClientTable = make(map[string]*Client)

	// mutex
	// Main controller (:D!!!) for [ACT]
	mutex sync.Mutex
)

// create
// Creates ACID (Active Client ID)
func create() string {
	acid := make([]byte, 16)
	_, err := rand.Read(acid)
	if err != nil {
		log.Fatalf("Unable to create [ACID]: %v", err)
		log.Warnf("Retry...")
		create()
	}
	acid[8] = acid[8]&^0xc0 | 0x80
	acid[6] = acid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", acid[0:4], acid[4:6], acid[6:8], acid[8:10], acid[10:])
}

// New
// Authenticate new client, based on temporary Active-Client Identification key
// Returns: [ACID] (Active Client ID)
func New(connection *websocket.Conn) string {
	name := create()
	init := model.Init{}
	battery := model.Battery{}

	client := &Client{
		ConnectPtr: connection,
		InitPtr:    &init,
		BatteryPtr: &battery,
	}

	mutex.Lock()
	ActiveClientTable[name] = client
	mutex.Unlock()

	return name
}

// Delete
// Removes client from global ACT (Active Clients Table)
func Delete(acid string) {
	mutex.Lock()
	delete(ActiveClientTable, acid)
	mutex.Unlock()
}
