package app

type StartupEventArgs struct {
	Address string `json:"address"`
	Port    string `json:"port"`
	Debug   bool   `json:"debug"`
}
