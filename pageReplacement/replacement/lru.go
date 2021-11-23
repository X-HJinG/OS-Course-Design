package replacement

import mysys "OS/pageReplacement/sys"

func LRU(block *mysys.Block, pageSequence []int) {
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
				bNo = cache.RemoveFirst()
				cache.Put(curPage, bNo)
			} else {
				cache.ReNew(curPage) //命中则更新页号位置
			}
			block.ChangeBlock(bNo, curPage)
		}
	}
}
