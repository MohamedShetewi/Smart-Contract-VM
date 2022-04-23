package main

type (
	BasicOperation func(*VMState)
	PushOperation  func(*VMState, *ContractByteCode)
)

func AddOP(state *VMState) {
	stack := state.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.Add(b))
}

func SubOP(state *VMState) {
	stack := state.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.Sub(b))
}

func MulOP(state *VMState) {
	stack := state.Stack

	b := stack.Pop()
	a := stack.Pop()

	result, _ := a.Multiply(b)
	stack.Push(result)
}

// GreaterOp Return 1 if a > b
func GreaterOp(state *VMState) {
	stack := state.Stack

	b := stack.Pop()
	a := stack.Pop()

	isGreater := a.GT(b)

	var result DataWord

	if isGreater {
		result.setUint32(1, 0)
	} else {
		result.setUint32(0, 0)
	}

	stack.Push(&result)
}

func XorOP(state *VMState) {
	stack := state.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.Xor(b))
}
func AndOP(state *VMState) {
	stack := state.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.And(b))
}
func OrOP(state *VMState) {
	stack := state.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.Or(b))
}

func NotOP(state *VMState) {
	stack := state.Stack

	a := stack.Pop()

	stack.Push(a.Not())
}

func PushOp(state *VMState, code *ContractByteCode) {
	stack := state.Stack
	newData := newDataWord()
	newData.setDataWord((*code)[state.pc:])
	stack.Push(newData)
}

func PopOp(state *VMState) {
	state.Stack.Pop()
}

func MStoreOp(state *VMState) {
	mem := state.Memory

	val, index := state.Stack.Pop(), state.Stack.Pop()
	mem.store(val, index)
}

func MLoadOp(state *VMState) {
	mem := state.Memory
	index := state.Stack.Pop()
	state.Stack.Push(mem.load(index.toInt()))
}
