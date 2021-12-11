package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var args = os.Args[1:]

	if len(args) != 1 {
		fmt.Print("Missing file argument")
		return
	}

	file, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Print(err)
		return
	}

	var instructions []byte

	for i := 0; i < len(file); i++ {
		if file[i] == 32 || file[i] == 10 || file[i] == 9 {
			instructions = append(instructions, file[i])
		}
	}

	for i := 0; i < len(instructions); i++ {
		switch instructions[i] {
		// S = stack manipulation
		case 32:
			fmt.Print("Space ")
			// SS = push number onto stack
			// SNS = duplicate top element of stack
			// STS = duplicate the 0 based n-th item from stack
			// SNT = swap top two items on stack
			// SNN = discard top item on stack
			// STN = discard n items from top of stack while keeping top item
		// N = flow control
		case 10:
			fmt.Print("Newline ")
			// NSS = mark a location in the program as a subroutine with a label
			// NST = call a subroutine with a given label
			// NSN = jump to given label
			// NTS = pop top integer off stack, if 0 jump to given label, else keep popped item removed and continue
			// NTT = pop top integer off stack and jump to given label if negative, else keep popped item removed and continue
			// NTN = end subroutine and return to caller
			// NNN = end program
		// T
		case 9:
			switch instructions[i+1] {
			// TS = Arithmetic
			case 32:
				fmt.Print("TabSpace ")
				// TSSS = add the top two items of the stack together
				// TSST = subtract the top item of stack from second item on the stack
				// TSSN = multiply the top two items on the stack together
				// TSTS = integer division the second item on the stack by the top item on the stack
				// TSTT = modulo of the second item on the stack with the top item on the stack
			// TT = Heap Access
			case 10:
				fmt.Print("TabTab ")
				// TTS = pop top two items of stack, and store the top item in
				// TTT = pop top item of stack, and push the item corresponding to that heap address to the top of the stack
			// TN = I/O
			case 9:
				fmt.Print("TabNewline ")
				// TNSS = pop the top integer and print as character
				// TNST = pop the top integer and print as integer
				// TNTS = pop the top integer, read a character from input, and save to heap with popped value as key, input as value
				// TNTT = pop the top integer, read an integer from input, and save to heap with popped value as key, input as value
			}
		}
	}
}
