package schedule

import (
	mysys "OS/diskSchedule/sys"
	"fmt"
)

type subQueues = []int  //切割序列后存储的子队列
type queues []subQueues //存储子队列的数组

//N步扫描法 (N -> 长度为N的子队列)
func NStepSCAN(diskSequence []int, disNo int, N int) mysys.Result {
	res := make(mysys.Result, 0)
	var list queues = make(queues, 0)
	//切割子队列
	for len(diskSequence) > 0 {
		if len(diskSequence) > N {
			//前N个序列号加入子队列
			var newQueue subQueues = diskSequence[:N]
			diskSequence = diskSequence[N:]
			list = append(list, newQueue)
		} else {
			//直接将最后的加入子队列
			var newQueue subQueues = diskSequence
			diskSequence = diskSequence[0:0]
			list = append(list, newQueue)
		}
	}
	fmt.Println(list)
	for i := 0; i < len(list); i++ {
		//每个子队列都采用SCAN算法
		tmp := SCAN(list[i], disNo)
		res = append(res, tmp...)
	}
	return res
}
