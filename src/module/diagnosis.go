package module

import "myconstant"

const THRESHOLD = 20

func diagnose(status [myconstant.MAX_STATUS_TYPE]int) (leakList []int) {
	for i := 0; i < myconstant.MAX_STATUS_TYPE; i++ {
		if status[i] < THRESHOLD {
			leakList = append(leakList, i)
		}
	}

	return
}
