package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/realHoangHai/go-assmin/config"
	"github.com/realHoangHai/go-assmin/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

// VERSION Usage: go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "1.0"

var (
	file string
)

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("      -c {config file}")
	fmt.Println("      -h (show help info)")
}

func parse() bool {
	flag.StringVar(&file, "c", "configs/sig.toml", "config file")
	help := flag.Bool("h", false, "help info")
	flag.Parse()

	if !config.Load(file) {
		return false
	}

	if *help {
		showHelp()
		return false
	}
	return true
}

// @title assmin
// @version 1.0
// @description Go Assmin.
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
// @license.name MIT
// @license.url http://github.com/realHoangHai/go-assmin/blob/main/LICENSE
// @contact.name realHoangHai
// @contact.email aflyingpenguin2lth@gmail.com
func main() {
	if !parse() {
		showHelp()
		os.Exit(-1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)

	log.Init(config.C.Log.Level)
	// init	server
	server, err := initializeServer(ctx)
	if err != nil {
		log.Fatalf("initialize server: %v", err)
	}
	server.Start()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Exiting...")
}
