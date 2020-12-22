package main

import (
	"github.com/hinha/sometor/provider/api"
	"github.com/hinha/sometor/provider/command"
	"github.com/hinha/sometor/provider/infrastructure"
	"os"
	"time"
)

func init() {
	_ = os.Setenv("TZ", "Asia/Jakarta")
	loc, _ := time.LoadLocation(os.Getenv("TZ"))
	time.Local = loc
}

func main() {
	cmd := command.Fabricate()

	// Infra
	infra, err := infrastructure.Fabricate()
	if err != nil {
		panic(err)
	}
	defer infra.Close()

	if err := infra.FabricateCommand(cmd); err != nil {
		panic(err)
	}

	// API
	apiEngine := api.Fabricate(9000)
	apiEngine.FabricateCommand(cmd)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
