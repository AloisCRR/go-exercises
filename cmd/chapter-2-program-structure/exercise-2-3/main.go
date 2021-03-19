package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Println(PopCount(1000))
}

func PopCount(x uint64) int {
	var count int

	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}

	return count
}
