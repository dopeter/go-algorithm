package Table

type BinarySearchTable struct {

	Keys []int
	Vals []interface{}


}

func NewBinarySearchTable() *BinarySearchTable{
	return &BinarySearchTable{Keys:make([]int,0,10),Vals:make([]interface{},0,10)}
}



func(this *BinarySearchTable) Get(key int) interface{}{
	index:=this.Rank(key);

	if index< len(this.Keys) && this.Keys[index] == key{
		return this.Vals[index]
	}else{
		return nil
	}
}

func(this *BinarySearchTable) Rank(key int) int{

	lowIndex:=0
	highIndex:=len(this.Keys)-1

	for{
		if lowIndex>highIndex{break}

		midKey:=lowIndex+(highIndex-lowIndex)/2

		if key<midKey{
			highIndex=midKey-1
		}else if key>midKey{
			lowIndex=midKey+1
		}else{
			return midKey
		}
	}

	return lowIndex
}

func(this *BinarySearchTable) Put(key int,val interface{}) {

	index:=this.Rank(key)

	if index< len(this.Keys) && this.Keys[index] == key{
		this.Vals[index]=val
		return
	}

	this.Keys=append(this.Keys,0)
	this.Vals=append(this.Vals,nil)

	//leftKeyRank:=this.Keys[0:index]
	//rightKeyRank:=this.Keys[index:]
	//leftKeyRank=append(leftKeyRank,key)
	//this.Keys=append(leftKeyRank,rightKeyRank...)
	//
	//leftValRank:=this.Vals[0:index]
	//rightValRank:=this.Vals[index:]
	//leftValRank=append(leftValRank,val)
	//this.Vals=append(leftValRank,rightValRank...)

	for curIndex:= len(this.Keys)-1 ; curIndex>index ; curIndex--{
		this.Keys[curIndex]=this.Keys[curIndex-1]
		this.Vals[curIndex]=this.Vals[curIndex-1]
	}

	this.Keys[index]=key
	this.Vals[index]=val

}








