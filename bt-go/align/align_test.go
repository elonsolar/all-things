package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
	内存对齐
	2020年从事golang 的时候就接触了这个概念，后来学习C 也同样接触这个概念，每次都花了时间学习了，但是不久就忘了。
	原因有2个，1. 只有对内存有比较高的要求时候才有用，我没有接触这样的项目，不接触自然忘记
			 2. 有一些关键的概念没有理解透
			 	2.1 内存对齐是什么？ 为何需要对齐？  怎么在实际中运用？

	一 内存对齐是什么？
		结构体在内存中的组织，相邻的结构体成员 并不是按照紧凑的排序的， 而是按照一定规则排列的，规则是：即每个类型 【相对结构体首地址的偏移量】 由根据对齐大小决定，

		对齐大小： 可以用 unsafe.AlignOf() 来算出来，或者说math.min(unsafe.sizeof(), ptrSize) 即字段大小和计算机指针标量大小 的中的最小值
				 结构体的对齐大小为所有成员中最大对齐大小的值

		对齐规则：

		第一个字段 偏移量 为0 ， 其他字段的偏移量是对齐大小的整数倍。
		注意；由于对齐大小的存在，结构体的大小不单单是所有成员的大小之和，除此之外还要加上由于对齐空出来的内存大小

	二  为何需要对齐
		网上说：1. 某些cpu 只能访问特定地址 ，这样兼容所有系统
			   2. 计算机每次按字长读取数据， 如果一个字段 的其中一个字节，被分配在了2个字节上了，那么就需要多读取一次

	三  怎么在实际中运用

		定义结构体一般是小字节字段放在一起（ 但还是要计算一下）

	更多详情：https://juejin.cn/post/7077833959047954463





*/

//相对于结构体首地址的偏移
// 对齐规则TODO

// 0000 0000 first;
// 0000 0000 second  s: 1 ,align :1
// 0000 0000 //-
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000 //-
// 0000 0000
// 0000 0000 //third start ,aligin 8
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000 //third end
// 0000 0000 // fourth
// 0000 0000 // fith start
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000 // fith start
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000
// 0000 0000

type RecordHeader struct {
	first  byte
	second byte
	third  int
	fourth byte
	fifth  string
}

func TestAligin(t *testing.T) {

	header := RecordHeader{}

	//  size: 表示内存大小， alignof: 表示对齐大小（和内存大小有关） ,offsetof :偏移首地址大小（对齐大小的整数倍）

	// output :  size:1  , align 1, offsetof: 0
	// 解释：   size: 大小 1 ，
	// 		   alginof : math.min(1,8) =1
	//         offset 由于是第一个 字段偏移  为0
	fmt.Printf("first ;size: %d alignof:%d, offsetof:%d \n", unsafe.Sizeof(header.first), unsafe.Alignof(header.first), unsafe.Offsetof(header.first))

	// output :  size:1  , align 1, offsetof: 1
	// 解释：   size: 大小 1 ，
	// 		   alginof : math.min(1,8) =1
	//         offset : 对齐大小的整数倍  m*alignof =1
	fmt.Printf("second ;size: %d alignof:%d, offsetof:%d \n", unsafe.Sizeof(header.second), unsafe.Alignof(header.second), unsafe.Offsetof(header.second))

	// output :  size:8  , align 8, offsetof: 8
	// 解释：   size: 大小 8 ，
	// 		   alginof : math.min(8,8) =8
	//         offset : 对齐大小的整数倍  m*alignof =8
	fmt.Printf("third ;size: %d alignof:%d, offsetof:%d \n", unsafe.Sizeof(header.third), unsafe.Alignof(header.third), unsafe.Offsetof(header.third))

	// output :  size:1  , align 1, offsetof:16
	// 解释：   size: 大小 1 ，
	// 		   alginof : math.min(1,8) =1
	//         offset : 对齐大小的整数倍  m*alignof =16 (上个字段int (8 byte) 占据了 8-15 个字节)
	fmt.Printf("fourth ;size: %d alignof:%d, offsetof:%d \n", unsafe.Sizeof(header.fourth), unsafe.Alignof(header.fourth), unsafe.Offsetof(header.fourth))

	// output :  size:16  , align 8, offsetof:16
	// 解释：   size: 大小 16 ，
	// 		   alginof : math.min(8,8) =8
	//         offset : 对齐大小的整数倍  m*alignof =24 (上个字段byte  占据了 偏移为 16-16，  由于对齐大小是8 整倍数，3*8=24>16 )
	fmt.Printf("fifth ;size: %d alignof:%d, offsetof:%d \n", unsafe.Sizeof(header.fifth), unsafe.Alignof(header.fifth), unsafe.Offsetof(header.fifth))

	// output :  size:40  , align 8
	// 解释： size =1[0-1) +1[1-2) +[2-7] + 8[8-15] +1[16-17) + + 16(24-40)  =40
	//		alignof: math.max(最大的成员对齐大小)=8
	fmt.Printf("header; size:%d ,alignof:%d \n", unsafe.Sizeof(header), unsafe.Alignof(header))

}
