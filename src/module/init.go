package module

import (
	"myconstant"
)

func InitStatus(status []int) {
	for i := 0; i < myconstant.MAX_STATUS_TYPE; i++ {
		status[i] = 100
	}
}

func StartDay() {
	var myAction []int
	time := 0

	for time <= myconstant.MINUTE_OF_DAY {
		time = nextTime(time)
		myAction = append(myAction, action())
	}

	PrintDay(myAction)
}

func PrintDay(actions []int) {
	time := 0
	actionIdx := 0

	for time <= myconstant.MINUTE_OF_DAY {
		time = nextTime(time)
		printAction(actionIdx)
	}
}
