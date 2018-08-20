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
