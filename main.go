package main

import (
	"wsst/internal/app"
)

func main() {
	app.UseServerConfig()
	app.UseDatabaseConfig()
	app.Run()
}
