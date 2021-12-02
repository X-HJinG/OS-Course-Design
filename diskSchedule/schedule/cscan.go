package schedule

import (
	mysys "OS/diskSchedule/sys"
	"math"
	"sort"
)

//循环扫描算法 (dir 为磁头移动方向,0为向里,1为向外)
func CSCAN(diskSequence []int, disNo int, dir int) mysys.Result {
	res := make(mysys.Result, 0)
	start := 0 // 磁头定位开始扫描的索引
	// 根据磁头移动方向 决定 序列为升序或是降序
	if dir == 0 {
		//降序 磁头向里
		sort.Ints(diskSequence)
	} else if dir == 1 {
		//升序 磁头向外
		sort.Slice(diskSequence, func(i, j int) bool {
			return diskSequence[i] > diskSequence[j]
		})
	}
	for i := 0; i < len(diskSequence); i++ {
		if dir == 0 && diskSequence[i] >= disNo { // 磁头向里 查找第一个比开始磁道大的索引
			start = i
			break
		} else if dir == 1 && diskSequence[i] <= disNo { // 磁头向外 查找第一个比开始磁道小的索引
			start = i
			break
		}
	}
	//完整扫描一次队列,循环扫描回去
	times := 0
	for i := start; times < len(diskSequence); i = (i + 1) % len(diskSequence) {
		//当前序列号
		curSequence := diskSequence[i]
		distance := math.Abs(float64(disNo) - float64(curSequence))
		res = append(res, mysys.NewItem(curSequence, int(distance)))
		disNo = curSequence
		times++
	}
	return res
}
