package main

type VMState struct {
	Memory       *Memory
	Stack        *Stack
	pc           uint64
	remainingGas uint64
}

func (state *VMState) newVM(gas uint64) *VMState {

	return &VMState{
		Memory:       newMemory(),
		Stack:        newStack(),
		pc:           0,
		remainingGas: gas,
	}
}
