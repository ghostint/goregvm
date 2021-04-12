package goregvm

import (
	"fmt"
	"log"
	"os"
)

type VM struct {
	stack *Stack
	pc    PC // special register.

	memory []int32

	registers map[int32]Register // all other registers.
}

func (v VM) Run() {

	v.stack = &Stack{}
	v.stack.sp = -1

	for {

		v.pc.value++ // during initial pc = -1
		if int(v.pc.value) > len(v.memory) {
			log.Fatalln("Program executed successfully")
		}

		ins := v.memory[v.pc.value] // first instruction.

		switch ins {

		case Mov:
			v.pc.value++
			reg := v.memory[v.pc.value]
			v.pc.value++

			switch reg {

			case Eax:
				v.registers[Eax].Set(v.memory[v.pc.value])

			case Ebx:
				v.registers[Ebx].Set(v.memory[v.pc.value])

			case Ecx:
				v.registers[Ecx].Set(v.memory[v.pc.value])

			case Edx:
				v.registers[Edx].Set(v.memory[v.pc.value])

			default:
				log.Fatalf("Illegal MOV instruction at %v\n", v.pc.value)

			}

		case Push:
			v.pc.value++

			var reg = v.memory[v.pc.value]

			switch reg {
			case Eax:
				v.stack.Push(v.registers[Eax].Get())
			case Ebx:
				v.stack.Push(v.registers[Ebx].Get())
			case Ecx:
				v.stack.Push(v.registers[Ecx].Get())

			case Esp:
				v.stack.Push(v.registers[Esp].Get())

			case Ebp:
				v.stack.Push(v.registers[Ebp].Get())
			case Edx:
				v.stack.Push(v.registers[Edx].Get())

			default:
				v.stack.Push(reg)

			}

		case Pop:
			v.pc.value++

			var nxt = v.memory[v.pc.value]

			switch nxt {
			case Eax:
				v.stack.Pop(v.registers[Eax])

			case Ebx:
				v.stack.Pop(v.registers[Ebx])

			case Ecx:
				v.stack.Pop(v.registers[Ecx])

			case Esp:
				v.stack.Pop(v.registers[Esp])

			case Ebp:
				v.stack.Pop(v.registers[Ebp])

			case Edx:
				v.stack.Pop(v.registers[Edx])

			default:
				log.Fatalf("Not a valid pop instruction at %v \n", v.pc.value)

			}

		case Add:

			v.pc.value++

			dst := v.registers[v.memory[v.pc.value]]
			v.pc.value++
			src := v.registers[v.memory[v.pc.value]]
			result := dst.Get() + src.Get()

			dst.Set(result)

		case Sub:
			v.pc.value++

			dst := v.registers[v.memory[v.pc.value]]
			v.pc.value++
			src := v.registers[v.memory[v.pc.value]]
			result := dst.Get() - src.Get()

			dst.Set(result)

		case Mul:
			v.pc.value++

			dst := v.registers[v.memory[v.pc.value]]
			v.pc.value++
			src := v.registers[v.memory[v.pc.value]]
			result := dst.Get() * src.Get()

			dst.Set(result)

		case Div:
			continue

		case Print:
			// TODO: Use switch statement for viewing different registers.
			v.pc.value++ // next keyword

			reg := v.memory[v.pc.value]

			switch reg {
			case Eax:
				fmt.Printf("EAX: %v\n", v.registers[Eax].Get())
			case Ebx:
				fmt.Printf("EBX: %v\n", v.registers[Ebx].Get())
			case Ecx:
				fmt.Printf("ECX: %v\n", v.registers[Ecx].Get())
			case Edx:
				fmt.Printf("EDX: %v\n", v.registers[Edx].Get())

			case Pc:
				fmt.Printf("PC: %v\n", v.pc.value)

			case Zf:
				fmt.Printf("ZF: %v\n", v.registers[Zf].Get())
			case Sf:
				fmt.Printf("SF: %v\n", v.registers[Sf].Get())

			}

		case Cmp:

			v.pc.value++

			n := v.registers[v.memory[v.pc.value]]
			v.pc.value++
			nn := v.registers[v.memory[v.pc.value]]

			result := n.Get() - nn.Get()

			if result == 0 {
				v.registers[Zf].Set(1)
			}
			if result < 0 {
				v.registers[Sf].Set(1)
			}

		case Jz:
			v.pc.value++
			if v.registers[Zf].Get() == 1 {
				new_pc := v.memory[v.pc.value]
				v.pc.value = new_pc
			}

		case Jnz:
			v.pc.value++
			if v.registers[Zf].Get() == 0 {
				new_pc := v.memory[v.pc.value]
				v.pc.value = new_pc
			}

		case Inc:
			v.pc.value++

			reg := v.registers[v.memory[v.pc.value]]
			reg.Set(reg.Get() + 1)

		case Dec:
			v.pc.value++
			reg := v.registers[v.memory[v.pc.value]]
			reg.Set(reg.Get() - 1)

		case Exit:

			// fmt.Println("exit opcode encountered")
			os.Exit(0)

		default:
			log.Fatalf("Illegal instruction at %v\n", v.pc.value)

		}

	}

}

func NewVM() VM {

	// initializing registers.
	var registers = map[int32]Register{
		Eax: &EAX{value: 0},
		Ebx: &EBX{value: 0},
		Ecx: &ECX{value: 0},
		Edx: &EDX{value: 0},
		Esp: &ESP{value: 0},
		Sf:  &SF{is_set: 0},
		Zf:  &ZF{is_set: 0},
	}

	// initializing program counter.
	pc := PC{}
	pc.value = -1

	// code to be executed.
	// code := []int32{

	// 	Mov, Eax, 10, // eax = 10
	// 	Mov, Ebx, 20, // ebx = 20
	// 	Add, Eax, Ebx, // eax = eax + ebx (30)
	// 	Push, Eax, // storing eax in stack i.e 30
	// 	Pop, Ecx, // ecx = 30; poping the value from stack into ecx.
	// 	Add, Eax, Ecx, // eax = eax + ecx = 30+30 = 60
	// 	Print, Eax,
	// 	Mul, Eax, Ebx,
	// 	Print, Eax,
	// 	Exit,
	// }

	code := []int32{
		Mov, Eax, 50,
		Mov, Ecx, 0,
		Print, Eax,
		Dec, Eax,
		Cmp, Ecx, Eax,
		Jnz, 5,

		Exit,
	}

	return VM{pc: pc, memory: code, registers: registers}

}
