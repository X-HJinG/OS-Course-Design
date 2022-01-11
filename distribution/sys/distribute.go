package mysys

import (
	"fmt"
	"sort"
)

var pre int

func (b *block) Distribute(mode string, processes ...*process) {
	var startAddr int
	var ok bool
	for _, p := range processes {
		switch mode {
		case "FF":
			startAddr, ok = FF(b, p)
		case "NF":
			startAddr, ok = NF(b, p, pre)
			if ok {
				pre = startAddr
			}
		case "BF":
			startAddr, ok = BF(b, p)
		case "WF":
			startAddr, ok = WF(b, p)
		}
		if ok {
			setBlock(b, p, startAddr)
			fmt.Printf("work%v---success!\n", p.pid)
		} else {
			fmt.Println("There isn't enough space")
		}
	}

}

func FF(b *block, p *process) (int, bool) {
	cursize := p.size
	cnt := 0
	for i := 0; i < len(b.allocation); i++ {
		for i < len(b.allocation) && b.allocation[i] != 0 {
			i++
		}
		startAddr := i
		for i < len(b.allocation) && b.allocation[i] == 0 && cnt < cursize {
			cnt++
			if cnt == cursize {
				return startAddr, true
			}
			i++
		}
		cnt = 0
	}
	return -1, false
}

func NF(b *block, p *process, pre int) (int, bool) {
	if len(b.workspace) > 0 {
		keys := make([]int, 0)
		free := getFreeSpace(b, p)
		for k := range free {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		//查找上次最后一块
		for i, v := range keys {
			if pre < v {
				if p.size <= free[v] {
					return v, true
				} else {
					if i+1 > len(keys) {
						keys = keys[:i]
					} else {
						keys = append(keys[:i], keys[i+1:]...)
					}
				}
			}
		}
		//从头找
		for _, k := range keys {
			if p.size <= free[k] {
				return k, true
			}
		}
		return -1, false
	} else {
		return FF(b, p)
	}
}

func BF(b *block, p *process) (int, bool) {
	if len(b.workspace) > 0 {
		var (
			min       int = 10000000
			startAddr int = -1
		)
		free := getFreeSpace(b, p)
		for k, v := range free {
			if p.size <= v && v-p.size < min {
				min = v - p.size
				startAddr = k
			}
		}
		if startAddr > -1 {
			return startAddr, true
		}
		return -1, false
	} else {
		return FF(b, p)
	}
}

func WF(b *block, p *process) (int, bool) {
	if len(b.workspace) > 0 {
		var (
			max       int = 0
			startAddr int = -1
		)
		free := getFreeSpace(b, p)
		for k, v := range free {
			if p.size <= v && v-p.size > max {
				max = v - p.size
				startAddr = k
			}
		}
		if startAddr > -1 {
			return startAddr, true
		}
		return -1, false
	} else {
		return FF(b, p)
	}
}
