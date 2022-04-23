package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func geneterateOpcode() {
	file, err := os.Open("./Opcodes")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	for i := 1; scanner.Scan(); i++ {
		opCode := strings.Split(scanner.Text(), " ")
		hexa := fmt.Sprintf("%x", i)
		if len(hexa) == 1 {
			hexa = "0" + hexa
		}
		fmt.Println(opCode[0] + "    OPCODE = 0x" + hexa)
	}
}

func main() {

}
