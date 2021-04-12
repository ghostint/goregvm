// Different type of opcodes present in this VM.
package goregvm

const (

	// Stack operations.

	Push int32 = iota
	Pop

	// Math operations.

	Add
	Sub
	Mul
	Div

	// Note: not implemented yet.
	//Bitwise operations.

	Xor
	And
	Or
	Not

	// flow control opcodes.

	Jmp
	Jz
	Jnz
	Cmp
	Exit

	// misc operations.

	Mov
	Print
	Inc
	Dec

	// register opcodes.

	Eax
	Ebx
	Ecx
	Edx
	Esp
	Ebp
	Pc
	Sf
	Zf
)
