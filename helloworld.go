package main

/*
#include <stdio.h>
*/

import (
	"C"
	"fmt"
	"runtime"
	"unsafe"
	"math"
)

func init() {
	fmt.Printf("Map:%v\n", m)
	info = fmt.Sprintf("OS:%s, Arch:%s", runtime.GOOS, runtime.GOARCH)
}

var m map[int]string = map[int]string{1: "A", 2: "B", 3: "C"}
var info string

func main() {
	fmt.Println(info)
	fmt.Println("Hello World!")
	fn, mn, ln, nn := getName()
	_, _, _, nickName := getName()
	fmt.Print("firstName:" + fn + "\n")
	fmt.Print("middleName:" + mn + "\n")
	fmt.Print("lastName:" + ln + "\n")
	fmt.Print("nickName:" + nn + "\n")
	fmt.Print("nickName:" + nickName + "\n")
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan
	fmt.Println("Result:", sum1, sum2, sum1+sum2)

	cstr := C.CString("hello,world! Cgo")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}

func getName() (firstName, middleName, lastName, nickName string) {
	//return "May","M","Chen","Babe"
	firstName = "SQL"
	middleName = "S"
	lastName = "Li"
	nickName = "SQL Li"
	return
}

//浮点数比较
//p为用户自定义的比较精度，例如：0.00001
func IsEqual(f1, f2 ,p float64) bool {
	return math.Fdim(f1,f2) < p
}

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum
}
