package mysys

import (
	"fmt"
	"os"
)

func ReadFile(filename string) {
	file, err := os.Open("./sample.txt")
	if err != nil {
		fmt.Println("Open file failed,err:", err)
		return
	}
	defer file.Close()
}
