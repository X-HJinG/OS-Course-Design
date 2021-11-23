package main

import (
	mysys "OS/distribution/sys"
)

func main() {
	b := mysys.NewBlock(10)
	p1 := mysys.NewProcess(1, 2)
	p2 := mysys.NewProcess(2, 4)
	p3 := mysys.NewProcess(3, 1)
	p4 := mysys.NewProcess(4, 2)
	p5 := mysys.NewProcess(5, 3)
	p6 := mysys.NewProcess(6, 1)
	b.Distribute("NF", p1)
	b.Distribute("NF", p2)
	b.Distribute("NF", p3)
	b.Release(2)
	b.Distribute("NF", p4)
	b.Distribute("NF", p5)
	b.Distribute("NF", p6)
	// b.Distribute("BF", p3)
	b.ShowWorkSpace()
}
