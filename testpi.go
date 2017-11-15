package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func main() {
	// t := time.Now()
	// currenttime := t.Format("15:04")
	c := cron.New()
	//Run the following once a minute
	//Check all 4 alarms to see if the current time matches any configurations
	c.AddFunc("0 * * * * *", func() {
		fmt.Println("YOLO")
	})
	c.Start()
}
