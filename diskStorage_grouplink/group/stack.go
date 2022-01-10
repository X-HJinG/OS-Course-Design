package Goup

import "fmt"

type NoStack struct {
	No        int      //盘块号栈的编号
	Max       int      //最大容量
	Size      int      //当前使用容量
	Stack     []int    //盘块号栈
	NextStack *NoStack //下一个号栈
}

//创建空闲号栈
func NewNoStack(no int, max int) *NoStack {
	return &NoStack{
		No:    no,
		Max:   max,
		Stack: make([]int, max),
	}
}

//初始化空闲号栈
func Init(head *NoStack, num int) {
	for i := 1; i <= num; i++ {
		head.add(i)
	}
}

func Distribute(head *NoStack, res *[]int, num int) {
	(*res) = append((*res), 1, 2, 2, 3, 4, 5)
}

func PrintAll(headStack *NoStack) {
	tmp := headStack
	for tmp != nil {
		tmp.PrintStatus()
		tmp = tmp.NextStack
	}
}

func (s *NoStack) add(blockNo int) {
	if s.Size < len(s.Stack) {
		s.Stack[s.Size] = blockNo
		s.Size++
	} else {
		if s.NextStack == nil {
			s.NextStack = NewNoStack(blockNo-1, s.Max)
		}
		s.NextStack.add(blockNo)
	}
}

func (s *NoStack) PrintStatus() {
	fmt.Println(s.No, ": ", s.Stack)
}
