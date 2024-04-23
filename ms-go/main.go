package main

import (
	"ms-go/app/consumers"
	_ "ms-go/db"
	"ms-go/router"
	"time"
)

func main() {
	go router.Run()
	time.Sleep(15 * time.Second)
	consumers.RailsConsumer()
}
