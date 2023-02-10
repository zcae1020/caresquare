package module

import (
	"fmt"
	"myconstant"
	"strconv"
)

func chooseAction() int {
	var action string

	fmt.Println("행동을 고르세요.\n0. 공부하기 1. 식사하기 2. 화장실가기 3. 숙면하기 4. 놀기 5. 씻기 6. 취업준비하기")
	fmt.Scan(&action)

	actionNum, err := strconv.Atoi(action)

	if err != nil || actionNum >= myconstant.MAX_STATUS_TYPE {
		fmt.Println("올바른 숫자로 입력하세요.")
		return chooseAction()
	}

	return actionNum
}
