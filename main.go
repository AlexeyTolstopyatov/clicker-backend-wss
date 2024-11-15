package main

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
	"wsst/Handling"
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
	Handling.SetClientTeam()
	init, marshalError := json.Marshal(Handling.SetInitMessage(Handling.GetClientTeam()))

	if marshalError != nil {
		log.Printf("[x]: %s", marshalError.Error())
	}

	_ = context.WriteMessage(websocket.TextMessage, init)

	for {
		var _, _, err = context.ReadMessage()
		if err != nil {
			log.Printf("[!]: %s\n", err)
			return
		}

		go Handling.Set(context)
	}
}
