package Table

type Node struct{
	Key int
	Val string
	Left *Node
	Right *Node
	Size int
}

func NewNode(key int,val string,size int) *Node{
	return &Node{Key:key,Val:val,Size:size}
}

type BinarySearchTree struct{
	Root *Node
}

func NewBinarySearchTree() *BinarySearchTree{
	return &BinarySearchTree{}
}

func(this *BinarySearchTree) Get(key int) string{
	return this.RetrieveFromRoot(this.Root,key)
}

func(this *BinarySearchTree) CompareKey(key1 int,key2 int) int{
	if key1<key2{
		return -1
	}else if key1==key2{
		return 0
	}else{
		return 1
	}
}

func(this *BinarySearchTree) RetrieveFromRoot(n *Node,key int) string{

	if n==nil{
		return "Not exist in tree."
	}

	compareResult:=this.CompareKey(key,n.Key)

	if compareResult<0{
		return this.RetrieveFromRoot(n.Left,key)
	}else if compareResult>0{
		return this.RetrieveFromRoot(n.Right,key)
	}else{
		return n.Val
	}

}

func(this *BinarySearchTree) CalcucateSize(n *Node){
	size:=0

	if(n.Left!=nil){
		size+=n.Left.Size
	}

	if(n.Right!=nil){
		size+=n.Right.Size
	}

	if size==0{
		n.Size=1
	}else{
		n.Size=size+1
	}
}

func(this *BinarySearchTree) Put(key int,val string){

	if(this.Root==nil){
		this.Root=NewNode(key,val,1)
		return
	}

	this.Root=this.InsertNode(this.Root,key,val)

}

func(this *BinarySearchTree) InsertNode(n *Node,key int,val string) *Node{

	//here will add new node at bottom level.
	if n==nil{
		return NewNode(key,val,1)
	}

	compareResult:=this.CompareKey(key,n.Key)

	if compareResult<0{
		n.Left= this.InsertNode(n.Left,key,val)
	}else if compareResult>0{
		n.Right= this.InsertNode(n.Right,key,val)
	}else{
		//If find same key ,update val
		n.Val=val
	}

	this.CalcucateSize(n)

	//This node is not new node that inserted.It's the parent node of inserted node.
	return n
}

func(this *BinarySearchTree) DeleteMin(){
	this.DeleteMinByNode(this.Root)
}

func(this *BinarySearchTree) DeleteMinByNode(n *Node) *Node {

	//最小节点的左子节点一定为nil，直接返回右节点，如果是右子节点是nil则不管；如果不为空，它也一定小于要删除节点的父节点，所以
	//直接嫁接在父节点的左子节点上即可
	if(n.Left==nil){
		return n.Right
	}

	//这里n是要删除节点的父节点，因为要删除的节点小于它，所以直接把要删除节点的右子节点嫁接在该父节点的左子节点下即可
	//此处是递归，所以最后一定会递归到第一个传入的节点
	n.Left=this.DeleteMinByNode(n.Left)

	this.CalcucateSize(n)

	return n
}

func(this *BinarySearchTree) Delete(key int){
	this.Root=this.DeleteByKey(this.Root,key)
}

func(this *BinarySearchTree) Min(n *Node) *Node{
	if n.Left==nil{
		return n
	}

	return this.Min(n.Left)
}


func(this *BinarySearchTree) DeleteByKey(n *Node,key int) *Node{
	if n==nil{
		return nil
	}

	compareResult:=this.CompareKey(key,n.Key)

	if compareResult<0{
		n.Left=this.DeleteByKey(n.Left,key)
	}else if compareResult>0{
		n.Right=this.DeleteByKey(n.Right,key)
	}else{
		//找到了对应key的Node
		//3种情况
		//1 右节点为空，直接返回左节点，嫁接在父节点的左节点即可
		//2 左节点为空，直接返回右节点，嫁接在父节点的右节点即可
		//3 左右节点都不为空，找到右节点下面的最小节点，

		if n.Right==nil{
			return n.Left
		}

		if n.Left==nil{
			return n.Right
		}

		tmpNode:=n
		//找到要删除节点的所有子节点中最小的节点
		n=this.Min(tmpNode.Right)

		//将要删除节点的右子节点下最小节点删除，并整理完毕后，返回右子节点，实际上删除的就是n这个节点
		//因为是最小节点，所以直接可以嫁接至右子节点
		n.Right=this.DeleteMinByNode(tmpNode.Right)

		//因为n是从tmpNode的右子节点中查找出的最小节点，但一定大于tmpNode，所以直接将tmpNode的左子节点嫁接给它即可
		n.Left=tmpNode.Left

	}

	this.CalcucateSize(n)

	return n
}




