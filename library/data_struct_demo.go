package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	size int
	number string
}

type ItemHeap []Item  // 定义一个类型

func (h ItemHeap) Len() int { return len(h) }  // 绑定len方法,返回长度
func (h ItemHeap) Less(i, j int) bool {  // 绑定less方法
	return h[i].size < h[j].size  // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h ItemHeap) Swap(i, j int) {  // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *ItemHeap) Pop() interface{} {  // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *ItemHeap) Push(x interface{}) {  // 绑定push方法，插入新元素
	*h = append(*h, x.(Item))
}

func HeapDemo() {
	item1 := Item{size: 1,number: "a"}
	item2 := Item{size: 1,number: "b"}
	item3 := Item{ size: 1, number: "c" }
	item4 := Item{size: 1, number: "d"}
	item5 := Item{size: 1,number: "e"}
	item6 := Item{ size: 1, number: "f"}
	h := &ItemHeap{item1, item2, item3, item4, item5, item6}  // 创建slice
	heap.Init(h)  // 初始化heap
	for len(*h) > 0 { // 排序输出
		fmt.Printf("%v\n", heap.Pop(h))
	}
}
