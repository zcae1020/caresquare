package module

import "myconstant"

func printTime(hour int, minute int) {

}

func PrintDay(actions []int) {
	time := 0
	actionIdx := 0

	for time <= myconstant.MINUTE_OF_DAY {
		time = nextTime(time)
		printAction(actions[actionIdx])
		actionIdx++
	}
}

func printAction(action int) {

}
