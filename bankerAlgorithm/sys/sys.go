package mysys

import "fmt"

type Process struct {
	PID        int
	Max        []int
	Allocation []int
	Need       []int
}

type ProcessSequence []Process //多个进程组成的切片 -> 进程序列

func NewResources(resource ...int) []int {
	return resource
}

//创建进程，初始化所需资源
func NewProcess(pid int, Max []int, Allocation []int) Process {
	need := make([]int, len(Max))
	for i := 0; i < len(Max); i++ {
		need[i] = Max[i] - Allocation[i]
	}
	return Process{
		PID:        pid,
		Max:        Max,
		Allocation: Allocation,
		Need:       need,
	}
}

func NewEmptySequence() ProcessSequence {
	return make([]Process, 0)
}

//创建进程序列
func NewSequence(p ...Process) ProcessSequence {
	return p
}

func IsAvailable(distribute []int, available []int) bool {
	if len(distribute) != len(available) {
		return false
	}
	for i := 0; i < len(distribute); i++ {
		if distribute[i] > available[i] {
			return false
		}
	}
	return true
}

//分配资源
func Distribute(processes ProcessSequence, PID int, distribute []int, available []int) {
	for i, v := range distribute {
		available[i] -= v
		processes[PID].Allocation[i] += v
		processes[PID].Need[i] -= v
	}
}

func Display(processes ProcessSequence) {
	fmt.Printf("%-4v|%v|%v|%v\n", "PID", "Max", "Allocation", "Need")
	for i := 0; i < len(processes); i++ {
		fmt.Printf("%-4v", processes[i].PID)
		fmt.Printf("%v", processes[i].Max)
		fmt.Printf("%v", processes[i].Allocation)
		fmt.Printf("%v\n", processes[i].Need)
	}
	fmt.Println()
}
