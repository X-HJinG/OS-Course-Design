package replacement

import mysys "OS/pageReplacement/sys"

func LFU(block *mysys.Block, pageSequence []int) {
	cache := mysys.NewCache(block.Size)
	freq := make(map[int]int) // 记录频次的表 [页号]访问次数
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
			freq[curPage]++
			block.ChangeBlock(bNo, curPage)
		} else {
			// 判断是否命中
			bNo, ok := cache.Get(curPage)
			if !ok {
				minFreq := 100000
				minPno := -1
				//7, 0, 1        2, 0, 3, 0, 4, 2, 3, 0, 3, 2, 1, 2, 0, 1, 7, 0, 1
				//按照缓存队列顺序遍历，获取访问次数最少页面（存在多个最少 保证是入队较早的页面）
				for e := cache.List.Front(); e != nil; e = e.Next() {
					if freq[e.Value.(int)] < minFreq {
						minFreq = freq[e.Value.(int)]
						minPno = e.Value.(int)
					}
				}
				// 删除访问次数最少的页面
				delete(freq, minPno)
				bNo = cache.Remove(minPno)
				cache.Put(curPage, bNo)
			} else {
				cache.ReNew(curPage) //命中则更新页号位置
			}
			freq[curPage]++
			block.ChangeBlock(bNo, curPage)
		}
	}
}
