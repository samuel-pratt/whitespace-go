package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parseInt(instructions []byte) (value, test int) {
	var offset int = 1

	// Determine positive or negative from fist byte
	var odd_even int
	if instructions[0] == 32 {
		odd_even = 1
	} else {
		odd_even = -1
	}

	// Create string in binary until hitting N
	var binary strings.Builder
	for _, curr := range instructions[1:] {
		offset = offset + 1
		switch curr {
		case 32:
			binary.WriteString("0")
		case 10:
			break
		case 9:
			binary.WriteString("1")
		}
	}

	decimal, err := strconv.ParseInt(binary.String(), 2, 10)
	if err != nil {
		fmt.Println(err)
	}

	return (int(decimal) * odd_even), offset
}

func main() {
	var stack []int
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
			if instructions[i+1] == 32 {
				// SS = push number onto stack
				value, offset := parseInt(instructions[i+2:])
				stack = append(stack, value)
				i = i + 2 + offset
				continue
			}

			var command = string(instructions[i+1]) + string(instructions[i+2])
			switch command {
			// SNS = duplicate top element of stack onto second element
			case "NS":
			// STS = duplicate the 0 based n-th item from stack onto top of stack
			case "TS":
			// SNT = swap top two items on stack
			case "NT":
			// SNN = discard top item on stack
			case "NN":
			// STN = discard n items from top of stack while keeping top item
			case "TN":

			}
		// N = flow control
		case 10:
			var command = string(instructions[i+1]) + string(instructions[i+2])
			switch command {
			// NSS = mark a location in the program as a subroutine with a label
			case "SS":
			// NST = call a subroutine with a given label
			case "ST":
			// NSN = jump to given label
			case "SN":
			// NTS = pop top integer off stack, if 0 jump to given label, else keep popped item removed and continue
			case "TS":
			// NTT = pop top integer off stack and jump to given label if negative, else keep popped item removed and continue
			case "TT":
			// NTN = end subroutine and return to caller
			case "TN":
			// NNN = end program
			case "NN":
			}
		// T
		case 9:
			switch instructions[i+1] {
			// TS = Arithmetic
			case 32:
				var command = string(instructions[i+2]) + string(instructions[i+3])
				switch command {
				// TSSS = add the top two items of the stack together
				case "SS":
				// TSST = subtract the top item of stack from second item on the stack
				case "ST":
				// TSSN = multiply the top two items on the stack together
				case "SN":
				// TSTS = integer division the second item on the stack by the top item on the stack
				case "TS":
				// TSTT = modulo of the second item on the stack with the top item on the stack
				case "TT":
				}
			// TT = Heap Access
			case 10:
				var command = string(instructions[i+2])
				switch command {
				// TTS = pop top two items of stack, and store the top item in
				case "S":
				// TTT = pop top item of stack, and push the item corresponding to that heap address to the top of the stack
				case "T":
				}
			// TN = I/O
			case 9:
				var command = string(instructions[i+2]) + string(instructions[i+3])
				switch command {
				// TNSS = pop the top integer and print as character
				case "SS":
				// TNST = pop the top integer and print as integer
				case "ST":
				// TNTS = pop the top integer, read a character from input, and save to heap with popped value as key, input as value
				case "TS":
				// TNTT = pop the top integer, read an integer from input, and save to heap with popped value as key, input as value
				case "TT":
				}
			}
		}
	}
	fmt.Print(stack)
}
