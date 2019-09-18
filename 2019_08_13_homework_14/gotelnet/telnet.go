package gotelnet

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type Telnet struct {
	conn net.Conn
}

func (t Telnet) Connect(address string, timeout time.Duration) Telnet {
	dialer, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Fatalln(err)
	}

	return Telnet{dialer}
}

func (t *Telnet) Start(ctx context.Context) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(ctx context.Context) {
		t.receiver(ctx)
		wg.Done()
	}(ctx)

	wg.Add(1)
	go func(ctx context.Context) {
		t.sender(ctx)
		wg.Done()
	}(ctx)

	wg.Wait()
	errorCheck(t.conn.Close)
}

func (t *Telnet) sender(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	bridge := make(chan string, 1)
	defer errorCheck(scanner.Err)

	go func(bridge chan<- string) {
		for scanner.Scan() {
			bridge <- scanner.Text()
		}
	}(bridge)

	for {
		select {
		case <-ctx.Done():
			close(bridge)
			log.Println("close sender")
			return
		case msg := <-bridge:
			log.Println("Send: " + msg)
			if _, err := t.conn.Write([]byte(msg + "\n")); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func (t *Telnet) receiver(ctx context.Context) {
	scanner := bufio.NewScanner(t.conn)
	bridge := make(chan string, 1)
	defer errorCheck(scanner.Err)

	go func(bridge chan<- string) {
		for scanner.Scan() {
			bridge <- scanner.Text()
		}
	}(bridge)

	for {
		select {
		case <-ctx.Done():
			close(bridge)
			log.Println("close receiver")
			return
		case msg := <-bridge:
			log.Println("Receive: " + msg)
		}
	}
}

func errorCheck(fn func() error) {
	if err := fn(); err != nil {
		log.Println(err)
	}
}
