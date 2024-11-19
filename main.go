package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"wsst/cmd/client"
	"wsst/cmd/game"
)

func main() {
	var application = fiber.New()

	application.Get("/", handleHttpContext)
	application.Use("/ws", trySwitchContext)
	application.Get("/ws", websocket.New(handleWsContext))

	err := application.Listen(":8080")

	if err != nil {
		print(err)
		return
	}
}

// trySwitchContext
// Обновляет схему передачи данных
func trySwitchContext(context *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(context) {
		context.Locals("allowed", true)
		return context.Next()
	}
	return fiber.ErrUpgradeRequired
}

// handleHttpContext
// Обрабатывает http соединение
func handleHttpContext(context *fiber.Ctx) error {
	return context.SendFile("index.html")
}

// handleWsContext
// Обрабатывает ws соединение
func handleWsContext(context *websocket.Conn) {
	name := client.New(context)
	state := client.ActiveClientTable[name]

	game.SetInitMessage(state)

	go game.UpdateBattery(state)
	go game.UpdateScore(state)

	for {
		_, _, err := context.ReadMessage()
		if err != nil {
			log.Errorf("%s\n", err.Error())
			break
		}
		go game.Click(state)
	}

	client.Delete(name)
	log.Warnf("Client %s disconnected\n", name)
}
