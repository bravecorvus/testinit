package main

import (
	"fmt"
	"time"

	"github.com/roylee0704/gron"
)

func main() {
	c := gron.New()
	c.AddFunc(gron.Every(1*time.Minute), func() {
		fmt.Println("runs every minute.")
	})
	c.Start()
}
