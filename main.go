package main

import (
	"fmt"
	"go-ds-algo/src"
)

func main() {
	var nums = [5]int{1, 3, 7, 11, 15}

	bs := src.NewBinarySearch(nums)
	fmt.Println(bs.Find(11))
}
