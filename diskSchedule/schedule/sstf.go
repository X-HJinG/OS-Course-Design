package schedule

import (
	mysys "OS/diskSchedule/sys"
	"math"
)

//最短寻道时间优先
func SSTF(diskSequence []int, disNo int) mysys.Result {
	res := make(mysys.Result, 0)
	for len(diskSequence) > 0 {
		// 最短寻道时间序列号的索引
		minIndex := 0
		// 最短寻道时间序列号
		minDiskNo := diskSequence[0]
		//最短寻道距离
		minDistance := int(math.Abs(float64(disNo) - float64(diskSequence[0])))
		//循环查找 最短寻道时间序列号
		for i := 1; i < len(diskSequence); i++ {
			cur := int(math.Abs(float64(disNo) - float64(diskSequence[i])))
			if cur < minDistance {
				minDistance = cur
				minDiskNo = diskSequence[i]
				minIndex = i
			}
		}
		//从切片中清除最短寻道时间序列号
		if minIndex >= len(diskSequence)-1 {
			diskSequence = diskSequence[:minIndex]
		} else {
			diskSequence = append(diskSequence[:minIndex], diskSequence[minIndex+1:]...)
		}
		res = append(res, mysys.NewItem(minDiskNo, minDistance))
		disNo = minDiskNo
	}
	return res
}
