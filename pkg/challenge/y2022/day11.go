package y2022

import (
	"sort"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/io/file"
)

type Day11 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

type (
	operation func(int64, int64) int64
	test      func(int64, int64) bool
)

var monkeys []*monkey

type monkey struct {
	Items             []*item
	Operation         operation
	Operand           int64
	Test              test
	TestVal           int64
	PassVal           int64
	FailVal           int64
	InspectionCount   int64
	SimplifyOperation operation
	SimplifyOperand   int64
}

func (m *monkey) TestItems() {
	for _, item := range m.Items {
		m.InspectionCount++
		item.WorryLevel = m.Operation(item.WorryLevel, m.Operand)
		item.WorryLevel = m.SimplifyOperation(item.WorryLevel, m.SimplifyOperand)
		passesTest := m.Test(item.WorryLevel, m.TestVal)
		if passesTest {
			monkeys[m.PassVal].Items = append(monkeys[m.PassVal].Items, item)
		} else {
			monkeys[m.FailVal].Items = append(monkeys[m.FailVal].Items, item)
		}
	}
	m.Items = make([]*item, 0)
}

type item struct {
	WorryLevel int64
}

func plus(old int64, addend int64) int64 {
	if addend == -1 {
		addend = old
	}
	return old + addend
}

func multiply(old int64, multiplier int64) int64 {
	if multiplier == -1 {
		multiplier = old
	}
	return old * multiplier
}

func divide(old int64, divisor int64) int64 {
	return old / divisor
}

func mod(old int64, modulo int64) int64 {
	return old % modulo
}

func divisible(value int64, divisor int64) bool {
	return value%divisor == 0
}

func (d *Day11) Solve() error {
	a1, err := d.Part1()
	if err != nil {
		return err
	}
	a2, err := d.Part2()
	if err != nil {
		return err
	}

	d.Logger.Infof("Day %v, %v completed. Part 1: %v, Part 2: %v", d.Day, d.Year, a1, a2)

	return nil
}

func (d *Day11) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	monkeys = make([]*monkey, 0)
	var currMonkey *monkey
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey ") {
			if currMonkey != nil {
				monkeys = append(monkeys, currMonkey)
			}
			currMonkey = &monkey{
				Items:             make([]*item, 0),
				SimplifyOperation: divide,
				SimplifyOperand:   3,
			}
		} else if strings.HasPrefix(line, "  Starting items: ") {
			itemLevels := strings.Split(strings.ReplaceAll(line, "  Starting items: ", ""), ", ")
			for _, itemLevel := range itemLevels {
				level, err := strconv.Atoi(itemLevel)
				if err != nil {
					return nil, err
				}
				currMonkey.Items = append(currMonkey.Items, &item{WorryLevel: int64(level)})
			}
		} else if strings.HasPrefix(line, "  Operation: new = ") {
			baseOp := strings.ReplaceAll(line, "  Operation: new = ", "")
			pieces := strings.Split(baseOp, " ")
			var operand int
			if pieces[2] == "old" {
				operand = -1
			} else {
				operand, err = strconv.Atoi(pieces[2])
				if err != nil {
					return nil, err
				}
			}
			currMonkey.Operand = int64(operand)
			switch pieces[1] {
			case "*":
				currMonkey.Operation = multiply
			case "+":
				currMonkey.Operation = plus
			}
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			baseTest := strings.ReplaceAll(line, "  Test: divisible by ", "")
			divisor, err := strconv.Atoi(baseTest)
			if err != nil {
				return nil, err
			}
			currMonkey.TestVal = int64(divisor)
			currMonkey.Test = divisible
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			baseTest := strings.ReplaceAll(line, "    If true: throw to monkey ", "")
			trueVal, err := strconv.Atoi(baseTest)
			if err != nil {
				return nil, err
			}
			currMonkey.PassVal = int64(trueVal)
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			baseTest := strings.ReplaceAll(line, "    If false: throw to monkey ", "")
			falseVal, err := strconv.Atoi(baseTest)
			if err != nil {
				return nil, err
			}
			currMonkey.FailVal = int64(falseVal)
		}
	}

	// Get the last monkey in the slice
	if currMonkey != nil {
		monkeys = append(monkeys, currMonkey)
	}

	roundCount := 20
	for i := roundCount; i > 0; i-- {
		for _, monkey := range monkeys {
			monkey.TestItems()
		}
	}

	inspectionCounts := make([]int, 0)
	for _, monkey := range monkeys {
		inspectionCounts = append(inspectionCounts, int(monkey.InspectionCount))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectionCounts)))
	result := inspectionCounts[0] * inspectionCounts[1]

	return result, nil
}

func (d *Day11) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	monkeys = make([]*monkey, 0)
	modulo := int64(1)
	var currMonkey *monkey
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey ") {
			if currMonkey != nil {
				monkeys = append(monkeys, currMonkey)
			}
			currMonkey = &monkey{
				Items:             make([]*item, 0),
				SimplifyOperation: mod,
			}
		} else if strings.HasPrefix(line, "  Starting items: ") {
			itemLevels := strings.Split(strings.ReplaceAll(line, "  Starting items: ", ""), ", ")
			for _, itemLevel := range itemLevels {
				level, err := strconv.Atoi(itemLevel)
				if err != nil {
					return nil, err
				}
				currMonkey.Items = append(currMonkey.Items, &item{WorryLevel: int64(level)})
			}
		} else if strings.HasPrefix(line, "  Operation: new = ") {
			baseOp := strings.ReplaceAll(line, "  Operation: new = ", "")
			pieces := strings.Split(baseOp, " ")
			var operand int
			if pieces[2] == "old" {
				operand = -1
			} else {
				operand, err = strconv.Atoi(pieces[2])
				if err != nil {
					return nil, err
				}
			}
			currMonkey.Operand = int64(operand)
			switch pieces[1] {
			case "*":
				currMonkey.Operation = multiply
			case "+":
				currMonkey.Operation = plus
			}
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			baseTest := strings.ReplaceAll(line, "  Test: divisible by ", "")
			divisor, err := strconv.Atoi(baseTest)
			if err != nil {
				return nil, err
			}
			modulo *= int64(divisor)
			currMonkey.TestVal = int64(divisor)
			currMonkey.Test = divisible
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			baseTest := strings.ReplaceAll(line, "    If true: throw to monkey ", "")
			trueVal, err := strconv.Atoi(baseTest)
			if err != nil {
				return nil, err
			}
			currMonkey.PassVal = int64(trueVal)
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			baseTest := strings.ReplaceAll(line, "    If false: throw to monkey ", "")
			falseVal, err := strconv.Atoi(baseTest)
			if err != nil {
				return nil, err
			}
			currMonkey.FailVal = int64(falseVal)
		}
	}

	// Get the last monkey in the slice
	if currMonkey != nil {
		monkeys = append(monkeys, currMonkey)
	}

	for _, monkey := range monkeys {
		monkey.SimplifyOperand = modulo
	}

	roundCount := 10000
	for i := roundCount; i > 0; i-- {
		for _, monkey := range monkeys {
			monkey.TestItems()
		}
	}

	inspectionCounts := make([]int, 0)
	for _, monkey := range monkeys {
		inspectionCounts = append(inspectionCounts, int(monkey.InspectionCount))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectionCounts)))
	result := inspectionCounts[0] * inspectionCounts[1]

	return result, nil
}
