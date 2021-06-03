package main

import (
	"fmt"
	"golang.org/x/exp/io/i2c"
	"github.com/maciej/bme280"
)

func main() {
	i2cAddr := 0x76
	device, _ := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, i2cAddr)
	driver := bme280.New(device)
	_ = driver.InitWith(bme280.ModeForced, bme280.Settings{
		Filter:                  bme280.FilterOff,
		Standby:                 bme280.StandByTime1000ms,
		PressureOversampling:    bme280.Oversampling16x,
		TemperatureOversampling: bme280.Oversampling16x,
		HumidityOversampling:    bme280.Oversampling16x,
	})

	response, _ := driver.Read()
	fmt.Println(response)
}
