package myaccessory

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type Humidity struct {
	*accessory.Accessory
	HumiditySensor *service.HumiditySensor
}

func NewHumiditySensor(info accessory.Info, humi, min, max, steps float64) *Humidity {
	acc := Humidity{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.HumiditySensor.CurrentRelativeHumidity.SetValue(humi)
	acc.HumiditySensor.CurrentRelativeHumidity.SetMinValue(min)
	acc.HumiditySensor.CurrentRelativeHumidity.SetMaxValue(max)
	acc.HumiditySensor.CurrentRelativeHumidity.SetStepValue(steps)

	acc.AddService(acc.HumiditySensor.Service)

	return &acc
}
