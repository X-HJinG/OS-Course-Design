package schedule

import mysys "OS/diskSchedule/sys"

type curQueue = []int
type nextQueue = []int

func FSCAN(diskSequence []int, nextSequence []int, disNo int) mysys.Result {
	res := make(mysys.Result, 0)
	cur := diskSequence
	next := nextSequence
	for len(next) != 0 {
		tmp := SCAN(cur, disNo)
		res = append(res, tmp...)
		cur = next
	}
	return res
}
