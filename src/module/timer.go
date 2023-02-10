package module

import "myconstant"

func getCurrentTime(time int) (int, int) {

}

func nextTime(time int) int {
	for i := 0; i <= myconstant.TIME_INTERVAL; i += 10 {
		printTime(getCurrentTime(time + i))
	}
	return time + myconstant.TIME_INTERVAL
}
