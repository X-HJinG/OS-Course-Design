package main

import (
	mysys "OS/distribution/sys"
)

func main() {
	b := mysys.NewBlock(25)
	p1 := mysys.NewProcess(1, 2)
	p2 := mysys.NewProcess(2, 4)
	p3 := mysys.NewProcess(3, 1)
	p4 := mysys.NewProcess(4, 2)
	p5 := mysys.NewProcess(5, 3)
	p6 := mysys.NewProcess(6, 1)
	b.Distribute("WF", p1, p2, p3, p4, p5)
	b.ShowWorkSpace()
	b.Release(2, 3)
	b.Distribute("WF", p3, p6, p2)
	b.ShowWorkSpace()
}
