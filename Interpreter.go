package main

import "fmt"

type ContractByteCode []byte

type Interpreter struct {
	state            *VMState
	code             *ContractByteCode
	operationMapping *OperationMapping
	gasLimit         uint64
}

func newInterpreter(code *ContractByteCode, gasLimit uint64) *Interpreter {
	return &Interpreter{state: newVM(), code: code, operationMapping: newInstructionInfo(), gasLimit: gasLimit}
}

func (interpreter *Interpreter) execute() error {

	for {

		consumedGas, pc := &interpreter.state.consumedGas, &interpreter.state.pc
		gasLimit := interpreter.gasLimit

		if *consumedGas > gasLimit {
			//outOfGasException
		}

		curInstruction := (*interpreter.code)[*pc]
		operationInfo := interpreter.operationMapping.getInstruction(curInstruction)

		if interpreter.state.Stack.Size() < operationInfo.stackArgsCount {
			//stack underflow exception
		}

		operationInfo.execute(interpreter.state, interpreter.code)

		*consumedGas += operationInfo.gasPrice

		*pc += operationInfo.pcJump

		fmt.Println(interpreter.state.toString())
		if int(*pc) >= len(*interpreter.code) {
			//
			fmt.Println("IndexOutofBound")
			return nil
		}
	}

	return nil
}
