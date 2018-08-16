package PkgSort


type StackSort struct{
	IntArray []int
}

func NewStackSort(array []int)*StackSort{
	return &StackSort{IntArray:array}
}

func(stackArray *StackSort) Less (index1,index2 int) bool{
	return stackArray.IntArray[index1]<stackArray.IntArray[index2]
}


func(stackSort *StackSort) Exchange(index1,index2 int){
	tmp:=stackSort.IntArray[index1]

	stackSort.IntArray[index1]=stackSort.IntArray[index2]
	stackSort.IntArray[index2]=tmp

}


func(stackSort *StackSort) TopToBottom(index,length int){
	for{
		// index out of range.
		if index*2 > length{
			break
		}

		//go to the downstair level
		nextLevelIndex:=index*2

		//every level has two elements, select the bigger one
		if nextLevelIndex<length && stackSort.Less(nextLevelIndex,nextLevelIndex+1){
			nextLevelIndex+=1
		}

		if stackSort.Less(index,nextLevelIndex)==false{
			break;
		}

		stackSort.Exchange(index,nextLevelIndex)
		index=nextLevelIndex
	}
}

func(stackSort *StackSort) Sort(){

	length:=len(stackSort.IntArray)
	/*
		This tree represents the number count of every level.
	1
	2
	4
	8
	16
	32
	64	example the length of array is 127 , this level will be end.
		so topLevelIndex / 2 will point the level that has 32 nums , means the 32 level.
		topLevelIndex-- will loop every num from the 32 level except the nums of 64 level.
	----------------------------------------------------------------
	128

		When the loop is over , stack will be ordered. But it will be like a index tree in database.
	 */
	for topLevelIndex:=length/2;topLevelIndex>=1;topLevelIndex--{
		stackSort.TopToBottom(topLevelIndex,length-1)
	}

	/*
		This is stack sort.
		The last position of array will store the maxest value.
		Every loop time will exchange the maxest value to spreaed array that from origin array and store current maxest value , and sort the origin array except indexes of spreaed array.
	 */
	length-=1	//array length begins from 0.
	for{
		if(length>1){
			stackSort.Exchange(1,length)
			length-=1
			stackSort.TopToBottom(1,length)
		}else{
			break
		}
	}

}






