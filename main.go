//
// start -> update ->end
package leap

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

type Action interface {
	Swipe()
	Circle()
}

type LeapMotion struct {
	act Action
}

var messages = make(chan string)
var isWork = false
var gesture = ""

func New(a Action) *LeapMotion {
	return &LeapMotion{a}
}

func (lp *LeapMotion) Start() {
	gbot := gobot.NewGobot()

	leapMotionAdaptor := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	l := leap.NewLeapMotionDriver(leapMotionAdaptor, "leap")

	work := func() {
		gobot.On(l.Event("message"), func(data interface{}) {
			//fmt.Println(data.(leap.Frame).ID)
			if len(data.(leap.Frame).Gestures) > 0 {
				isWork = true
				if data.(leap.Frame).Gestures[0].State == StateStop {
					//fmt.Println(data.(leap.Frame).Gestures[0].Type)
					//messages <- data.(leap.Frame).Gestures[0].Type
					gesture = data.(leap.Frame).Gestures[0].Type
				}
			} else if isWork {
				isWork = false
				messages <- gesture
				fmt.Println("end")
			}

		})
	}

	robot := gobot.NewRobot("leapBot",
		[]gobot.Connection{leapMotionAdaptor},
		[]gobot.Device{l},
		work,
	)

	go lp.allocate()
	gbot.AddRobot(robot)
	gbot.Start()
}

func (lp *LeapMotion) allocate() {
	for {
		msg := <-messages

		switch msg {
		case TypeSwipe:
			lp.act.Swipe()
			break
		case TypeCircle:
			lp.act.Circle()
			break
		}
		fmt.Println(msg)
	}
}
