package main

import (
	"fmt"
	"unsafe"
)

func main() {

	println(call_incr(3000))

	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	p := &arr[0][0]

	fmt.Println(*p)
	unsafepointer := unsafe.Pointer(p)
	up := uintptr(unsafepointer)
	up++
	np := *up
	fmt.Println(&up)
}

func call_incr(x int64) int64 {
	var v1 int64 = 15213
	v2 := incr(&v1, 3000)
	return x + v2
}

func incr(p *int64, val int64) int64 {
	x := *p
	y := x + val
	*p = y
	return x
}
