package subsection

import (
	"fmt"

	mysys "OS/storage/sys"
)

//段表项
type Item struct {
	addr int //内存始址
	len  int //段长
}

type SectionTable struct {
	Items map[int]*Item //段表项
}

func NewItem(addr string, len string) *Item {
	addrToInt := mysys.ParseInt(addr)
	lenToInt := mysys.ParseInt(len)
	return &Item{
		addr: addrToInt,
		len:  lenToInt,
	}
}

func NewSectionTable() *SectionTable {
	return &SectionTable{
		Items: make(map[int]*Item),
	}
}

func MappingToPhysical(addr *mysys.LogicAddr, sectionTable *SectionTable) {
	no := addr.No //段号
	m, ok := sectionTable.Items[no]
	if !ok {
		fmt.Println("error! Out of Index")
	} else {
		if m.len < addr.Addr {
			fmt.Println("error! Out of Length")
		} else {
			physicalAddress := m.addr + addr.Addr
			fmt.Println("Physical Address:  ", physicalAddress)
		}
	}
}
