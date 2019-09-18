package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IamStubborN/otus-golang/2019_08_13_homework_14/gotelnet"
	"github.com/spf13/pflag"
)

var address *string
var timeout *time.Duration

func init() {
	address = pflag.StringP("address", "a", "0.0.0.0:3302", "gotelnet -a 0.0.0.0:3302")
	timeout = pflag.DurationP("timeout", "t", 30*time.Second, "gotelnet -t 30s")
	pflag.Parse()
}

func main() {
	client := gotelnet.Telnet{}.Connect(*address, *timeout)
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)

	go gracefulShutdown(cancel)
	client.Start(ctx)
}

func gracefulShutdown(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	cancel()
}
