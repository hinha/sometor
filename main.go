package main

import (
	"github.com/hinha/sometor/provider/api"
	"github.com/hinha/sometor/provider/command"
	"github.com/hinha/sometor/provider/infrastructure"
	"github.com/hinha/sometor/provider/scheduler"
	"github.com/hinha/sometor/provider/socket"
	"github.com/hinha/sometor/provider/socket_stream"
	"github.com/hinha/sometor/provider/socmed"
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

	celery := infra.Celery()

	s3Management, err := infra.S3()
	if err != nil {
		panic(err)
	}

	if err := infra.FabricateCommand(cmd); err != nil {
		panic(err)
	}

	// User
	userFabricate := user.FabricateStream(db)
	keywordStream := user.FabricateStreamKeyword(db)

	// API
	apiEngine := api.Fabricate(9000)
	apiEngine.FabricateCommand(cmd)

	keywordStreamAPI := socmed.FabricateKeyword(keywordStream)
	keywordStreamAPI.FabricateAPI(apiEngine)

	// Socket
	socketEngine := socket.Fabricate(7000)
	socketEngine.FabricateCommand(cmd)

	// Socket Stream Twitter, Instagram and facebook
	twitterSocket := socket_stream.Fabricate(userFabricate)
	twitterSocket.FabricateSocket(socketEngine)

	// Scheduler Cron Local
	cronJobLocal := scheduler.FabricateLocal("cron_local")
	cronJobLocal.FabricateCommand(cmd)
	cronJobServer := scheduler.FabricateServer("cron_server")
	cronJobServer.FabricateCommand(cmd)

	keywordJobLocal := scheduler.FabricateKeyword(userFabricate, celery, s3Management)
	keywordJobLocal.FabricateSchedule(cronJobLocal)

	keywordJobServer := scheduler.FabricateKeywordServer(userFabricate, celery, s3Management)
	keywordJobServer.FabricateSchedule(cronJobServer)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
