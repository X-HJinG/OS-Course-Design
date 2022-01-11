package main

import (
	Group "OS/diskStorage_grouplink/group"
	"fmt"
)

const (
	N int = 20 //每个号栈的最大容量
)

func main() {
	allocatedBlock := make([]int, 0, 1000)
	headStack := Group.NewNoStack(0, N)
	Group.Init(headStack, 120)
	fmt.Println("Init:")
	Group.PrintAll(headStack)
	Group.Distribute(headStack, &allocatedBlock, 40)
	fmt.Printf("\nafter distribute:\n")
	Group.PrintAll(headStack)
	Group.Recycle(headStack, allocatedBlock)
	fmt.Printf("\nafter recycle:\n")
	Group.PrintAll(headStack)
}
