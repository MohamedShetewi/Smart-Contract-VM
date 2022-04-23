package main

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
