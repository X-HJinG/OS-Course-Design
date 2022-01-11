package main

import (
	mysys "OS/bankerAlgorithm/sys"
	"fmt"
)

func main() {
	available := mysys.NewResources(3, 3, 2)
	//NewProcess(进程号,最大需求资源,已分配资源)
	p0 := mysys.NewProcess(0, mysys.NewResources(7, 5, 3), mysys.NewResources(0, 1, 0))
	p1 := mysys.NewProcess(1, mysys.NewResources(3, 2, 2), mysys.NewResources(2, 0, 0))
	p2 := mysys.NewProcess(2, mysys.NewResources(9, 0, 2), mysys.NewResources(3, 0, 2))
	p3 := mysys.NewProcess(3, mysys.NewResources(2, 2, 2), mysys.NewResources(2, 1, 1))
	p4 := mysys.NewProcess(4, mysys.NewResources(4, 3, 3), mysys.NewResources(0, 0, 2))
	//两种创建进程序列的方式
	// sequence := mysys.NewEmptySequence()
	// sequence = append(sequence, p0, p1, p2, p3, p4)
	sequence := mysys.NewSequence(p0, p1, p2, p3, p4)
	resources := mysys.NewResources(1, 0, 0) //待分配资源
	fmt.Println("Init:")
	mysys.Display(sequence)
	//进行预分配
	tmpSeqence, tmpAvailable, isSuccess := mysys.PreDistribute(sequence, p2.PID, resources, available)
	if isSuccess {
		//进行安全检测
		isSafe, Order := mysys.SecurityCheck(tmpSeqence, tmpAvailable)
		if isSafe {
			fmt.Println("It's safe :", Order)
			//真正进行分配
			mysys.Distribute(sequence, p2.PID, resources, available)
			fmt.Printf("After distribute: %v -> %v\n", p2.PID, resources)
			mysys.Display(sequence)
		} else {
			fmt.Println("It's not safe!")
		}
	} else {
		fmt.Println("bad request!")
	}
}
