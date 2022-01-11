package Group

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
	for i := 1; i < num; i++ {
		head.Add(i)
	}
	if num%head.Max == 0 {
		head.Add(-1)
	} else {
		head.Add(num)
	}
}

//分配空闲区块 num->为分配个数
func Distribute(head *NoStack, res *[]int, num int) {
	for i := 0; i < num; i++ {
		hasNext, blockNo := head.Remove()
		*res = append(*res, blockNo)
		if !hasNext {
			if head.NextStack != nil {
				*head = *head.NextStack
				i--
			} else {
				newStack := NewNoStack(0, head.Max)
				*head = *newStack
				fmt.Println("error-> out of index !!")
				break
			}
		}
	}
}

//回收号栈
func Recycle(head *NoStack, allocatedBlock []int) {
	for len(allocatedBlock) > 0 {
		head.Add(allocatedBlock[0])
		allocatedBlock = allocatedBlock[1:]
	}
}

func PrintAll(headStack *NoStack) {
	tmp := headStack
	for tmp != nil {
		tmp.PrintStatus()
		tmp = tmp.NextStack
	}
}

func (s *NoStack) Add(blockNo int) {
	//判断是否仍有空闲位置未分配
	if s.Size < s.Max {
		s.Stack[s.Size] = blockNo
		s.Size++
	} else {
		if s.NextStack == nil {
			//判断是否为终止位置
			if s.Stack[s.Max-1] == -1 {
				s.Stack[s.Max-1] = blockNo
				s.NextStack = NewNoStack(blockNo, s.Max)
				return
			} else {
				newStackNo := s.Stack[s.Max-1]
				s.NextStack = NewNoStack(newStackNo, s.Max)
			}
		}
		s.NextStack.Add(blockNo)
	}
	//添加终止标志
	if s.Size < s.Max {
		s.Stack[s.Size] = -1
	}
}

func (s *NoStack) Remove() (bool, int) {
	//从空闲号栈顶删除一个，直到栈底
	if len(s.Stack) > 1 {
		blockNo := s.Stack[0]
		s.Stack = s.Stack[1:]
		return true, blockNo
	}
	return false, s.No
}

func (s *NoStack) PrintStatus() {
	fmt.Printf("%-4d: %v\n", s.No, s.Stack)
}
