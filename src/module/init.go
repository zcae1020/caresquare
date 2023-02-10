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
	var myStatus [myconstant.MAX_STATUS_TYPE]int
	time := 0

	for time <= myconstant.MINUTE_OF_DAY {
		time = nextTime(myStatus, time)
		myAction = append(myAction, action(myStatus))
	}

	PrintDay(myAction)
}
