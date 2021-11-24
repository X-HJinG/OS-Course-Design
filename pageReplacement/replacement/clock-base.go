package replacement

import (
	mysys "OS/pageReplacement/sys"
)

func ClockBase(block *mysys.Block, pageSequence []int) {
	clock := mysys.NewClock(block.Size)
	//条件为页面序列未空
	for len(pageSequence) > 0 {
		curPage := pageSequence[0]
		pageSequence = pageSequence[1:]
		newPageInfo := mysys.PageInfo{
			Pno:       curPage,
			IsVisited: true,
		}
		//查看是否命中
		bNo, isExist := clock.Get(curPage)
		if !isExist {
			//找到循环队列中的应该置换出的页面
			curLno := clock.FindByCheckOne()
			//将循环队列中页面置换
			bNo = clock.OutAndSwap(curLno, newPageInfo)
			//指针移动到下一个页面
			clock.CurLno = (curLno + 1) % clock.Size
		} else {
			clock.Visit(curPage)
		}
		block.ChangeBlock(bNo, curPage)
	}
}
