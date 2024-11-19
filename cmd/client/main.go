package client

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
	"math/rand"
	"sync"
	"wsst/shared"
)

// Client
// Структура данных клиента
type Client struct {
	ConnectPtr *websocket.Conn
	InitPtr    *shared.Init
	BatteryPtr *shared.Battery
}

var (
	// ActiveClientTable (ACT)
	// Содержит всех online клиентов
	ActiveClientTable = make(map[string]*Client)

	// mutex
	// Управляет состоянием в функциях
	// создания и удаления
	mutex sync.Mutex
)

// create
// Создает ACID (Active Client ID)
// для клиента
func create() string {
	acid := make([]byte, 16)
	_, err := rand.Read(acid)
	if err != nil {
		log.Fatalf("Не удалось создать ACID: %v", err)
	}
	acid[8] = acid[8]&^0xc0 | 0x80
	acid[6] = acid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", acid[0:4], acid[4:6], acid[6:8], acid[8:10], acid[10:])
}

// New
// Регистрирует клиента (Client) в ACT таблице (Active clients table)
// Возвращает ACID зарегистрированного пользователя
func New(connection *websocket.Conn) string {
	name := create()
	init := shared.Init{}
	battery := shared.Battery{}

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
// Удаляет клиента (Client) из ACT (Active Client Table)
// на основе переданного ACID
func Delete(acid string) {
	mutex.Lock()
	delete(ActiveClientTable, acid)
	mutex.Unlock()
}
