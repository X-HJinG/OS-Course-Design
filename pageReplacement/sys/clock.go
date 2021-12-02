package mysys

type PageInfo struct {
	Pno        int  //页号
	IsVisited  bool // [访问位]--是否访问过
	IsModified bool // [修改位]--是否修改
}

type Clock struct {
	Size      int         //时钟大小 == 页面数量
	Map       map[int]int //记录[页号]块号的映射
	CycleList []PageInfo  // 用切片模拟循环队列 -> 存页号
	CurLno    int         //当前访问的循环队列的位置
}

func NewClock(size int) *Clock {
	//创建时，将循环队列中的页号初始化为-1
	newPageInfos := make([]PageInfo, size)
	for i := 0; i < len(newPageInfos); i++ {
		newPageInfos[i].Pno = -1
	}
	return &Clock{
		Size:      size,
		Map:       make(map[int]int),
		CycleList: newPageInfos,
		CurLno:    0,
	}
}

func (c *Clock) Get(pNo int) (bNo int, ok bool) {
	bNo, ok = c.Map[pNo]
	return
}

func (c *Clock) Put(pNo int, bNo int) {
	c.Map[pNo] = bNo
}

//调整访问位
func (c *Clock) Visit(pNo int) {
	for i := 0; i < len(c.CycleList); i++ {
		if c.CycleList[i].Pno == pNo {
			c.CycleList[i].IsVisited = true
			break
		}
	}
}

//通过检查访问位，查找被置换的页面
func (c *Clock) FindByCheckOne() (curLno int) {
	for i := c.CurLno; i < c.Size; i = (i + 1) % c.Size {
		if !c.CycleList[i].IsVisited {
			//标记位为false，直接记录当前位置
			curLno = i
			break
		} else {
			//访问过,标记位改为false
			c.CycleList[i].IsVisited = false
		}
	}
	return
}

//通过检查访问位与修改位，查找被置换的页面
func (c *Clock) FindByCheckTwo() (curLno int) {
	//是否扫完一圈
	isFinished := false
	for {
		//第一次查找访问位与修改位均为 false
		for i := c.CurLno; i < c.Size; i = (i + 1) % c.Size {
			if i == c.CurLno && isFinished {
				break
			}
			if !c.CycleList[i].IsVisited && !c.CycleList[i].IsModified {
				curLno = i
				return
			}
			isFinished = true
		}
		//重置标记
		isFinished = false
		//第二轮查找访问位为false 且 修改位为true，同时将所有查找过的页面的访问位置为false
		for i := c.CurLno; i < c.Size; i = (i + 1) % c.Size {
			if i == c.CurLno && isFinished {
				break
			}
			if !c.CycleList[i].IsVisited && c.CycleList[i].IsModified {
				curLno = i
				return
			} else {
				c.CycleList[i].IsVisited = false
			}
			isFinished = true
		}
		// 若仍未找到则重置标记, 重复第一次扫描
		isFinished = false
	}
}

func (c *Clock) OutAndSwap(Lno int, newPageInfo PageInfo) int {
	//获取即将出列的队列元素的页号
	pNo := c.CycleList[Lno].Pno
	//若存在，则删除该页号与块号的映射，若不存在说明内存块仍有剩余
	bNo, ok := c.Map[pNo]
	if ok {
		delete(c.Map, pNo)
	} else {
		bNo = c.CurLno
	}
	//将换出的页面位置更换为新进入的页面
	c.CycleList[Lno] = newPageInfo
	//更新映射关系
	c.Map[newPageInfo.Pno] = bNo
	return bNo
}
