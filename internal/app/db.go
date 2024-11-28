package app

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

// Here is initialization using "db.json"

type PostgresProfile struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

// UseDatabaseConfig
// Searches db.json inside the application domain.
// Includes to configuration if .json profile exists
// Else ~Database profile 'll be filled by application defaults~
//
// username="admin"
// password="admin"
func UseDatabaseConfig() {
	// "postgres://user:password@host:port/dbname?sslmode=disable"
}

// seekDB
// Returns 'true' if .json profile file
// exists
func seekDB() bool {
	if _, err := os.Stat("db.json"); errors.Is(err, os.ErrNotExist) {
		return false
	}
	segs, _ := os.Open("db.json")
	reader := bufio.NewReader(segs)

	err := json.NewDecoder(reader).Decode(&StartupArgs)
	if err != nil {
		log.Warnf("Database Profile opening error:"+
			"%s", err.Error())
		return false
	}

	return true
}
