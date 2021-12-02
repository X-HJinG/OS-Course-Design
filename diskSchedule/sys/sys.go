package mysys

import "fmt"

type diskSequence = []int

type item struct {
	DisNo    int //磁道号
	Distance int //移动距离
}

type Result []item //输出结果封装

func NewDiskSequence(sequence ...int) diskSequence {
	return sequence
}

func NewItem(disNo int, distance int) item {
	return item{
		DisNo:    disNo,
		Distance: distance,
	}
}

func (res Result) Print() {
	sum := 0
	fmt.Printf("\n|%-12v|%-12v|\n", "disNo", "distance")
	for _, item := range res {
		fmt.Printf("|%-12v|%-12v|\n", item.DisNo, item.Distance)
		sum += item.Distance
	}
	fmt.Printf("|AverageSeekLength:%-7.1f|\n", float64(sum)/float64(len(res)))
	fmt.Println()
}
