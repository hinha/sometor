package main

import (
	"fmt"
	"github.com/hinha/sometor/provider/api"
	"github.com/hinha/sometor/provider/command"
	"github.com/hinha/sometor/provider/infrastructure"
	"github.com/hinha/sometor/provider/scheduler"
	"github.com/hinha/sometor/provider/socket"
	"github.com/hinha/sometor/provider/user"
	"github.com/subosito/gotenv"
	"os"
	"time"
)

func init() {
	_ = os.Setenv("TZ", "Asia/Jakarta")
	loc, _ := time.LoadLocation(os.Getenv("TZ"))
	time.Local = loc
}

func main() {
	_ = gotenv.Load()
	cmd := command.Fabricate()

	// Infra
	infra, err := infrastructure.Fabricate()
	if err != nil {
		panic(err)
	}
	defer infra.Close()

	db, err := infra.DB()
	if err != nil {
		panic(err)
	}

	s3Management, err := infra.S3()
	if err != nil {
		panic(err)
	}
	fmt.Println("s3: ", s3Management)

	if err := infra.FabricateCommand(cmd); err != nil {
		panic(err)
	}

	// User
	userFabricate := user.FabricateStream(db)

	// API
	apiEngine := api.Fabricate(9000)
	apiEngine.FabricateCommand(cmd)

	// Socket
	socketEngine := socket.Fabricate(7000)
	socketEngine.FabricateCommand(cmd)

	// Scheduler Cron
	cronJob := scheduler.Fabricate("test")
	cronJob.FabricateCommand(cmd)

	keywordJob := scheduler.FabricateKeyword(userFabricate)
	keywordJob.FabricateSchedule(cronJob)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
