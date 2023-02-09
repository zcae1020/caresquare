package myconstant

type Status int

const MAX_STATUS_TYPE = 7
const TIME_INTERVAL = 30
const MINUTE_OF_DAY = 1440

const (
	Study Status = iota
	Eat
	Toliet
	Sleep
	Play
	Wash
	Job
)
