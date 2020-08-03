package main // 包声明
// go 是一个天生支持多核计算的云开发时代C语言
// node + go 后端语言 跟服务器有关 
// require('fs') fileSystem 
import (
	"fmt" // 向命令行输出hello world  format
	"math"
)

// es6 
// node 脚本语言
// go  c 二进制文件
// 编译器  编译原理   func => function
func minEatingSpeed(piles []int, H int) int {
	lo := 1 // let lo = 1
  hi := maxPiles(piles)
	//  最大pile的数量是 11  交给一个函数
	// 是否可以吃完   吃香蕉的速度范围 1-11
	// 函数 是 组织代码的最小单元 {块级作用域}
	for {
		if lo > hi {
			break
		}
		if canEatAllBananas(piles, H, lo) {
			return lo
		} else { 
			lo++ 
		}
		
	}

	return lo
}

func canEatAllBananas(piles []int ,H int, k int) bool {
	timer := 0
	for _, n := range piles{
		timer = timer + int(math.Ceil( float64(n) / float64(k) ))
	}
	if timer <= H {
		return true
	}
	return false
}

func maxPiles(piles []int) int {
	h := 0
	// for i := 0; i < len(piles); i++ {
	// 	 if h < piles[i]{
	// 		  h = piles[i]
	// 	 }
	// }
	for _, n := range piles{
		h = max(h, n)
	}
	return h
}
func max(a int, b int) int {
	// return a > b ? a : b
	if a > b {
		return a
	}
	return b
}

func main() { // func 函数关键字
	// 从main 开始
	// go 是一个静态的， 存在类型约束
	fmt.Printf("%d小时内吃完香蕉的最慢速度是%d只/小时",8 ,minEatingSpeed([]int{3, 6, 7, 11}, 8));
}