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

var messages = make(chan string)

func main() {
	gbot := gobot.NewGobot()

	leapMotionAdaptor := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	l := leap.NewLeapMotionDriver(leapMotionAdaptor, "leap")

	work := func() {
		gobot.On(l.Event("message"), func(data interface{}) {
			fmt.Println(data.(leap.Frame).ID)
			if len(data.(leap.Frame).Gestures) > 0 {

				if data.(leap.Frame).Gestures[0].State == StateStop {
					//fmt.Println(data.(leap.Frame).Gestures[0].Type)
					messages <- data.(leap.Frame).Gestures[0].Type
				}
			}

		})
	}

	robot := gobot.NewRobot("leapBot",
		[]gobot.Connection{leapMotionAdaptor},
		[]gobot.Device{l},
		work,
	)

	go allocate()
	gbot.AddRobot(robot)
	gbot.Start()
}

func allocate() {
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
