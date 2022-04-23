package main

type OperationInfo struct {
	basicOperation BasicOperation
	pushOperation  PushOperation
	stackArgsCount uint // number of arguments needed for the operation
	gasPrice       uint
	pcJump         uint
}

const (
	lowGasPrice  = 2
	midGasPrice  = 4
	highGasPrice = 7
)

const (
	onePCJump = 1
)

type InstructionMap *[100]OperationInfo

func newInstructionInfo() (oppArray *InstructionMap) {

	(*oppArray)[ADD] =
		OperationInfo{
			basicOperation: AddOP,
			stackArgsCount: 2,
			gasPrice:       lowGasPrice,
			pcJump:         onePCJump,
		}
	(*oppArray)[SUB] = OperationInfo{
		basicOperation: SubOP,
		stackArgsCount: 2,
		gasPrice:       lowGasPrice,
		pcJump:         onePCJump,
	}
	(*oppArray)[MUL] = OperationInfo{
		basicOperation: MulOP,
		stackArgsCount: 2,
		gasPrice:       midGasPrice,
		pcJump:         onePCJump,
	}
	(*oppArray)[DIV] = OperationInfo{}
	(*oppArray)[GT] = OperationInfo{
		basicOperation: GreaterOp,
		stackArgsCount: 2,
		gasPrice:       lowGasPrice,
		pcJump:         onePCJump,
	}
	(*oppArray)[OR] = OperationInfo{
		basicOperation: OrOP,
		stackArgsCount: 2,
		gasPrice:       lowGasPrice,
		pcJump:         onePCJump,
	}
	(*oppArray)[XOR] = OperationInfo{
		basicOperation: XorOP,
		stackArgsCount: 2,
		gasPrice:       lowGasPrice,
		pcJump:         onePCJump,
	}
	(*oppArray)[AND] = OperationInfo{
		basicOperation: AndOP,
		stackArgsCount: 2,
		gasPrice:       lowGasPrice,
		pcJump:         onePCJump,
	}
	(*oppArray)[NOT] = OperationInfo{
		basicOperation: NotOP,
		stackArgsCount: 2,
		gasPrice:       lowGasPrice,
		pcJump:         onePCJump,
	}
	(*oppArray)[PUSH] = OperationInfo{
		pushOperation:  PushOp,
		stackArgsCount: 0,
		gasPrice:       lowGasPrice,
		pcJump:         33,
	}
	(*oppArray)[POP] = OperationInfo{
		basicOperation: PopOp,
		stackArgsCount: 0,
		gasPrice:       lowGasPrice,
		pcJump:         1,
	}

	(*oppArray)[MSTORE] = OperationInfo{
		basicOperation: MStoreOp,
		stackArgsCount: 2,
		gasPrice:       midGasPrice,
		pcJump:         1,
	}
	(*oppArray)[MLOAD] = OperationInfo{
		basicOperation: MLoadOp,
		stackArgsCount: 1,
		gasPrice:       midGasPrice,
		pcJump:         onePCJump,
	}
	return
}
