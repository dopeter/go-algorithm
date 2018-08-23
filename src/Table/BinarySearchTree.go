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
		return ""
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

func(this *BinarySearchTree) Put(key int,val string){
	this.InsertNode(this.Root,key,val)
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

	n.Size=n.Left.Size+n.Right.Size

	//This node is not new node that inserted.It's the parent node of inserted node.
	return n
}