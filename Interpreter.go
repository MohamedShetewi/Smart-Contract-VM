package main

type ContractByteCode []byte

type Interpreter struct {
	state           VMState
	code            ContractByteCode
	instructionInfo *InstructionMap
}

func newInterpreter(state VMState, code ContractByteCode) *Interpreter {
	return &Interpreter{state: state, code: code, instructionInfo: newInstructionInfo()}
}

func (interpreter *Interpreter) execute() error {

	return nil
}
