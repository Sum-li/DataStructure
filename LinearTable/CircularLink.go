package main

/*

	循环链表和单链表差不多，就是将单链表的尾节点指向头节点，即形成循环链表

	作用： 可以通过循环找到当前节点的前驱节点（不一定是直接前驱）

	注意： 循环链表须存储节点个数，防止插入、删除或读取等操作时 超出节点数时，count MOD i 出错

*/


// *************************************循环链表实战：约瑟夫环问题*****************************************************