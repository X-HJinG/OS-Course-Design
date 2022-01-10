package mysys

//进行一次预分配
func PreDistribute(processes ProcessSequence, PID int, distribute []int, available []int) ([]Process, []int) {
	//将传递的切片进行拷贝 防止修改原先传入的变量
	tempProcessSequence := make([]Process, len(processes))
	tempAvailable := make([]int, len(available))
	copy(tempAvailable, available)
	//拷贝序列
	copy(tempProcessSequence, processes)
	Distribute(tempProcessSequence, PID, distribute, tempAvailable)
	return tempProcessSequence, tempAvailable
}

func SecurityCheck(processes ProcessSequence, available []int) (bool, []int) {
	securityOrder := make([]int, 0)
	//记录遍历一次进程序列,是否找到可以分配完资源的进程并回收资源
	isFind := true
	for len(processes) > 0 && isFind {
		isFind = false
		for i := 0; i < len(processes); i++ {
			//如果把资源一次分配完，则返回已分配资源
			if IsAvailable(processes[i].Need, available) {
				securityOrder = append(securityOrder, processes[i].PID)
				for j := 0; j < len(processes[i].Allocation); j++ {
					available[j] += processes[i].Allocation[j]
				}
				//将获得所需资源的进程删除
				isFind = true
				processes = append(processes[:i], processes[i+1:]...)
			}
		}
	}
	return len(processes) == 0, securityOrder
}
