package app

import (
	"errors"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"wsst/internal/client"
	"wsst/internal/game"
)

var (
	EmptyEnvironmentError = errors.New("empty environment file")
)

// Run
// Preparing and running general server's instance
// bases on StartupEventArgs. You must to recorgize
// startup arguments for calling it.
func Run() {
	var application = fiber.New()

	application.Get("/", handleHttpContext)
	application.Use("/ws", trySwitchContext)
	application.Get("/ws", websocket.New(handleWsContext))

	err := application.Listen(
		StartupArgs.Address + ":" + StartupArgs.Port)

	if err != nil {
		print(err)
		return
	}
}

// Shutdown
// I see it in template of GO backend project
// Usually it must correctly shut down
// server's internals (ex. Postgres services closing...)
func Shutdown() {

}

// trySwitchContext
// Updates connection schema (changes protocol)
func trySwitchContext(context *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(context) {
		context.Locals("allowed", true)
		return context.Next()
	}
	return fiber.ErrUpgradeRequired
}

// handleHttpContext
// Handles HTTP connection.
// Sends main page with information message
// (http://address:port/)
func handleHttpContext(context *fiber.Ctx) error {
	return context.SendFile("index.html")
}

// handleWsContext
// Handles WS connection.
// Starts the game package. Here is the main game-logic body
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
