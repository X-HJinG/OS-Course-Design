package simplework

import (
	"container/list"
	"fmt"
	"sort"
)

func FCFS(works []work) {
	//根据到达时间降序排列
	res := make([]work, 0)
	sort.Slice(works, func(i, j int) bool {
		return works[i].arriveTime < works[j].arriveTime
	})
	for i := 0; i < len(works); i++ {
		res = append(res, works[i])
	}
	fmt.Println(res)
	fmt.Println()
	showStatus(res)
}

func SJF(works []work) {
	curTime := 0
	res := make([]work, 0)
	//根据到达时间降序排列
	sort.Slice(works, func(i, j int) bool {
		return works[i].arriveTime < works[j].arriveTime
	})
	res = append(res, works[0])
	curTime = works[0].serviceTime
	works = works[1:]
	//从时间降序挑选短作业
	for len(works) > 0 {
		min := works[0].serviceTime
		index := 0
		for i := 1; i < len(works); i++ {
			if works[i].arriveTime <= curTime && works[i].serviceTime < min {
				min = works[i].serviceTime
				index = i
			}
			if works[i].arriveTime > curTime {
				break
			}
		}
		curTime += works[index].serviceTime
		res = append(res, works[index])
		works = append(works[:index], works[index+1:]...)
	}
	fmt.Println(res)
	fmt.Println()
	showStatus(res)

}

func FB(works []work, isPreemptive bool) {
	res := make([]work, 0)
	queues := make([]*list.List, 1, 10)
	queues[0] = list.New()
	sort.Slice(works, func(i, j int) bool {
		return works[i].arriveTime < works[j].arriveTime
	})
	//非抢占式
	if !isPreemptive {
		curTime := 0
		isEmpty := false
		for len(works) > 0 || !isEmpty {
			//作业是否到达
			if len(works) > 0 && works[0].arriveTime <= curTime {
				//到达队列置于一级队列末端
				curWork := works[0]
				queues[0].PushBack(curWork)
				works = works[1:]
			}
			//特殊情况：就绪队列之中没有作业，作业队列里仍有未抵达的队列
			if isEmpty && len(works) > 0 {
				curWork := works[0]
				queues[0].PushBack(curWork)
				works = works[1:]
			}
			//从队列从取一个作业执行n个时间片的操作
			for i := 0; i < len(queues); i++ {
				//顺序查找，队列不为空，在当前队列对作业进行一个时间片的操作
				if queues[i].Len() != 0 {
					e := queues[i].Front()
					w := e.Value.(work)
					queues[i].Remove(e)
					timeSlice := 1 << i
					//如果一个时间片内完成任务
					if w.serviceTime-timeSlice <= 0 {
						curTime += w.serviceTime
						res = append(res, w)
					} else {
						curTime += timeSlice
						tmp := work{
							serviceTime: timeSlice,
							arriveTime:  w.arriveTime,
							pName:       w.pName,
						}
						res = append(res, tmp)
						w.serviceTime = w.serviceTime - timeSlice
						//未执行完的进入下一级队列,不存在则创建，存在则插入
						if i+1 >= len(queues) {
							// queues = append(queues, list.List{})
							queues = append(queues, list.New())
						}
						queues[i+1].PushBack(w)
					}
					break
				}
			}
			//每次结束队列的扫描，判断多级队列是否为空
			for i := 0; i < len(queues); i++ {
				if queues[i].Len() != 0 {
					isEmpty = false
					break
				}
				isEmpty = true
			}
		}
		fmt.Println(res)
		fmt.Println()
		showStatus(res)
	} else {
		//抢占
		isEmpty := false
		curTime := 0
		for len(works) > 0 || !isEmpty {
			if len(works) > 0 && works[0].arriveTime <= curTime {
				curWork := works[0]
				queues[0].PushBack(curWork)
				works = works[1:]
			}
			if isEmpty && len(works) > 0 {
				curWork := works[0]
				queues[0].PushBack(curWork)
				works = works[1:]
			}
			for i := 0; i < len(queues); i++ {
				if queues[i].Len() != 0 {
					e := queues[i].Front()
					w := e.Value.(work)
					queues[i].Remove(e)
					timeSlice := 1 << i
					targetTime := w.serviceTime
					isInterrupt := false
					//抢占的特点，在作业执行的时机，每单个时间片都需要进行判断
					for t := 1; t <= timeSlice; t++ {
						targetTime--
						curTime++
						//是否会发生中断
						if len(works) > 0 && curTime == works[0].arriveTime {
							newWork := works[0]
							queues[0].PushBack(newWork)
							works = works[1:]
							isInterrupt = true
							//中断且完成任务
							if targetTime == 0 {
								res = append(res, w)
							}
							break
						}
						//未中断，且完成任务
						if targetTime == 0 {
							res = append(res, w)
							break
						}
					}
					//时间片未执行完,判断是否发生中断，选择是否进入下一级队列
					if targetTime != 0 {
						tmp := work{
							serviceTime: w.serviceTime - targetTime,
							arriveTime:  w.arriveTime,
							pName:       w.pName,
						}
						//将已运转的时间片加入结果
						res = append(res, tmp)
						//更新未执行完的作业
						w.serviceTime = targetTime
						if isInterrupt {
							queues[i].PushBack(w)
						} else {
							//未中断且未执行完 ，进入下一级就绪队列
							if i+1 >= len(queues) {
								// queues = append(queues, list.List{})
								queues = append(queues, list.New())
							}
							queues[i+1].PushBack(w)
						}
					}
					break
				}
			}
			//每次结束队列的扫描，判断多级队列是否为空
			for i := 0; i < len(queues); i++ {
				if queues[i].Len() != 0 {
					isEmpty = false
					break
				}
				isEmpty = true
			}
		}
		fmt.Println(res)
		fmt.Println()
		showStatus(res)
	}
}
