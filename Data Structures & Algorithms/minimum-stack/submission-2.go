type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	stack := make([]int, 0)
	minStack := make([]int, 0)
	return MinStack{stack : stack, minStack : minStack}
}

func (this *MinStack) Push(val int) {
	minVal := val
	if len(this.minStack) > 0{
		minVal = min(this.minStack[len(this.minStack)-1], minVal) 
	}
	this.minStack = append(this.minStack, minVal)
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
