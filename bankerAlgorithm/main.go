package main

import (
	mysys "OS/bankerAlgorithm/sys"
	"fmt"
)

func main() {
	available := mysys.NewResources(3, 3, 2)
	p0 := mysys.NewProcess(0, mysys.NewResources(7, 5, 3), mysys.NewResources(0, 1, 0))
	p1 := mysys.NewProcess(1, mysys.NewResources(3, 2, 2), mysys.NewResources(2, 0, 0))
	p2 := mysys.NewProcess(2, mysys.NewResources(9, 0, 2), mysys.NewResources(3, 0, 2))
	p3 := mysys.NewProcess(3, mysys.NewResources(2, 2, 2), mysys.NewResources(2, 1, 1))
	p4 := mysys.NewProcess(4, mysys.NewResources(4, 3, 3), mysys.NewResources(0, 0, 2))
	sequence := mysys.NewEmptySequence()
	sequence = append(sequence, p0, p1, p2, p3, p4)
	isSuccess, Order := mysys.SecurityCheck(sequence, available)
	if isSuccess {
		fmt.Println(Order)
	}
}
