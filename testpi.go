package main

import (
	// "github.com/robfig/cron"
	"fmt"
	"os"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func VibOn() {
	// fmt.Println("VibOn")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	Input1 := rpio.Pin(5)
	Input1.Output()
	Input1.High()
	Input2 := rpio.Pin(6)
	Input2.Output()
	Input2.Low()
	Enable := rpio.Pin(17)
	Enable.Output()
	Enable.High()
}

//Sends the signal to turn off the bed vibrator by sending a Low (false) signal to GPIO 17
func VibOff() {
	// fmt.Println("VibOff")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	Input1 := rpio.Pin(5)
	Input1.Output()
	Input1.High()
	Input2 := rpio.Pin(6)
	Input2.Output()
	Input2.Low()
	Enable := rpio.Pin(17)
	Enable.Output()
	Enable.Low()
}

func main() {
	VibOn()
	time.Sleep(time.Millisecond * 50)
	VibOff()
}
