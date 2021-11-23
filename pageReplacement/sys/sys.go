package mysys

import "fmt"

type Block struct {
	Size  int
	Store []int // 块号:页号
}

func NewBlock(size int) *Block {
	store := make([]int, size)
	for i := 0; i < len(store); i++ {
		store[i] = -1
	}
	return &Block{
		Size:  size,
		Store: store,
	}
}

func (b *Block) ChangeBlock(bNo int, pNo int) {
	flag := false
	if b.Store[bNo] != pNo {
		b.Store[bNo] = pNo
	} else {
		flag = true
	}
	fmt.Print(pNo, "  [")
	for i := 0; i < len(b.Store); i++ {
		fmt.Printf("%4d", b.Store[i])
	}
	fmt.Printf("%8v]\n", flag)
}
