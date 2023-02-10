package module

import "myconstant"

const MINUTE_OF_HOUR = 60

func getCurrentTime(time int) (int, int) {
	return time / MINUTE_OF_HOUR, time % MINUTE_OF_HOUR
}

func changeStatusByTime() {

}

func nextTime(status [myconstant.MAX_STATUS_TYPE]int, time int) int {
	for i := 0; i <= myconstant.TIME_INTERVAL; i += 10 {
		printTime(getCurrentTime(time + i))
		changeStatusByTime()
		printStatus(status)
	}
	return time + myconstant.TIME_INTERVAL
}
