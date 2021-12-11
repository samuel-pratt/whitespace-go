package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parseInt(instructions []byte) (value, offset int) {
	if len(instructions) == 2 {
		return 0, 2
	} else if len(instructions) == 1 {
		return 0, 1
	} else if len(instructions) == 0 {
		return 0, 0
	}

	// Counts how many steps to take after retun
	offset = 1

	// Determine positive or negative from fist byte
	var pos_neg int
	if instructions[0] == 32 {
		pos_neg = 1
	} else {
		pos_neg = -1
	}

	// Create string in binary until hitting N
	var binary strings.Builder
loop:
	for _, curr := range instructions[1:] {
		offset = offset + 1
		switch curr {
		case 32:
			binary.WriteString("0")
		case 10:
			break loop
		case 9:
			binary.WriteString("1")
		}
	}

	// Convert binary to base10
	decimal, err := strconv.ParseInt(binary.String(), 2, 10)
	if err != nil {
		fmt.Println(err)
	}

	return (int(decimal) * pos_neg), offset
}

func main() {
	var stack []int
	var heap = make(map[int]int)
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
			// SS = push number onto stack
			if instructions[i+1] == 32 {
				value, offset := parseInt(instructions[i+2:])
				stack = append(stack, value)
				i = i + 1 + offset
				continue
			}

			var command = string(instructions[i+1]) + string(instructions[i+2])
			switch command {
			// SNS = duplicate top element of stack
			case "\n ":
				index := len(stack) - 1
				stack = append(stack, stack[index])

				i = i + 3
				continue

			// STS = duplicate the 0 based n-th item from stack onto top of stack
			case "\t ":
				value, offset := parseInt(instructions[i+3:])
				stack = append(stack, stack[value])
				i = i + 2 + offset

				i = i + 2
				continue

			// SNT = swap top two items on stack
			case "\n\t":
				// Pop first item
				index := len(stack) - 1
				itemOne := stack[index]
				stack = stack[:index]

				// Pop second item
				index = len(stack) - 1
				itemTwo := stack[index]
				stack = stack[:index]

				stack = append(stack, itemOne)
				stack = append(stack, itemTwo)

				i = i + 2
				continue

			// SNN = discard top item on stack
			case "\n\n":
				// Pop
				index := len(stack) - 1
				stack = stack[:index]

				i = i + 2
				continue

			// STN = discard n items from top of stack while keeping top item
			case "\t\n":
				index := len(stack) - 1
				item := stack[index]
				stack = stack[:index]

				value, offset := parseInt(instructions[i+3:])
				for j := 0; j < value; j++ {
					index := len(stack) - 1
					stack = stack[:index]
				}

				stack = append(stack, item)

				i = i + 2 + offset
				continue
			}
		// N = flow control
		case 10:
			var command = string(instructions[i+1]) + string(instructions[i+2])
			switch command {
			// NSS = mark a location in the program as a subroutine with a label
			case "  ":

			// NST = call a subroutine with a given label
			case " \t":

			// NSN = jump to given label
			case " \n":

			// NTS = pop top integer off stack, if 0 jump to given label, else keep popped item removed and continue
			case "\t ":

			// NTT = pop top integer off stack and jump to given label if negative, else keep popped item removed and continue
			case "\t\t":

			// NTN = end subroutine and return to caller
			case "\t\n":

			// NNN = end program
			case "\n\n":
				os.Exit(0)
			}
		// T
		case 9:
			switch instructions[i+1] {
			// TS = Arithmetic
			case 32:
				var command = string(instructions[i+2]) + string(instructions[i+3])
				switch command {
				// TSSS = add the top two items of the stack together
				case "  ":
					indexOne := len(stack) - 1
					itemOne := stack[indexOne]

					indexTwo := len(stack) - 1
					itemTwo := stack[indexTwo]

					stack[indexOne] = itemOne + itemTwo

					i = i + 3
					continue

				// TSST = subtract the top item of stack from second item on the stack
				case " \t":
					index := len(stack) - 1
					itemOne := stack[index]

					index = len(stack) - 1
					itemTwo := stack[index]

					stack[index] = itemTwo - itemOne

					i = i + 3
					continue

				// TSSN = multiply the top two items on the stack together
				case " \n":
					indexOne := len(stack) - 1
					itemOne := stack[indexOne]

					indexTwo := len(stack) - 1
					itemTwo := stack[indexTwo]

					stack[indexOne] = itemOne * itemTwo

					i = i + 3
					continue

				// TSTS = integer division the second item on the stack by the top item on the stack
				case "\t ":
					index := len(stack) - 1
					itemOne := stack[index]

					index = len(stack) - 1
					itemTwo := stack[index]

					stack[index] = int(itemTwo / itemOne)

					i = i + 3
					continue

				// TSTT = modulo of the second item on the stack with the top item on the stack
				case "\t\t":
					index := len(stack) - 1
					itemOne := stack[index]

					index = len(stack) - 1
					itemTwo := stack[index]

					stack[index] = itemTwo % itemOne

					i = i + 3
					continue
				}
			// TN = I/O
			case 10:
				var command = string(instructions[i+2]) + string(instructions[i+3])
				switch command {
				// TNSS = pop the top integer and print as character
				case "  ":
					index := len(stack) - 1
					item := stack[index]
					stack = stack[:index]

					fmt.Print(string(byte(item)))

					i = i + 3
					continue

				// TNST = pop the top integer and print as integer
				case " \t":
					index := len(stack) - 1
					item := stack[index]
					stack = stack[:index]

					fmt.Print(item)

					i = i + 3
					continue

				// TNTS = pop the top integer, read a character from input, and save to heap with popped value as key, input as value
				case "\t ":
					index := len(stack) - 1
					item := stack[index]
					stack = stack[:index]

					reader := bufio.NewReader(os.Stdin)
					char, err := reader.ReadByte()
					if err != nil {
						fmt.Print(err)
						return
					}

					heap[item] = int(char)

					i = i + 3
					continue

				// TNTT = pop the top integer, read an integer from input, and save to heap with popped value as key, input as value
				case "\t\t":
					index := len(stack) - 1
					item := stack[index]
					stack = stack[:index]

					var integer int
					_, err := fmt.Scanf("%d", &integer)
					if err != nil {
						fmt.Print(err)
						return
					}

					heap[item] = integer

					i = i + 3
					continue
				}
			// TT = Heap Access
			case 9:
				var command = string(instructions[i+2])
				switch command {
				// TTS = pop top two items of stack, and store the top item in heap
				case " ":
					// Pop first item
					index := len(stack) - 1
					itemOne := stack[index]
					stack = stack[:index]

					// Pop second item
					index = len(stack) - 1
					itemTwo := stack[index]
					stack = stack[:index]

					heap[itemTwo] = itemOne

					i = i + 2
					continue

				// TTT = pop top item of stack, and push the item corresponding to that heap address to the top of the stack
				case "\t":
					// Pop first item
					index := len(stack) - 1
					item := stack[index]
					stack = stack[:index]

					stack = append(stack, heap[item])

					i = i + 2
					continue
				}
			}

		}
	}
	fmt.Println(stack)
	fmt.Println(heap)
}
