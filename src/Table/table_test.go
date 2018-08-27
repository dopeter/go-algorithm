package Table

import (
	"testing"
	"fmt"
)

func TestBinarySearchTable(t *testing.T) {

	bsTable:=NewBinarySearchTable()

	for loop:=0;loop<10000000;loop++{
		bsTable.Put(loop,loop)
	}

	//bsTable.Put(0,0)
	//bsTable.Put(1,1)

	bsTable.Put(len(bsTable.Keys),len(bsTable.Keys))

	val2:=bsTable.Get(3298)

	fmt.Println(val2)
}

func TestBinarySearchTree(t *testing.T) {

	/*
	如果输入的数是顺序的，则二叉会退化成无序链表
	12,11,10,9 会让树变成只有右子树，不会形成lgn的树高度
	12
		11
			10
				9

	 */

	bsTree:=NewBinarySearchTree()

	bsTree.Put(12,"12")
	bsTree.Put(3,"3")
	bsTree.Put(2,"2")
	bsTree.Put(5,"5")
	bsTree.Put(4,"4")

	bsTree.Put(11,"11")
	bsTree.Put(13,"13")

	bsTree.Put(15,"15")

	fmt.Println(bsTree.Root.Val)
	fmt.Println(bsTree.Get(5))
	fmt.Println(bsTree.Get(11))
	fmt.Println(bsTree.Root.Left.Right.Left.Val)

	//bsTree.DeleteMin()

	key3Val:=bsTree.Get(3)
	fmt.Println(key3Val)

	/*
	The tree is:
				12
		3				13
	2		5
		4		11

	delete 3 , node 4 will be set the positoin of node 3.
	 */

	bsTree.Delete(3)
	fmt.Println(bsTree.Root.Left.Val)
}
