package main

import (
	"golang.org/x/exp/io/i2c"
	"github.com/maciej/bme280"
	"github.com/gin-gonic/gin"
)

type SensingData struct {
	Temperature	float64	`json:"temperature"`
	Pressure	float64 `json:"pressure"`
	Humidity	float64	`json:"humidity"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	d, _ := readDevice()
	r.GET("/sensing", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"temperature": d.Temperature,
			"pressure": d.Pressure,
			"humidity": d.Humidity,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func readDevice() (SensingData, error) {
	i2cAddr := 0x76
	device, _ := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, i2cAddr)
	driver := bme280.New(device)
	err := driver.InitWith(bme280.ModeForced, bme280.Settings{
		Filter:                  bme280.FilterOff,
		Standby:                 bme280.StandByTime1000ms,
		PressureOversampling:    bme280.Oversampling16x,
		TemperatureOversampling: bme280.Oversampling16x,
		HumidityOversampling:    bme280.Oversampling16x,
	})
	if err != nil {
		return SensingData{}, err
	}

	r, err := driver.Read()
	if err != nil {
		return SensingData{}, err
	}
	return SensingData{r.Temperature, r.Pressure, r.Humidity}, nil
}
