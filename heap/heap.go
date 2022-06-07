package heap

type node struct{
	Value int
	Priority int
}

type Heap struct {
	heap []node
	Less func(i,j int)bool
}

func Constructor(heap []node, Less func(i,j int)bool)Heap{
	H := Heap{heap,Less}
	if len(heap) <= 1{
		return H
	}

	end := len(heap)
	for i := end/2 - 1 ; i >= 0; i--{
		H.Adjust2Down(i)
	}
	return H
}

func (this *Heap)Adjust2Down(id int){
	if id > len(this.heap)/2-1 || len(this.heap) <= 1{
		return
	}
	left := 2*id + 1
	right := 2*id + 2
	end := len(this.heap)
	high := id

	if this.Less(left,high){
		high = left
	}
	if right < end &&  this.Less(right,high){
		high = right
	}
	if high == id{
		return
	}
	this.heap[id], this.heap[high] = this.heap[high], this.heap[id]
	this.Adjust2Down(high)
}

func (this *Heap)Adjust2Up(id int){
	if id == 0 || len(this.heap) <= 1{
		return
	}
	high := id
	//如果Id为大于0的偶数，则存在左兄弟节点
	if id % 2 == 0 && id > 0{
		left := id - 1
		if this.Less(left,high){
			high = left
		}
	}
	f := (id - 1)/2
	if this.Less(f,high){
		high = f
	}

	if high == f{
		return
	}
	this.heap[f], this.heap[high] = this.heap[high], this.heap[f]
	this.Adjust2Up(high)
}

func (this *Heap)Top()node{
	if len(this.heap) <= 0{
		return node{}
	}
	return this.heap[0]
}
