package structs

type Process struct {
	PID int16
	AT int16
	BT int16
	WaitTime int16
	TurnaroundTime int16
}

type Chart struct {
	AlgorithmName string
	AvgWaitingTime int16
	AvgTurnaroundTime int16
}