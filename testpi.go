package main

import (
	"fmt"
	"time"

	"github.com/roylee0704/gron"
)

func main() {
	// t := time.Now()
	// currenttime := t.Format("15:04")
	c := gron.New()
	//Run the following once a minute
	//Check all 4 alarms to see if the current time matches any configurations
	c.AddFunc(gron.Every(1*time.Minute), func() {
		fmt.Println("YOLO")
	})
	c.Start()
}
