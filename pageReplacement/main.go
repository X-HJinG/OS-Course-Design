package main

import (
	"OS/pageReplacement/replacement"
	mysys "OS/pageReplacement/sys"
	"fmt"
	"time"
)

func main() {
	pageSequence := []int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2, 1, 2, 0, 1, 7, 0, 1}
	block := mysys.NewBlock(3)
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
	time.Sleep(time.Hour)
	fmt.Printf("s")
}
