package main

import (
	"testing"
	"machine"
	"tinygo.org/x/drivers/bme280"
)

func TestCreateNewConnection(t *testing.T) {
	machine.I2C0.Configure(machine.I2CConfig{})
	sensor := bme280.New(machine.I2C0)
	sensor.Configure()

	connected := sensor.Connected()
	if !connected {
		t.Errorf("BME280 not detected")
	}
}
