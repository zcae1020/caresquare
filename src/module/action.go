package module

import "myconstant"

const STATE_CHANGE_BY_ACTION = 50

func action(status *[myconstant.MAX_STATUS_TYPE]int) int {
	action := chooseAction()
	status[action] -= STATE_CHANGE_BY_ACTION
	return action
}
