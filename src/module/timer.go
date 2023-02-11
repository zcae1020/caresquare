package module

import "myconstant"

const MINUTE_OF_HOUR = 60
const STATE_CHANGE_PER_TIME_INTREVAL = 3
const HOUR_OF_DAY = 24
const WAKE_HOUR = 8

func getCurrentTime(time int) (int, int) {
	time += WAKE_HOUR * MINUTE_OF_HOUR
	return (time / MINUTE_OF_HOUR) % 24, time % MINUTE_OF_HOUR
}

func changeStatusByTime(status *[myconstant.MAX_STATUS_TYPE]int) {
	for i := 0; i < myconstant.MAX_STATUS_TYPE; i++ {
		status[i] -= STATE_CHANGE_PER_TIME_INTREVAL
	}
}

func nextTime(status *[myconstant.MAX_STATUS_TYPE]int, time int) int {
	for i := 0; i < myconstant.TIME_INTERVAL; i += 10 {
		printTime(getCurrentTime(time + i))
	}
	changeStatusByTime(status)
	printStatus(*status)
	return time + myconstant.TIME_INTERVAL
}
