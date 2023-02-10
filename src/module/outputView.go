package module

import (
	"fmt"
	"myconstant"
)

func printTime(hour int, minute int) {
	fmt.Printf("현재 시간: %d시, %d분\n", hour, minute)
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
	switch myconstant.Status(action) {
	case myconstant.Study:
		fmt.Println("프로그래밍 공부중...")
	case myconstant.Eat:
		fmt.Println("치킨 먹는중...")
	case myconstant.Toliet:
		fmt.Println("화장실입니다...")
	case myconstant.Sleep:
		fmt.Println("자는중...")
	case myconstant.Play:
		fmt.Println("유튜브 보는중...")
	case myconstant.Wash:
		fmt.Println("씻는중...")
	case myconstant.Job:
		fmt.Println("자소서 쓰는중...")
	}
}

func printStatus(status [myconstant.MAX_STATUS_TYPE]int) {

}
