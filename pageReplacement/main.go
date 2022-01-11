package main

import (
	"OS/pageReplacement/replacement"
	mysys "OS/pageReplacement/sys"
	"fmt"
)

func main() {
	pageSequence := []int{4, 3, 2, 1, 4, 3, 5, 4, 3, 2, 1, 5}
	block := mysys.NewBlock(4)
	fmt.Print("-  {")
	for i := 0; i < block.Size; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Printf("%8v}\n", "ISHIT")
	// replacement.Optimal(block, pageSequence)
	// replacement.FIFO(block, pageSequence)
	// replacement.LRU(block, pageSequence)
	// replacement.LFU(block, pageSequence)
	// replacement.ClockBase(block, pageSequence)
	replacement.ClockPlus(block, pageSequence)
}
