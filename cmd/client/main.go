package client

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
	"math/rand"
	"sync"
	"wsst/shared"
)

type Client struct {
	ConnectPtr *websocket.Conn
	InitPtr    *shared.Init
	BatteryPtr *shared.Battery
}

// ActiveClientTable (ACT)
// contains all online users
var (
	ActiveClientTable = make(map[string]*Client)
	Mutex             sync.Mutex
)

// create
// Generates GUID for Client
func create() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// New
// Inserts Client (GUID) into ACT (Active clients table)
func New(connection *websocket.Conn) string {
	name := create()
	init := shared.Init{}
	battery := shared.Battery{}

	client := &Client{
		ConnectPtr: connection,
		InitPtr:    &init,
		BatteryPtr: &battery,
	}

	Mutex.Lock()
	ActiveClientTable[name] = client
	Mutex.Unlock()

	return name
}

// Delete
// Removes Client (GUID) from ACT (Active Client Table)
func Delete(guid string) {
	Mutex.Lock()
	delete(ActiveClientTable, guid)
	Mutex.Unlock()
}
