package schedule

import (
	mysys "OS/diskSchedule/sys"
	"sort"
)

//扫描算法
func SCAN(diskSequence []int, disNo int) mysys.Result {
	//当前的序列进行升序排列
	sort.Ints(diskSequence)
	//找到数组的旋转点 [1,2,3,4,5] -> [4,5,1,2,3] : 4即为旋转点
	//先将旋转点前部分倒置 ->[3,2,1,4,5]
	// 最后完成旋转 ->[4,5,3,2,1]
	for i := 0; i < len(diskSequence); i++ {
		//找到旋转点
		if diskSequence[i] >= disNo {
			front := diskSequence[:i]
			back := diskSequence[i:]
			//前部分倒置
			sort.Slice(front, func(i, j int) bool {
				return front[i] > front[j]
			})
			//切片拼接,实现旋转
			diskSequence = append(back, front...)
			break
		}
	}
	res := FCFS(diskSequence, disNo)
	return res
}
