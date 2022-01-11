package main

import (
	Bitmap "OS/diskStorage_bitmap/bitmap"
	"bufio"
	"fmt"
	"os"
)

const (
	SizeofClinders int = 10 //柱面数
	MachineWord    int = 16 //机器字长
	SizeofTrack    int = 4  //每个柱面的磁道数
)

var (
	rblockNo   int //相对块号
	cylinderNo int //柱面号
	trackNo    int //磁道号
	blockNo    int //物理块号
)

func main() {
	choice := 0
	bitmap := Bitmap.New(MachineWord, SizeofClinders)
	bitmap.Init()
	for {
		fmt.Println("1.add  2.remove 3.printBitMap")
		fmt.Print("Enter the value of your choice:")
		fmt.Scanf("%d\n", &choice)
		fmt.Println()
		switch choice {
		case 1:
			isAdd, row, column := bitmap.Add()
			if isAdd {
				calResult(row, column)
				fmt.Print("CylinderNo:", cylinderNo, "\tTrackNo:", trackNo, "\tBlockNo:", blockNo, "\n\n")
			} else {
				fmt.Print("These areas are full\n\n")
			}
		case 2:
			fmt.Print("Enter the cylinderNo,trackNo and blockNo:")
			fmt.Scanf("%d%d%d\n", &cylinderNo, &trackNo, &blockNo)
			rblockNo = cylinderNo*MachineWord + trackNo*SizeofTrack + blockNo
			isRemove := bitmap.Remove(rblockNo/MachineWord, rblockNo%MachineWord)
			if isRemove {
				fmt.Println("rblockNo is ", rblockNo)
				fmt.Print("blockNo:", blockNo, " has been released.\n\n")
			} else {
				fmt.Print("Remove failed\n\n")
			}
		case 3:
			bitmap.Print()
		default:
			fmt.Println("bad choice ! !")
			r := bufio.NewReader(os.Stdin)
			r.ReadLine()
		}
	}
}

func calResult(row int, column int) {
	rblockNo = row*MachineWord + column
	cylinderNo = rblockNo / MachineWord
	trackNo = rblockNo % MachineWord / SizeofTrack
	blockNo = rblockNo % MachineWord % SizeofTrack
}
