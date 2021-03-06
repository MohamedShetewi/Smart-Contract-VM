package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	OperationType func(*VMState, *ContractByteCode) error
)

func AddOP(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.Add(b))

	return nil
}

func SubOP(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.Sub(b))

	return nil
}

func MulOP(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack

	b := stack.Pop()
	a := stack.Pop()

	result, _ := a.Multiply(b)
	stack.Push(result)

	return nil
}

// GreaterOp Return 1 if a > b
func GreaterOp(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack

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

	return nil
}

func XorOP(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.Xor(b))

	return nil
}
func AndOP(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.And(b))

	return nil
}
func OrOP(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack

	b := stack.Pop()
	a := stack.Pop()

	stack.Push(a.Or(b))
	return nil
}

func NotOP(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack

	a := stack.Pop()

	stack.Push(a.Not())
	return nil
}

func PushOp(state *VMState, code *ContractByteCode) error {
	stack := state.Frame.Stack
	newData := newDataWord()
	newData.setDataWord((*code)[state.Frame.pc+1 : state.Frame.pc+33])
	stack.Push(newData)
	return nil

}

func PopOp(state *VMState, code *ContractByteCode) error {
	state.Frame.Stack.Pop()
	return nil

}

func MStoreOp(state *VMState, code *ContractByteCode) error {
	mem := state.Memory

	offset, value := state.Frame.Stack.Pop().toInt32(), state.Frame.Stack.Pop().toByteArray()
	mem.store(int(offset), value)
	return nil

}

func MLoadOp(state *VMState, code *ContractByteCode) error {
	mem := state.Memory
	offset := state.Frame.Stack.Pop().toInt32()
	size := state.Frame.Stack.Pop().toInt32()

	data, err := mem.load(int(offset), int(size))

	if err != nil {
		return err
	}
	x := newDataWord()
	x.setDataWord(data)
	state.Frame.Stack.Push(x)
	return nil
}

func OracleOp(state *VMState, code *ContractByteCode) error {
	keyOffset := state.Frame.Stack.Pop().toInt32()
	keySize := state.Frame.Stack.Pop().toInt32()

	urlOffset := state.Frame.Stack.Pop().toInt32()
	urlSize := state.Frame.Stack.Pop().toInt32()

	url := string(state.Memory[urlOffset : urlOffset+urlSize])
	key := string(state.Memory[keyOffset : keyOffset+keySize])

	var messageChanel chan Message
	state.Frame.localVariables = append(state.Frame.localVariables, messageChanel)
	var executionTime *time.Duration
	go func() {
		startingTime := time.Now()

		defer func(start time.Time) {
			*executionTime = time.Since(startingTime)
		}(startingTime)

		res, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
		}

		resBody, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		bodyMap := make(map[string]interface{})
		err = json.Unmarshal(resBody, &bodyMap)

		var typeOfData dataType

		if fmt.Sprintf("%T", bodyMap[key]) == "string" {
			typeOfData = String
		} else if fmt.Sprintf("%T", bodyMap[key]) == "int" {
			typeOfData = integer
		}

		var msg = Message{
			dataType: typeOfData,
			val:      fmt.Sprint(bodyMap[key]),
		}
		messageChanel <- msg
	}()
	if executionTime.Seconds() > 5 {
		return errors.New("outdated value")
	}

	return nil
}

func JumpOp(state *VMState, code *ContractByteCode) error {
	pc := &state.Frame.pc
	*pc = uint(state.Frame.Stack.Pop().toInt32())
	return nil

}

// JumpIOp conditional Jump
func JumpIOp(state *VMState, code *ContractByteCode) error {
	pc := &state.Frame.pc
	check := state.Frame.Stack.Pop().toInt32()
	nextInstruction := state.Frame.Stack.Pop().toInt32()

	if check == 1 {
		*pc = uint(nextInstruction)
	} else {
		*pc++
	}
	return nil
}
