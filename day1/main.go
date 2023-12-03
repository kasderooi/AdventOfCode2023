package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

// numbers defines the number literals that can be found in the strings
var numbers = [...][]rune{
	{'o', 'n', 'e'},           // one
	{'t', 'w', 'o'},           // two
	{'t', 'h', 'r', 'e', 'e'}, // three
	{'f', 'o', 'u', 'r'},      // four
	{'f', 'i', 'v', 'e'},      // five
	{'s', 'i', 'x'},           // six
	{'s', 'e', 'v', 'e', 'n'}, // seven
	{'e', 'i', 'g', 'h', 't'}, // eight
	{'n', 'i', 'n', 'e'},      // nine
}

// direction defines the start and end values as well as the step direction for traversing the strings
type direction struct {
	start int
	end   int
	step  int
}

// newDirection is a constructional function for the direction object, it takes a rune slice and a boolean that is false
// for normal traversing and true for reverse
func newDirection(runes []rune, reverse bool) direction {
	strLen := len(runes)
	if reverse {
		return direction{
			start: strLen - 1,
			end:   -1,
			step:  -1,
		}
	} else {
		return direction{
			start: 0,
			end:   strLen,
			step:  1,
		}
	}
}

// traverseString converts the argument to a rune slice and iterates over it checking for digit's and number literals
func traverseString(str string, reverse bool, part2 bool) int {
	runes := []rune(str)
	dir := newDirection(runes, reverse)
	for i := dir.start; i != dir.end; i += dir.step {
		char := runes[i]
		if unicode.IsDigit(char) {
			return int(char - 48)
		}
		if part2 {
			found, n := findLiteral(runes, reverse, i)
			if found {
				return n
			}
		}
	}
	return 0
}

// findLiteral checks the rune from its provided index for 9 number literals as defined by numbers
func findLiteral(source []rune, reverse bool, index int) (bool, int) {
	var cut []rune
	if reverse {
		cut = source[:index+1]
	} else {
		cut = source[index:]
	}
	srcDir := newDirection(cut, reverse)
	for n, number := range numbers {
		dir := newDirection(number, reverse)
		j := srcDir.start
		for i := dir.start; i != dir.end && j != srcDir.end; i += dir.step {
			if cut[j] != number[i] {
				break
			}
			if i == dir.end-dir.step {
				return true, n + 1
			}
			j += srcDir.step
		}
	}
	return false, -1
}

// main the program entry point, call without flags for part one or -part2 for part2
func main() {
	part2 := flag.Bool("part2", false, "part two")
	flag.Parse()
	args := flag.Arg(0)
	var total int
	for _, arg := range strings.Split(args, "\n") {
		currentNumber := traverseString(arg, false, *part2)
		currentNumber *= 10
		currentNumber += traverseString(arg, true, *part2)
		total += currentNumber
	}
	fmt.Println(total)
}
