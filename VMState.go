package main

import "strconv"

type VMState struct {
	Memory      *Memory
	Stack       *Stack
	pc          uint
	consumedGas uint64
}

func newVM() *VMState {

	return &VMState{
		Memory:      newMemory(),
		Stack:       newStack(),
		pc:          0,
		consumedGas: 0,
	}
}

func (vm *VMState) toString() string {
	return vm.Stack.toString() +
		"\n" + "PC ---->" + strconv.Itoa(int(vm.pc)) +
		"\n" + "Consumed Gas ----->" + strconv.Itoa(int(vm.consumedGas))
}
