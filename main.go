package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
	"wsst/cmd/client"
	"wsst/cmd/game"
)

func main() {
	var application = fiber.New()

	application.Get("/", handleHttpContext)
	application.Use("/ws", trySwitchContext)
	application.Get("/ws", websocket.New(handleWebSocketContext))

	err := application.Listen(":8080")

	if err != nil {
		print(err)
		return
	}
}

func trySwitchContext(context *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(context) {
		context.Locals("allowed", true)
		return context.Next()
	}
	return fiber.ErrUpgradeRequired
}

func handleHttpContext(context *fiber.Ctx) error {
	return context.SendFile("index.html")
}

func handleWebSocketContext(context *websocket.Conn) {
	name := client.New(context)
	state := client.ActiveClientTable[name]

	go game.SetInitMessage(state)
	go game.UpdateBattery(state)

	for {
		_, _, err := context.ReadMessage()
		if err != nil {
			log.Printf("[!]: %s\n", err.Error())
			break
		}
		go game.Click(state)
	}

	client.Delete(name)
	log.Printf("Client %s disconnected\n", name)
}
