package main

import (
	"OS/storage/pagination"
	"OS/storage/subsection"
	mysys "OS/storage/sys"
)

func main() {
	pt := pagination.NewPageTable()
	pt.Item[0] = 2
	pt.Item[1] = 4
	pt.Item[2] = 6
	pt.Item[3] = 7
	pageSize := mysys.NewPageSize("1K")
	addr := mysys.NewLogicAddrForPaging(1023)
	pagination.MappingToPhysical(addr, pageSize, pt)

	st := subsection.NewSectionTable()
	st.Items[0] = subsection.NewItem("50K", "10K")
	st.Items[1] = subsection.NewItem("60K", "3K")
	st.Items[2] = subsection.NewItem("70K", "5K")
	st.Items[3] = subsection.NewItem("120K", "8K")
	st.Items[4] = subsection.NewItem("150K", "4K")
	section_addr := mysys.NewLogicAddrForSection(0, 137)
	subsection.MappingToPhysical(section_addr, st)
}
