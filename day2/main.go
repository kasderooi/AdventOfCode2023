package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// colourPair defines a type for key-value pairs
type colourPair struct {
	colour string
	number int
}

// minColours keeps track of the minimal value needed to make a game possible in part 2
type minColours struct {
	red   int
	green int
	blue  int
}

// isPossible returns true if the number of this colourPair is smaller than the maximum allowed
func (c colourPair) isPossible() bool {
	return c.number <= maxColours[c.colour]
}

// getPower returns the product of all three colours of this minColours object
func (m *minColours) getPower() int {
	return m.red * m.green * m.blue
}

// setColour takes a colourPair and sets the corresponding colour number to whatever is higher: existing or provided
func (m *minColours) setColour(pair colourPair) {
	switch pair.colour {
	case "red":
		if m.red < pair.number {
			m.red = pair.number
		}
	case "green":
		if m.green < pair.number {
			m.green = pair.number
		}
	case "blue":
		if m.blue < pair.number {
			m.blue = pair.number
		}
	}
}

// initMinColour is a constructor function for colourPair
func initColourPair(str string) colourPair {
	pair := strings.Split(str, " ")
	n, _ := strconv.Atoi(pair[1])
	return colourPair{
		colour: pair[2],
		number: n,
	}
}

// initMinColour is a constructor function for minColours
func initMinColour() minColours {
	return minColours{
		red:   1,
		green: 1,
		blue:  1,
	}
}

// splitToColour splits the string into substrings with one colour-number pair
func splitToColour(str string) []string {
	return strings.Split(str, ",")
}

// splitToSets splits the string into substrings that each contain one game subset
func splitToSets(str string) []string {
	return strings.Split(str, ";")
}

// parseGamePart1 parses the line with the game and returns true if the game is allowed following rules of part 1
func parseGamePart1(str string) bool {
	for _, set := range splitToSets(str) {
		for _, colour := range splitToColour(set) {
			if !initColourPair(colour).isPossible() {
				return false
			}
		}
	}
	return true
}

// parseGamePart2 parses the line with the game and returns the power of the game following the rules of part 2
func parseGamePart2(str string) int {
	colours := initMinColour()
	for _, set := range splitToSets(str) {
		for _, colour := range splitToColour(set) {
			colours.setColour(initColourPair(colour))
		}
	}
	return colours.getPower()
}

// maxColours is the max amount of stones available to make a game possible in part 1
var maxColours = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// main the program entry point, call without flags for part one or -part2 for part2
func main() {
	part2 := flag.Bool("part2", false, "part two")
	flag.Parse()
	args := flag.Arg(0)
	var total int
	for n, arg := range strings.Split(args, "\n") {
		slice := strings.Split(arg, ":")
		if len(slice) > 1 {
			if *part2 {
				total += parseGamePart2(slice[1])
			} else if parseGamePart1(slice[1]) {
				total += n + 1
			}
		}
	}
	fmt.Println(total)
}
