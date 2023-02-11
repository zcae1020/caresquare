package module

import (
	"fmt"
	"myconstant"
)

func initStatus(status *[myconstant.MAX_STATUS_TYPE]int) {
	for i := 0; i < myconstant.MAX_STATUS_TYPE; i++ {
		status[i] = 100
	}
}

func StartDay() {
	var myAction []int
	var myStatus [myconstant.MAX_STATUS_TYPE]int
	time := 0

	initStatus(&myStatus)
	fmt.Println("현재 시각 8시, 기상!")

	for time <= myconstant.MINUTE_OF_DAY {
		myAction = append(myAction, action(&myStatus))
		time = nextTime(&myStatus, time)
	}

	printDay(myStatus, myAction)
}
