package module

import "myconstant"

func printTime(hour int, minute int) {

}

func printDay(status [myconstant.MAX_STATUS_TYPE]int, actions []int) {
	time := 0
	actionIdx := 0

	for time <= myconstant.MINUTE_OF_DAY {
		time += myconstant.TIME_INTERVAL
		printTime(getCurrentTime(time))
		printAction(actions[actionIdx])
		actionIdx++
	}

	printStatus(status)
}

func printAction(action int) {

}

func printStatus(status [myconstant.MAX_STATUS_TYPE]int) {

}