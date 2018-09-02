package Table

type BRNode struct{
	Key int
	Val string
	Left *BRNode
	Right *BRNode
	Size int
	IsRed bool
}

func NewBRNode(key int , val string , size int , isRed bool) *BRNode{
	return &BRNode{Key:key,Val:val,Size:size,IsRed:isRed}
}

type BRTree struct {
	Root *BRNode
}

func(this *BRTree) CalcucateSize(n *BRNode){
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

func (this *BRTree) RotateLeft(h *BRNode) *BRNode{

	x:=h.Right
	h.Right=x.Left
	x.Left=h
	x.IsRed=h.IsRed
	h.IsRed=true

	x.Size=h.Size
	this.CalcucateSize(h)

	return x
}

func (this *BRTree) RotateRight(h *BRNode) *BRNode{

	x:=h.Left
	h.Left=x.Right
	x.Right=h
	x.IsRed=h.IsRed
	h.IsRed=true

	x.Size=h.Size
	this.CalcucateSize(h)

	return x
}

func (this *BRTree) ValidRed(h *BRNode) bool{
	if h ==nil{
		return false
	}

	return h.IsRed==true
}

func(this *BRTree) CompareKey(key1 int,key2 int) int{
	if key1<key2{
		return -1
	}else if key1==key2{
		return 0
	}else{
		return 1
	}
}

func (this *BRTree) FlipColors(h *BRNode){
	h.IsRed=true
	if h.Left!=nil{
		h.Left.IsRed=false
	}
	if h.Right!=nil{
		h.Right.IsRed=false
	}
}


func (this *BRTree) Put(key int , val string){

	this.Root=this.put(this.Root,key,val)
	this.Root.IsRed=false

}

func (this *BRTree) put(h *BRNode,key int , val string) *BRNode{

	if h == nil{
		return NewBRNode(key,val,1,true)
	}

	compareRes:=this.CompareKey(key,h.Key)

	if compareRes < 0{
		h.Left=this.put(h.Left,key,val)
	}else if compareRes > 0{
		h.Right=this.put(h.Right,key,val)
	}else{
		h.Val=val
	}

	//red black node valid
	if this.ValidRed(h.Right) && this.ValidRed(h.Left)==false{
		h=this.RotateLeft(h)
	}

	if this.ValidRed(h.Left) && this.ValidRed(h.Left.Left) {
		h=this.RotateRight(h)
	}

	if this.ValidRed(h.Left) && this.ValidRed(h.Right){
		this.FlipColors(h)
	}

	this.CalcucateSize(h)

	return h
}





