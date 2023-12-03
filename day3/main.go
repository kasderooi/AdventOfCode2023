package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

// symbol is an object that defines the location and its gear ratio
type symbol struct {
	x     int
	y     int
	gear  bool
	ratio int
}

// number defines the location, value and counted status of the part numbers that are found
type number struct {
	start   int
	end     int
	value   int
	counted bool
}

// chart contains all data for retrieving numbers and symbols
type chart struct {
	x       int
	xMax    int
	y       int
	yMax    int
	numbers [][]number
	symbols []symbol
	rows    []rune
}

// getPart checks if the part intersects with the given location and is not counted before, if so, returns its value
func (n *number) getPart(location int) int {
	isNext := location >= n.start-1 && location <= n.end+1
	if !n.counted && isNext {
		n.counted = true
		return n.value
	}
	return 0
}

// createNumber is a constructor function for a number object
func createNumber(c *chart) number {
	n := number{value: 0, counted: false}
	n.start = c.x
	for c.x < c.xMax && unicode.IsDigit(c.rows[c.x]) {
		n.value *= 10
		n.value += int(c.rows[c.x] - 48)
		c.x++
	}
	n.end = c.x - 1
	return n
}

// resolveGear resolves gears per row
func resolveGear(c *chart, sym *symbol, parts *int, offset int) {
	for i := 0; i < len(c.numbers[sym.y+offset]); i++ {
		part := c.numbers[sym.y+offset][i].getPart(sym.x)
		if part > 0 {
			*parts++
			sym.ratio *= part
		}
	}
}

// resolveGears checks all parts for crossing with gears AKA symbol that have the gear flag
func (c *chart) resolveGears() int {
	total := 0
	for _, sym := range c.symbols {
		parts := 0
		if sym.gear {
			if sym.y > 0 {
				resolveGear(c, &sym, &parts, -1)
			}
			resolveGear(c, &sym, &parts, 0)
			if sym.y < c.yMax {
				resolveGear(c, &sym, &parts, 1)
			}
		}
		if parts > 1 {
			total += sym.ratio
		}
	}
	return total
}

// resolveSymbols checks all parts for crossing with symbol
func (c *chart) resolveSymbols() int {
	total := 0
	for _, sym := range c.symbols {
		if sym.y > 0 {
			for i := 0; i < len(c.numbers[sym.y-1]); i++ {
				total += c.numbers[sym.y-1][i].getPart(sym.x)
			}
		}
		for i := 0; i < len(c.numbers[sym.y]); i++ {
			total += c.numbers[sym.y][i].getPart(sym.x)
		}
		if sym.y < c.yMax {
			for i := 0; i < len(c.numbers[sym.y+1]); i++ {
				total += c.numbers[sym.y+1][i].getPart(sym.x)
			}
		}
	}
	return total
}

// parseRow parses the string for its symbols and numbers
func (c *chart) parseRow() {
	c.x = 0
	for c.x < c.xMax {
		if unicode.IsDigit(c.rows[c.x]) {
			c.numbers[c.y] = append(c.numbers[c.y], createNumber(c))
		} else {
			if c.rows[c.x] != '.' {
				c.symbols = append(c.symbols, symbol{x: c.x, y: c.y, gear: c.rows[c.x] == '*', ratio: 1})
			}
			c.x++
		}
	}
}

// createChart is a constructor function for creating the chart, parsing the data and initializing the symbol and number
// slices
func createChart(str string) *chart {
	stringSlice := strings.Split(str, "\n")
	ret := chart{
		x:       0,
		xMax:    len(stringSlice[0]),
		y:       0,
		yMax:    len(stringSlice),
		numbers: make([][]number, 150),
		symbols: make([]symbol, 150),
	}
	for i, slice := range stringSlice {
		ret.y = i
		ret.rows = []rune(slice)
		ret.xMax = len(ret.rows)
		ret.parseRow()
	}
	return &ret
}

// main the program entry point, call without flags for part one or -part2 for part2
func main() {
	part2 := flag.Bool("part2", false, "part two")
	flag.Parse()
	args := flag.Arg(0)
	c := createChart(args)
	if *part2 {
		fmt.Println(c.resolveGears())
	} else {
		fmt.Println(c.resolveSymbols())
	}
}
