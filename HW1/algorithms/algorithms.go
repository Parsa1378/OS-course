package algorithms

import (
	"OS-course/HW1/structs"
	"sort"
	"sync"
)

func FCFS(processes []*structs.Process) {
	//sorting based on AT
	sort.Sort(structs.AtSorter(processes))
	completionTime := 0
	var wg sync.WaitGroup
	for _, process := range processes {
		wg.Add(1)
		go func(p *structs.Process) {
			defer wg.Done()
			if completionTime < p.AT {
				completionTime = p.AT
			}
			p.CompletionTime = completionTime + p.BT
			p.TurnaroundTime = float32(p.CompletionTime - p.AT)
			p.WaitingTime = p.TurnaroundTime - float32(p.BT)
			completionTime = p.CompletionTime
		}(process)
	}
	wg.Wait()
}

func RR(processes []*structs.Process) {
	totalTime := 0
	// time := 0
	// n := len(processes) - 1
	// queue := make([]structs.Process, n)
	// remainingProcesses := len(processes)
	timeQuantum := 2
	// sort.Sort(structs.AtSorter(processes))
	for i := 0; i < len(processes); i++ {
		processes[i].RemainingTime = processes[i].BT
		totalTime += processes[i].BT
		// queue = append(queue, *processes[i])
	}
	// without goroutine
	// for len(queue) > 0 {
	// 	topProcess := queue[0]
	// 	if topProcess.RemainingTime > 0 {
	// 		topProcess.RemainingTime -= timeQuantum
	// 		time += timeQuantum
	// 		if topProcess.RemainingTime <= 0 {
	// 			queue = queue[1:]
	// 			n--
	// 			topProcess.CompletionTime = time
	// 			topProcess.TurnaroundTime = float32(time - topProcess.AT)
	// 			topProcess.WaitingTime = float32(topProcess.TurnaroundTime) - float32(topProcess.BT)
	// 		} else {
	// 			for i := n; i < 0; i-- {
	// 				queue[i-1] = queue[i]
	// 			}
	// 			queue[n] = topProcess
	// 		}
	// 	}
	// }

	queue := make(chan *structs.Process)
	done := make(chan bool)
	go func() {
		for _, process := range processes {
			queue <- process
		}
		done <- true
	}()
	go func() {
		time := 0
		activeProcesses := &structs.Process{}
		for {
			if activeProcesses.RemainingTime > 0 {
				activeProcesses.RemainingTime -= timeQuantum
				time += timeQuantum
				if activeProcesses.RemainingTime <= 0 {
					activeProcesses.CompletionTime = time
					activeProcesses.TurnaroundTime = float32(time - activeProcesses.AT)
					activeProcesses.WaitingTime = float32(activeProcesses.TurnaroundTime) - float32(activeProcesses.BT)
					queue <- activeProcesses
				}
			}
			if activeProcesses.RemainingTime <= 0 {
				if len(queue) > 0 {
					activeProcesses = <-queue
				} else if time >= totalTime {
					done <- true
					return
				} else {
					time++
				}
			}
		}
	}()
	for i := 0; i < len(processes); i++ {
		<-queue
	}
}
