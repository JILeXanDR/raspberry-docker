package main

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

const HardStop = 1000000

type HCSR04 struct {
	echoPin rpio.Pin
	pingPin rpio.Pin
}

func (sensor *HCSR04) MeasureDistanceSm() float32 {

	sensor.echoPin.Output()
	sensor.pingPin.Output()

	sensor.echoPin.Low()
	sensor.pingPin.Low()

	sensor.echoPin.Input()

	strobeZero := 0
	strobeOne := 0

	// strobe
	delayUs(200)
	sensor.pingPin.High()

	delayUs(15)
	sensor.pingPin.Low()

	// wait until strobe back

	for strobeZero = 0; strobeZero < HardStop && sensor.echoPin.Read() != rpio.High; strobeZero++ {
	}
	start := time.Now()
	for strobeOne = 0; strobeOne < HardStop && sensor.echoPin.Read() != rpio.Low; strobeOne++ {
		delayUs(1)
	}
	end := time.Now()

	return float32(end.UnixNano()-start.UnixNano()) / (58.0 * 1000)
}
func (sensor *HCSR04) Loop(callback func(distance float32)) {
	for {
		callback(sensor.MeasureDistanceSm())
		time.Sleep(500 * time.Millisecond)
	}
}

func delayUs(ms int) {
	time.Sleep(time.Duration(ms) * time.Microsecond)
}

// NewHCSR04 creates instance for sensor HC-SR04
// https://arduino.ua/prod182-yltrazvykovoi-datchik-rasstoyaniya-hc-sr04
func NewHCSR04(echo int, trigger int) (*HCSR04, error) {
	return &HCSR04{
		echoPin: rpio.Pin(echo),
		pingPin: rpio.Pin(trigger),
	}, nil
}
