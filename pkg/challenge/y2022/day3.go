package y2022

import (
	"AdventOfCode/pkg/data_structure/slice"
	"AdventOfCode/pkg/io/file"
	"go.uber.org/zap"
)

type Day3 struct {
	Year      int
	Day       int
	InputFile string
	Logger    *zap.SugaredLogger
}

type elfGroup struct {
	rucksacks []*rucksack
}

type rucksack struct {
	items        string
	compartments []*compartment
}

type compartment struct {
	items []rune
}

func newRucksack(items string) *rucksack {
	numItems := len(items)
	compartmentSize := numItems / 2

	var compartments []*compartment
	firstCompartment := newCompartment([]rune(items[0:compartmentSize]))
	secondCompartment := newCompartment([]rune(items[compartmentSize:numItems]))
	compartments = append(compartments, firstCompartment)
	compartments = append(compartments, secondCompartment)

	ruck := &rucksack{
		items:        items,
		compartments: compartments,
	}

	return ruck
}

func newCompartment(items []rune) *compartment {
	return &compartment{
		items: items,
	}
}

func (e *elfGroup) TotalPriority() int {

	priority := 0
	intersect := []rune(e.rucksacks[0].items)
	for i, _ := range e.rucksacks {
		if i > 0 {
			intersect = slice.Intersection(intersect, []rune(e.rucksacks[i].items))
		}
	}

	countedItems := make(map[rune]bool)
	for _, item := range intersect {
		if !countedItems[item] {
			priority += getPriority(item)
			countedItems[item] = true
		}
	}

	return priority
}

func (r *rucksack) TotalPriority() int {
	compartmentOne := r.compartments[0]
	compartmentTwo := r.compartments[1]

	priority := 0
	intersect := slice.Intersection(compartmentOne.items, compartmentTwo.items)

	countedItems := make(map[rune]bool)
	for _, item := range intersect {
		if !countedItems[item] {
			priority += getPriority(item)
			countedItems[item] = true
		}
	}

	return priority
}

func getPriority(item rune) int {
	ascii := int(item)
	if ascii >= int('A') && ascii <= int('Z') {
		return ascii - int('&')
	}

	return ascii - int('`')
}

func (d *Day3) Solve() error {
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

func (d *Day3) Part1() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	totalPriority := 0
	for _, line := range lines {
		rucksack := newRucksack(line)
		priority := rucksack.TotalPriority()
		totalPriority += priority

		//d.Logger.Infof("Rucksack %v, Priority: %v, Cumulative: %v", i, priority, totalPriority)
	}

	return totalPriority, nil
}

func (d *Day3) Part2() (interface{}, error) {
	fr := file.Reader{Path: d.InputFile}
	lines, err := fr.GetContents()
	if err != nil {
		return nil, err
	}

	var rucksacks []*rucksack
	groupCount := 0
	totalPriority := 0
	for _, line := range lines {
		rucksacks = append(rucksacks, newRucksack(line))

		if len(rucksacks) == 3 {
			groupCount++
			eg := &elfGroup{rucksacks: rucksacks}
			priority := eg.TotalPriority()
			totalPriority += priority
			//d.Logger.Infof("Group %v, Priority: %v, Cumulative: %v", groupCount, priority, totalPriority)

			rucksacks = rucksacks[:0]
		}
	}

	return totalPriority, nil
}
