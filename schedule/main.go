package main

import (
	"OS/schedule/simplework"
)

func main() {
	w1 := simplework.NewWork("A", 0, 3)
	w2 := simplework.NewWork("B", 2, 6)
	w3 := simplework.NewWork("C", 4, 4)
	w4 := simplework.NewWork("D", 6, 5)
	w5 := simplework.NewWork("E", 8, 2)
	queue := simplework.NewQueue(*w1, *w2, *w3, *w4, *w5)
	// simplework.FCFS(queue)
	// simplework.SJF(queue)
	// simplework.FB(queue, false) //非抢占
	simplework.FB(queue, true) //抢占
}
