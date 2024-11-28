package app

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"strconv"
)

var StartupArgs = StartupEventArgs{}

// seekEnv
// Returns 'true' if .json or .env configuration file
// exists.
func seekEnv() bool {
	if _, err := os.Stat("server.json"); errors.Is(err, os.ErrNotExist) {
		return false
	}
	segs, _ := os.Open("server.json")
	reader := bufio.NewReader(segs)

	err := json.NewDecoder(reader).Decode(&StartupArgs)
	if err != nil {
		log.Warnf("StartupArgs opening error:"+
			"%s", err.Error())
		return false
	}

	return true
}

// seekArgs
// Returns 'true' if server runs without environment,
// but has command line arguments for running.
//
// ./server localhost 8080 debug
func seekArgs() bool {
	if len(os.Args) != 3 {
		// Prepare default parameters
		StartupArgs = StartupEventArgs{
			Address: "localhost",
			Port:    "8080",
			Debug:   false,
		}
	} else {
		env, _ := strconv.ParseBool(os.Args[3])
		StartupArgs = StartupEventArgs{
			Address: os.Args[1],
			Port:    os.Args[2],
			Debug:   env,
		}
	}
	return false
}

// UseServerConfig
// Starts reading startup arguments and scanning
// run's parameters
func UseServerConfig() {
	if seekEnv() {
		log.Infof("environment exists!\n")
	} else {
		seekArgs()
	}

	if StartupArgs.Debug {
		log.Warnf("%s\n", StartupArgs)
	}
}
