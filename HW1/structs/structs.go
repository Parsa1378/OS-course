package structs

type Process struct {
	PID            string
	AT             int
	BT             int
	RemainingTime  int
	CompletionTime int
	WaitingTime    float32
	TurnaroundTime float32
}

func newProcess(pid string, at int, bt int) Process {
	return Process{PID: pid, AT: at, BT: bt}
}
