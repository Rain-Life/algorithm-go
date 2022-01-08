package main

import (
	"algorithm-go/BackTracking"
	"fmt"
)

var path []int
var Aes [][]int
func main() {
	//BackTracking.BackTracking(1,1, 1)
	//fmt.Println(BackTracking.Res)

	nums := []int{1,2,3}
	fmt.Println(BackTracking.Permute(nums))
	//path = append(path, 1)
	//path = append(path, 2)
	//fmt.Printf("---path1---\n")
	//fmt.Println(path)
	//Aes = append(Aes, path)
	//fmt.Printf("---Aes1---\n")
	//fmt.Println(Aes)
	//
	//path = path[:len(path)-1]
	//fmt.Printf("---path2---\n")
	//fmt.Println(path)
	//path = append(path, 3)
	//fmt.Println(path)
	//
	//Aes = append(Aes, path)
	//fmt.Println(Aes)
	//fmt.Printf("---Aes2---\n")

}
