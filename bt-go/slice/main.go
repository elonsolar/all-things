package main

import (
	// "strings"
	"fmt"
	"math/bits"
)

// go build -gcflags=-S main.go
func main() {
	msgs :=make([]string, 0,3)

	msgs=append(msgs,"a")
	msgs=append(msgs,"b")
	msgs=append(msgs,"c")
	msgs=append(msgs,"d")

	// strings.IndexByte()
	fmt.Println(msgs)
	 var size  uint64=8
 	fmt.Println(uintptr(bits.TrailingZeros64(size)) )
 	// fmt.Println(uintptr(bits.TrailingZeros64(size)) & 63 )
}

//  1. 扩容
//      runtime.growslice(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *_type)slice
//   	oldPtr: 指向数据的指针
//      newLen: oldLen +num
//      num : 新增的元素数量
//      oldCap : 老切面的容量
//      et  : 切片元素底层类型  runtime._type
//

//   1.1 计算扩容后的容量大小
//    如果newLen  大于 两倍 的 oldCap ,   那么   新容量 newCap = newLen
//    如果newLen   小于 threadHold(256)  那么   新容量 newCap = 2 * oldCap ,也就是我们常说的2倍
//    如果newLen   大于 threadHold(256)  那么   新容量 在老容量的基础上 oldCap 不断的累加迭代 ，newCap+=(newCap+3*threadhold)/4 ,直到 大于newLen

//  1.2 计算分配后的内存
//  capmem=	 newCap * et.size ,  为了提升效率， 还判断是size=1 ,和 size 是2的平方的情况， 前者不需要做乘法，后者用移位
//  lenMem = oldLen *et.size , 计算出原来数据内存大小，用来拷贝

// 1.3 分配内存 ， 拷贝原始的 lenMen (内存大小)
//     mallocgc 分配内存
//     memmove 复制内存



//  2 分配内存


