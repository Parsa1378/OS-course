package algorithms

import (
	"OS-course/HW1/structs"
	"sort"
)

func FCFS(processes []structs.Process) {
	//sorting based on AT
	sort.Sort(structs.AtSorter(processes))
}
