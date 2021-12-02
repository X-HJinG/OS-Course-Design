package schedule

import (
	mysys "OS/diskSchedule/sys"
	"math"
)

//先来先服务
func FCFS(diskSequence []int, disNo int) mysys.Result {
	res := make(mysys.Result, 0)
	for len(diskSequence) > 0 {
		//当前序列号
		curSequence := diskSequence[0]
		//修改序列队列
		diskSequence = diskSequence[1:]
		//求绝对值
		distance := math.Abs(float64(disNo) - float64(curSequence))
		res = append(res, mysys.NewItem(curSequence, int(distance)))
		disNo = curSequence
	}
	return res
}
