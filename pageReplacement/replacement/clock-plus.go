package replacement

import (
	mysys "OS/pageReplacement/sys"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.Mutex

func ClockPlus(block *mysys.Block, pageSequence []int) {
	clock := mysys.NewClock(block.Size)
	wg.Add(2)
	//并发读取页面序列
	go func() {
		//条件为页面序列未空
		for len(pageSequence) > 0 {
			lock.Lock()
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
				curLno := clock.FindByCheckTwo()
				//将循环队列中页面置换
				bNo = clock.OutAndSwap(curLno, newPageInfo)
				//指针移动到下一个页面
				clock.CurLno = (curLno + 1) % clock.Size
			} else {
				clock.Visit(curPage)
			}
			block.ChangeBlock(bNo, curPage)
			lock.Unlock()
			time.Sleep(20 * time.Microsecond)
		}
		wg.Done()
	}()

	// 并发实现随机修改
	go func() {
		for len(pageSequence) > 0 {
			lock.Lock()
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(block.Size)
			if clock.CycleList[n].Pno != -1 {
				clock.CycleList[n].IsModified = true
				fmt.Printf("-----------Modify-----------PNo:%v\n", clock.CycleList[n].Pno)
			}
			lock.Unlock()
			time.Sleep(20 * time.Nanosecond)
		}
		wg.Done()
	}()
	wg.Wait()
}
