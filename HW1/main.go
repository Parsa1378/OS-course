package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Parsa1378/OS-course/algorithms"
	"github.com/Parsa1378/OS-course/structs"
)

func main() {
	processes := readProcesses("./test.txt")
	algorithms.FCFS(processes)
	fmt.Println(result(processes))
}

// for reading from file
func readProcesses(filename string) []structs.Process {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var processes []structs.Process
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		at, _ := strconv.Atoi(line[2])
		bt, _ := strconv.Atoi(line[1])
		newProcess := structs.NewProcess(line[0], at, bt)
		processes = append(processes, newProcess)
	}
	return processes
}

func result(processes []structs.Process) (float32, float32) {
	var avgTurnaroundTime, avgWaitingTime float32
	for _, process := range processes {
		avgWaitingTime += float32(process.WaitingTime)
		avgTurnaroundTime += float32(process.TurnaroundTime)
	}
	avgTurnaroundTime /= float32(len(processes))
	avgWaitingTime /= float32(len(processes))
	return avgWaitingTime, avgTurnaroundTime
}
