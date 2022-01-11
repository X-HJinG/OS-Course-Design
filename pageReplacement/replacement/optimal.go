package replacement

import mysys "OS/pageReplacement/sys"

func Optimal(block *mysys.Block, pageSequence []int) {
	cache := mysys.NewCache(block.Size)
	//条件为页面序列未空
	for len(pageSequence) > 0 {
		curPage := pageSequence[0]
		pageSequence = pageSequence[1:]
		// 缓存若未满，则初始化直接插入
		if !cache.IsFull() {
			bNo := cache.List.Len()
			if oldBno, ok := cache.Get(curPage); !ok {
				cache.Put(curPage, bNo)
			} else {
				bNo = oldBno
			}
			block.ChangeBlock(bNo, curPage)
		} else {
			// 判断是否命中
			bNo, ok := cache.Get(curPage)
			if !ok {
				flag := false
				// 创建一个map 记录当前缓存队列的页面顺序 --- [页号]默认最近使用顺序
				record := make(map[int]int)
				bigInt := 100000
				for e := cache.List.Front(); e != nil; e = e.Next() {
					record[e.Value.(int)] = bigInt
					bigInt--
				}
				// 从后向前遍历页面序列，更新页面最近使用顺序
				for i := len(pageSequence) - 1; i >= 0; i-- {
					_, isExist := cache.Get(pageSequence[i])
					if isExist {
						record[pageSequence[i]] = i
						flag = true
					}
				}
				// 如果找到缓存队列中有页面在未来仍被使用，则选取最久未使用
				if flag {
					max := -1
					longTimePage := 0
					for k, v := range record {
						if v > max {
							max = v
							longTimePage = k
						}
					}
					bNo = cache.Remove(longTimePage)
				} else { // 若无页面在未来被使用，则选取最早进入缓存队列的页面
					bNo = cache.RemoveFirst()
				}
				//更新缓存队列
				cache.Put(curPage, bNo)
			}
			//更新物理块信息
			block.ChangeBlock(bNo, curPage)
		}
	}
}
