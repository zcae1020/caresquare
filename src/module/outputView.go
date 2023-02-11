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
	for i := 0; i < myconstant.MAX_STATUS_TYPE; i++ {
		fmt.Printf("%s: %d ", statusToString(i), status[i])
	}
	fmt.Println()

	fmt.Print("보충 필요: ")
	for _, value := range diagnose(status) {
		fmt.Printf("%s ", statusToString(value))
	}
	fmt.Println()
}

func statusToString(status int) string {
	switch myconstant.Status(status) {
	case myconstant.Study:
		return "성장"
	case myconstant.Eat:
		return "식사"
	case myconstant.Toliet:
		return "화장실"
	case myconstant.Sleep:
		return "숙면"
	case myconstant.Play:
		return "놀이"
	case myconstant.Wash:
		return "씻기"
	case myconstant.Job:
		return "취준"
	}
	return ""
}
