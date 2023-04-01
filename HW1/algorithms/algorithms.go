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
