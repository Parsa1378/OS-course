package structs

type Process struct {
	PID            int
	AT             int
	BT             int16
	WaitTime       int16
	TurnaroundTime int16
}

type Chart struct {
	AlgorithmName     string
	AvgWaitingTime    int16
	AvgTurnaroundTime int16
	Processes         []Process
}

func newProcess(pid int, at int, bt int16) Process {
	return Process{PID: pid, AT: at, BT: bt}
}

// sort interface implementation for AT
type AtSorter []Process

func (ats AtSorter) Len() int           { return ats.Len() }
func (ats AtSorter) Swap(i, j int)      { ats[i], ats[j] = ats[j], ats[i] }
func (ats AtSorter) Less(i, j int) bool { return ats[i].AT < ats[j].AT }
