package pagination

import (
	"fmt"

	mysys "OS/storage/sys"
)

type PageTable struct {
	Item map[int]int //页表项
}

func NewPageTable() *PageTable {
	return &PageTable{
		Item: make(map[int]int),
	}
}

func MappingToPhysical(addr *mysys.LogicAddr, pageSize *mysys.PageSize, table *PageTable) {
	no := addr.Addr / pageSize.Size // 页号
	blockNo, ok := table.Item[no]
	if !ok {
		fmt.Println("error! Out of Index")
	} else {
		d := addr.Addr % pageSize.Size //页内偏移量
		physical := blockNo*pageSize.Size + d
		fmt.Println("Physical Address:  ", physical)
	}
}
