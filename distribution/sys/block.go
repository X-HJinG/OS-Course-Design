package mysys

import (
	"fmt"
	"sort"
)

type block struct {
	allocation []int
	workspace  map[int]map[string]int
}

func NewBlock(size int) *block {
	return &block{
		allocation: make([]int, size),
		workspace:  make(map[int]map[string]int),
	}
}

func setBlock(b *block, p *process, startAddr int) {
	p.startAddr = startAddr
	allocation := b.allocation[startAddr : startAddr+p.size]
	for j := range allocation {
		allocation[j] = p.pid
	}
	info := map[string]int{
		"pid":       p.pid,
		"size":      p.size,
		"startAddr": startAddr,
	}
	b.workspace[startAddr] = info
}

func getFreeSpace(b *block, p *process) map[int]int {
	// 空闲分区排序
	length := len(b.workspace)
	keys := make([]int, 0)
	free := make(map[int]int)
	for k := range b.workspace {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	if keys[0] != 0 {
		//若起始地址未分配，先从起始开始检索
		free[0] = keys[0]
	}
	for i := 0; i < length; i++ {
		//空闲分区的起点
		startAddr := b.workspace[keys[i]]["size"] + keys[i]
		//最后一块分区判断
		if i == length-1 {
			free[startAddr] = len(b.allocation) - startAddr
			break
		}
		if startAddr == keys[i+1] {
			continue
		} else {
			size := keys[i+1] - startAddr
			free[startAddr] = size
		}
	}
	return free
}

func (b *block) Release(pid int) {
	var startAddr int
	var size int
loop:
	for key, values := range b.workspace {
		if values["pid"] == pid {
			size = values["size"]
			startAddr = key
			break loop
		}
	}
	delete(b.workspace, startAddr)
	allocation := b.allocation[startAddr : startAddr+size]
	for i := 0; i < len(allocation); i++ {
		allocation[i] = 0
	}
	fmt.Printf("work%v---release\n", pid)
}

func (b *block) ShowWorkSpace() {
	var keys = make([]int, 0)
	for key := range b.workspace {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("%v\n", b.workspace[key])
	}
}
