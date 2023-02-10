package module

import "myconstant"

func getCurrentTime(time int) (int, int) {

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
