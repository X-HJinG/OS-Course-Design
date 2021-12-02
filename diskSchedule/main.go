package main

import (
	"OS/diskSchedule/schedule"
	mysys "OS/diskSchedule/sys"
)

type diskNo = int

func main() {
	var result mysys.Result
	sequence := mysys.NewDiskSequence(55, 58, 39, 18, 90, 160, 150, 38, 184)
	// result = schedule.FCFS(sequence, 100) //先来先服务
	// result = schedule.SSTF(sequence, 100) //最短寻道时间优先
	// result = schedule.SCAN(sequence, 100) //扫描算法
	// result = schedule.CSCAN(sequence, 100, 0) //循环扫描算法，向里
	// result = schedule.CSCAN(sequence, 100, 1) //循环扫描算法，向外
	result = schedule.NStepSCAN(sequence, 100, 4) //N步扫描算法
	result.Print()
}
