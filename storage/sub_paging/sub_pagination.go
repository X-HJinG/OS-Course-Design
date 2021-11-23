package subpaging

type PageTable struct {
	Item map[int]int
}

type Item struct {
	Size int
	Pt   PageTable
}

type Sub_PageTable struct {
	Items map[int]*Item
}
