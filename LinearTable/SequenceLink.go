package main

import (
	"fmt"
)

//数据结构，线性表之   顺序表


// 定义线性表中的数据类型
type Elem int

type SqList struct {
	maxsize int
	length int
	data []Elem
}

//初始化
func SNew(maxsize int) *SqList {
	return &SqList{
		maxsize:maxsize,
		data:make([]Elem,maxsize),
	}
}

//判断线性表是否为空
func (slist *SqList) IsEmpty() bool {
	return slist.length == 0
}

//判断线性表是否已满
func (slist *SqList) IsFull() bool {
	return slist.length == slist.maxsize
}

//在第i个元素前插入e
func (slist *SqList) Insert(i int, e Elem) bool {
	if i < 1 || i >slist.length {
		fmt.Println("输入的i不符合要求：",i)
		return false
	}
	if slist.IsFull() {
		fmt.Println("该线性表已满")
		return false
	}
	for k := slist.length; k > i-1; k-- {
		slist.data[k] = slist.data[k-1]
	}
	slist.data[i-1] = e
	slist.length++
	return true
}

//删除第i个元素
func (slist *SqList) Del(i int) bool {
	if i < 1 || i >slist.length {
		fmt.Println("输入的i不符合要求：",i)
		return false
	}
	for k := i-1; k < slist.length -1;  k++ {
		slist.data[k] = slist.data[k+1]
	}
	slist.data[slist.length-1] = 0
	slist.length--
	return true
}

//获取位置i的值
func (slist *SqList) GetElem(i int) Elem {
	if i < 1 || i >slist.length {
		fmt.Println("输入的i不符合要求：",i)
		return -1
	}
	return slist.data[i-1]
}

//追加一个元素
func (slist *SqList) Sappend(e Elem) bool {
	if slist.IsFull() {
		fmt.Println("该线性表已满")
		return false
	}
	slist.data[slist.length] = e
	slist.length++
	return true
}


//测试函数
func main() {
	sq := SNew(10)

	sq.Sappend(1)
	sq.Sappend(2)
	sq.Sappend(3)
	sq.Sappend(4)
	sq.Sappend(5)
	sq.Sappend(6)
	sq.Sappend(7)
	sq.Sappend(8)
	sq.Sappend(9)

	sq.Insert(6,60)
	fmt.Println(sq)

	sq.Del(7)
	fmt.Println(sq)

	fmt.Println(sq.GetElem(6))

}
