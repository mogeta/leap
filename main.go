//
// start -> update ->end
package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/leap"
)

const (
	StateStart  = "start"
	StateUpdate = "update"
	StateStop   = "stop"

	TypeSwipe  = "swipe"
	TypeCircle = "circle"
)

func main() {
	gbot := gobot.NewGobot()

	leapMotionAdaptor := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	l := leap.NewLeapMotionDriver(leapMotionAdaptor, "leap")

	work := func() {
		gobot.On(l.Event("message"), func(data interface{}) {
			if len(data.(leap.Frame).Gestures) > 0 {

				if data.(leap.Frame).Gestures[0].State == StateStop {
					fmt.Println(data.(leap.Frame).Gestures[0].Type)
				}

			}

		})
	}

	robot := gobot.NewRobot("leapBot",
		[]gobot.Connection{leapMotionAdaptor},
		[]gobot.Device{l},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
