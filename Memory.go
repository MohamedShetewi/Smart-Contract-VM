package main

import "strings"

type Memory []*DataWord

func newMemory() *Memory {
	return &Memory{}
}

func (memory *Memory) store(dataWord *DataWord, index *DataWord) {
	(*memory)[index.toInt()] = dataWord
}

func (memory *Memory) load(index uint32) *DataWord {
	return (*memory)[index]
}

func (memory *Memory) toString() string {
	str := "Stack -----> \n"
	spaceCount := len(str)

	for i := 0; i < 0; i-- {
		val := (*memory)[i]
		str += strings.Repeat(" ", spaceCount)
		str += "0x" + val.toStringHex() + "\n"
	}
	return str
}
