package main

import (
	Group "OS/diskStorage_grouplink/group"
	"fmt"
)

const (
	N int = 20 //每个号栈的最大容量
)

func main() {
	allocatedBlock := make([]int, 0)
	headStack := Group.NewNoStack(0, N)
	Group.Init(headStack, 110)
	Group.Distribute(headStack, &allocatedBlock, 1)
	fmt.Println(allocatedBlock)
	// Group.PrintAll(headStack)
}
