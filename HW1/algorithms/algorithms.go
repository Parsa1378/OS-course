package algorithms

import (
	"fmt"
	"sort"

	"github.com/Parsa1378/OS-course/structs"
)

func FCFS(processes []structs.Process) {
	//sorting based on AT
	l := len(processes)
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].AT < processes[j].AT
	})
	for i := 0; i < l; i++ {
		if i == 0 {
			processes[i].CompletionTime = processes[i].BT + processes[i].AT
		} else {
			processes[i].CompletionTime = processes[i-1].CompletionTime + processes[i].BT
		}
		processes[i].TurnaroundTime = processes[i].CompletionTime - processes[i].AT
		processes[i].WaitingTime = processes[i].TurnaroundTime - processes[i].BT
	}
}

func RR(processes []structs.Process) {
	n := len(processes)
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].AT < processes[j].AT
	})

	for i := 0; i < n; i++ {
		processes[i].RemainingTime = processes[i].BT
	}
	lastAT := processes[n-1].AT
	var queue []*structs.Process
	timeQuantum := 2
	time := 0
	queue = append(queue, &processes[0])
	lastIn := 0
	for len(queue) > 0 || time < lastAT {
		if queue[0].RemainingTime > timeQuantum {
			queue[0].RemainingTime -= timeQuantum
			time += timeQuantum
			if queue[0].RemainingTime <= 0 {
				queue[0].RemainingTime = 0
				queue[0].CompletionTime = queue[0].BT
				queue[0].TurnaroundTime = queue[0].CompletionTime - queue[0].AT
				queue[0].WaitingTime = queue[0].TurnaroundTime - queue[0].BT
				queue = queue[1:]
			}

			for i := lastIn + 1; i < n; i++ {
				if processes[i].AT <= time {
					queue = append(queue, &processes[i])
					lastIn = i
				}
			}
			top := queue[0]
			for j := 0; j < len(queue)-1; j++ {
				queue[j] = queue[j+1]
			}
			queue[len(queue)-1] = top
		} else {
			queue[0].CompletionTime = queue[0].RemainingTime + time
			time += queue[0].RemainingTime
			queue[0].RemainingTime = 0
			queue[0].TurnaroundTime = queue[0].CompletionTime - queue[0].AT
			queue[0].WaitingTime = queue[0].TurnaroundTime - queue[0].BT
			queue = queue[1:]
			for i := lastIn + 1; i < n; i++ {
				if processes[i].AT <= time {
					queue = append(queue, &processes[i])
					lastIn = i
				}
			}
		}
	}
	fmt.Println(processes)
}
