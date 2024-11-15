package main

import (
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
	log.Println("Client connected")
	// waiting for initialize message
	for {
		var msgType, msgBytes, err = context.ReadMessage()

		Handling.FixRequestTime()

		// Проверка на ошибки
		if err != nil {
			log.Printf("error reading: %s\n", err)
			return
		}

		if err = context.WriteMessage(msgType, msgBytes); err != nil {
			log.Printf("error writing: %s", err)
			return
		}

		err = Handling.HandleMessage(msgBytes)
		if err != nil {
			log.Printf("Client killed. (error: %s)\n", err)
			return
		}
	}
}
