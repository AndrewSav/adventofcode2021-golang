package day24

import (
	"aoc2021/util"
	"fmt"
	"regexp"
	"strconv"
)

type instruction struct {
	dst      int
	src      int
	isSrcVar bool
	op       func(vm *vm, i instruction)
}

type vm struct {
	input        []int
	inputPointer int
	memory       [4]int
	program      []instruction
}

func inp(vm *vm, i instruction) {
	vm.memory[i.dst] = vm.input[vm.inputPointer]
	vm.inputPointer++
}

func getSrc(vm *vm, i instruction) (src int) {
	if i.isSrcVar {
		src = vm.memory[i.src]
	} else {
		src = i.src
	}
	return
}

func add(vm *vm, i instruction) {
	vm.memory[i.dst] += getSrc(vm, i)
}

func mul(vm *vm, i instruction) {
	vm.memory[i.dst] *= getSrc(vm, i)
}

func div(vm *vm, i instruction) {
	vm.memory[i.dst] /= getSrc(vm, i)
}

func mod(vm *vm, i instruction) {
	vm.memory[i.dst] %= getSrc(vm, i)
}

func eql(vm *vm, i instruction) {
	if vm.memory[i.dst] == getSrc(vm, i) {
		vm.memory[i.dst] = 1
	} else {
		vm.memory[i.dst] = 0
	}
}

var ops = map[string]func(vm *vm, i instruction){
	"inp": inp,
	"add": add,
	"mul": mul,
	"div": div,
	"mod": mod,
	"eql": eql,
}

func run(vm *vm) {
	for _, i := range vm.program {
		i.op(vm, i)
	}
}

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	r := regexp.MustCompile(`(\w\w\w) (w|x|y|z) ?(w|x|y|z|-?\d+)?`)
	vm := vm{}

	for _, l := range lines {
		match := r.FindStringSubmatch(l)
		var isSrcVar bool
		var val int
		if len(match[3]) > 0 {
			v, err := strconv.Atoi(match[3])
			if err == nil {
				val = v
			} else {
				isSrcVar = true
				val = int([]rune(match[3])[0] - 'w')
			}
		}
		i := instruction{
			src:      val,
			isSrcVar: isSrcVar,
			dst:      int([]rune(match[2])[0] - 'w'),
			op:       ops[match[1]],
		}
		vm.program = append(vm.program, i)
	}
	input := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	vm.input = input[:]
	for {
		run(&vm)
		if vm.memory[3] == 0 {
			return fmt.Sprint(input)
		}
		vm.inputPointer = 0
		vm.memory = [4]int{}
		for i := 13; i >= 0; i-- {
			if input[i] > 1 {
				input[i]--
				break
			}
			input[i] = 9
		}
	}
}
