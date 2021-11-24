package mysys

import "container/list"

type Cache struct {
	Size int
	Map  map[int]int //[页号]块号
	List *list.List  // 缓存队列
}

func NewCache(size int) *Cache {
	return &Cache{
		Size: size,
		List: list.New(),
		Map:  make(map[int]int),
	}
}

func (c *Cache) IsFull() bool {
	return c.List.Len() >= c.Size
}

func (c *Cache) ReNew(pNo int) {
	bNo := c.Remove(pNo)
	c.Put(pNo, bNo)
}

func (c *Cache) Put(pNo int, bNo int) {
	c.Map[pNo] = bNo
	c.List.PushBack(pNo)
}

func (c *Cache) Get(pNo int) (int, bool) {
	bNo, ok := c.Map[pNo]
	return bNo, ok
}

func (c *Cache) Remove(pNo int) int {
	bNo := c.Map[pNo]
	delete(c.Map, pNo)
	for e := c.List.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == pNo {
			c.List.Remove(e)
			break
		}
	}
	return bNo
}

func (c *Cache) RemoveFirst() int {
	e := c.List.Front()
	pNo := e.Value.(int)
	c.List.Remove(e)
	bNo := c.Map[pNo]
	delete(c.Map, pNo)
	return bNo
}
