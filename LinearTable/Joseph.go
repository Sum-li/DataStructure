package main

import "fmt"

/*
	约瑟夫环问题---循环链表实现
*/

type cnode struct {
	data int //具体问题就直接将类型写死了
	Next *cnode
}

type clink struct {
	count int //当前所含的节点数
	node *cnode //当前节点的前驱节点指针
}

//初始化
func cNew(i,th int) *clink {
	head := &cnode{1,nil}
	p := head
	q := p
	for j := 2; j <= i ; j++ {
		e := &cnode{j,nil}		//默认赋值为序号值，可自行更改
		p.Next = e
		p = e
		if j == th-1 {
			q = p
		}
	}
	if th == 1 {
		q = p
	}
	p.Next = head    //产生循环链表
	return &clink{i,q}   //存的是尾节点指针
}

func (link *clink) Cdel(i int)  {
	for j := 1; j < i; j++ {
		link.node = link.node.Next
	}
	fmt.Printf("%v\t",link.node.Next.data)
	link.node.Next = link.node.Next.Next
	link.count--
}

func joseph(count,th,number int)  {
	cl := cNew(count,th)
	for cl.count > 0 {
		cl.Cdel(number)
	}
	fmt.Println()
	fmt.Println("完成")
}

func main() {
	joseph(10,2,3)
}

