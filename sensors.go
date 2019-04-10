package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"log"
)

const (
	sensorButtonPin        = 16
	echoLocatorPinReceiver = 23
	echoLocatorPinSender   = 24
)

func setUpGPIO() func() {
	if err := rpio.Open(); err != nil {
		log.Fatal("cant open gpio: ", err)
	}
	log.Println("GPIO connected")

	go func() {
		log.Printf("read touch button state...")
		states := map[bool]string{
			false: "off",
			true:  "on",
		}
		lastState := false
		sensor := NewTTP223B(sensorButtonPin)
		sensor.Loop(func(state bool) {
			if state != lastState {
				println(state)
				sendToAllWsConnections(states[state])
			}
			lastState = state
		})
	}()

	go func() {
		log.Printf("read echo locator state...")
		sensor, err := NewHCSR04(echoLocatorPinReceiver, echoLocatorPinSender)
		if err != nil {
			panic(err.Error())
		}
		sensor.Loop(func(distance float32) {
			res := fmt.Sprintf("%v centimeters", distance)
			println(res)
			sendToAllWsConnections(res)
		})
	}()

	return func() {
		rpio.Close()
	}
}
