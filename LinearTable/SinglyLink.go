package main

import "fmt"

//数据结构   线性表 之  单链表


type Elem int

type LinkNode struct {
	Data Elem
	Next *LinkNode
}

//初始化头节点 （头节点中存储链表的长度）
func lNew() *LinkNode {
	return &LinkNode{0,nil}
}

//在链表的i前插入元素(不包含头节点)
func (head *LinkNode) Insert(i int,e Elem) bool {
	p := head
	index := 1
	for p != nil && index < i{
		p = p.Next
		index++
	}
	if p == nil || index > i {
		fmt.Print("输入的i不符合要求",i)
		return false
	}
	val := &LinkNode{Data:e}
	val.Next = p.Next
	p.Next = val
	head.Data++
	return true
}

//遍历链表
func (head *LinkNode) Traverse()  {
	p := head.Next
	for p != nil{
		fmt.Printf("%v\t",p.Data)
		p = p.Next
	}
	fmt.Println()
	fmt.Println("完成")
}

//删除链表中的第i个节点
func (head *LinkNode) Del(i int) bool {
	p := head
	index := 1
	for p != nil && index < i{
		p = p.Next
		index++
	}
	if p == nil || index > i {
		fmt.Print("输入的i不符合要求",i)
		return false
	}
	p.Next = p.Next.Next
	head.Data--
	return true
}

//获取第i个元素的值
func (head *LinkNode) GetElem(i int) Elem {
	p := head.Next
	j := 1
	for ;j < i; j++ {
		if p == nil {
			fmt.Println("输入的i不符合要求",i)
			return -1
		}
		p = p.Next
	}
	if i < j || p == nil {
		fmt.Println("输入的i不符合要求",i)
		return -1
	}

	return p.Data
}

//测试函数
func main() {
	link := lNew()
	link.Insert(1,1)
	link.Insert(2,2)
	link.Insert(3,3)
	link.Insert(2,4)

	link.Traverse()
	fmt.Println(link.Data)

	fmt.Println(link.GetElem(3))

	link.Del(2)
	link.Traverse()
	fmt.Println(link.Data)
}
