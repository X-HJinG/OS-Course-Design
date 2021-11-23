package simplework

import (
	"fmt"
	"sort"
)

type work struct {
	pName       string
	arriveTime  int
	serviceTime int
}

type result struct {
	arriveTime             int     //到达时间
	serviceTime            int     //服务时间
	completionTime         int     //完成时间
	turnaroundTime         int     //周转时间
	weightedTurnaroundTime float64 //带权周转时间
}

func NewWork(name string, t1 int, t2 int) *work {
	return &work{
		pName:       name,
		arriveTime:  t1,
		serviceTime: t2,
	}
}

func NewQueue(works ...work) []work {
	queue := make([]work, 0)
	queue = append(queue, works...)
	return queue
}

func showStatus(res []work) {
	details := make(map[string]*result)
	curTime := 0
	for _, v := range res {
		m, ok := details[v.pName]
		curTime += v.serviceTime
		if !ok {
			details[v.pName] = &result{
				arriveTime:             v.arriveTime,
				serviceTime:            v.serviceTime,
				completionTime:         curTime,
				turnaroundTime:         curTime - v.arriveTime,
				weightedTurnaroundTime: float64(curTime-v.arriveTime) / float64(v.serviceTime),
			}
		} else {
			m.serviceTime += v.serviceTime
			m.completionTime = curTime
			m.turnaroundTime = m.completionTime - m.arriveTime
			m.weightedTurnaroundTime = float64(m.turnaroundTime) / float64(m.serviceTime)
		}
	}
	keys := make([]string, 0, len(details))
	for k := range details {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Print(key, ":")
		fmt.Printf("{arriveTime:%v ", details[key].arriveTime)
		fmt.Printf("serviceTime:%v ", details[key].serviceTime)
		fmt.Printf("completionTime:%v ", details[key].completionTime)
		fmt.Printf("turnaroundTime:%v ", details[key].turnaroundTime)
		fmt.Printf("weightedTurnaroundTime:%.2f}\n", details[key].weightedTurnaroundTime)
		fmt.Println()
	}
}
