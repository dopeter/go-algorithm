package PkgSort

/*
The first index of array is not useable.
Because its value is 0, the expression that "index/2" can't reach this level.
So the actuall value of first index is 1.
 */
type StackArray struct {
	IntArray []int
}

func NewStackArray(arrayLength int)*StackArray{
	return &StackArray{IntArray:make([]int,arrayLength,10)}
}

func NewStackArray2(array []int)*StackArray{
	return &StackArray{IntArray:array}
}

func(stackArray *StackArray) Less (index1,index2 int) bool{
	return stackArray.IntArray[index1]<stackArray.IntArray[index2]
}


func(stackArray *StackArray) Exchange(index1,index2 int){
	tmp:=stackArray.IntArray[index1]

	stackArray.IntArray[index1]=stackArray.IntArray[index2]
	stackArray.IntArray[index2]=tmp

}
/*
Exchange element from bottom to top.
 */
func(stackArray *StackArray) BottomToTop(index int){
	for{
		// here index>1 , see the first index description.
		if index>1 && stackArray.Less(index/2,index){
			stackArray.Exchange(index/2,index)

			//to top level
			index=index/2
		}else{
			break
		}

	}
}

func(stackArray *StackArray) TopToBottom(index int){
	for{
		// index out of range.
		if index*2 > len(stackArray.IntArray)-1{
			break
		}

		//go to the downstair level
		nextLevelIndex:=index*2

		//every level has two elements, select the bigger one
		if nextLevelIndex<len(stackArray.IntArray)-1 && stackArray.Less(nextLevelIndex,nextLevelIndex+1){
			nextLevelIndex+=1
		}

		if stackArray.Less(index,nextLevelIndex)==false{
			break;
		}

		stackArray.Exchange(index,nextLevelIndex)
		index=nextLevelIndex
	}
}

/*
Put the new element to the bottom of array.
 */
func(stackArray *StackArray) InsertElem (val int){

	stackArray.IntArray=append(stackArray.IntArray,val)
	stackArray.BottomToTop(len(stackArray.IntArray)-1)
}

func(stackArray *StackArray) RemoveMax() int{
	maxVal:=stackArray.IntArray[1]

	lastIndex:=len(stackArray.IntArray)-1

	stackArray.Exchange(1,lastIndex)

	//In golang , sub slice,delete lastindex
	stackArray.IntArray=stackArray.IntArray[0:lastIndex]

	stackArray.TopToBottom(1)

	return maxVal
}



