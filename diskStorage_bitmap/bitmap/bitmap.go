package Bitmap

import (
	"fmt"
)

type BitMap struct {
	bits   [][]int
	Row    int
	Column int
}

//生成bitmap
func New(x int, y int) *BitMap {
	bitmap := make([][]int, y)

	for i := 0; i < y; i++ {
		bitmap[i] = make([]int, x)
	}
	return &BitMap{
		bits:   bitmap,
		Row:    x,
		Column: y,
	}
}

//根据作业要求进行的初始化
func (bm *BitMap) Init() {
	bitmap := bm.bits
	for i := 0; i < len(bitmap[0]); i++ {
		bitmap[0][i] = 1
	}
	bitmap[0][6] = 0
	bitmap[0][8] = 0
	bitmap[0][9] = 0
	bitmap[0][15] = 0
}

func (bm *BitMap) Add() (bool, int, int) {
	bitmap := bm.bits
	for row := 0; row < len(bitmap); row++ {
		for column := 0; column < len(bitmap[row]); column++ {
			if bitmap[row][column] == 0 {
				bitmap[row][column] = 1
				return true, row, column
			}
		}
	}
	return false, 0, 0
}

func (bm *BitMap) Remove(row int, column int) bool {
	bitmap := bm.bits
	if row >= bm.Row || row < 0 || column >= bm.Column || column < 0 || bitmap[row][column] == 0 {
		return false
	}
	bitmap[row][column] = 0
	return true
}

func (bm *BitMap) Print() {
	bitmap := bm.bits
	for i := 0; i < len(bitmap); i++ {
		fmt.Println(bitmap[i])
	}
	fmt.Print("\n\n")
}
