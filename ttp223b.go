package main

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type TTP223B struct {
	pin rpio.Pin
}

func (s *TTP223B) Loop(callback func(state bool)) {
	s.pin.Input()
	for {
		val := s.pin.Read()
		if val == rpio.High {
			s.pin.Write(rpio.High)
			callback(true)
		} else {
			s.pin.Write(rpio.Low)
			callback(false)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// NewTTP223B creates instance for sensor TTP223B
// https://arduino.ua/prod1151-sensornii-datchik-knopka
func NewTTP223B(in int) *TTP223B {
	return &TTP223B{
		pin: rpio.Pin(in),
	}
}
