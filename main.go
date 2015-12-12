package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/ttyACM0")
	sensor := gpio.NewAnalogSensorDriver(firmataAdaptor, "sensor", "0")

	work := func() {
		gobot.On(sensor.Event("data"), func(data interface{}) {
			var distance int

			if data.(int) != 0 {
				distance = (6787 / (data.(int) - 3)) - 4
			}

			if distance < 30 {
				fmt.Println("ちかすぎるぷり〜〜〜!!!")
			} else {
				fmt.Println("とおいぷり〜〜〜")
			}
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
