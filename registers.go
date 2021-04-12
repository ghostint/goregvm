package goregvm

// register interface.
// every register is supposed to implement these methods.
// size of register is 32bit.
type Register interface {
	Set(int32)
	Get() int32
}

// EAX register.
type EAX struct {
	value int32
}

func (e *EAX) Set(val int32) {
	e.value = val
}

func (e *EAX) Get() int32 {
	return e.value
}

// EBX register.
type EBX struct {
	value int32
}

func (e *EBX) Set(val int32) {
	e.value = val
}

func (e *EBX) Get() int32 {
	return e.value
}

// ECX register.
type ECX struct {
	value int32
}

func (e *ECX) Set(val int32) {
	e.value = val
}

func (e *ECX) Get() int32 {
	return e.value
}

// ECX register.
type EDX struct {
	value int32
}

func (e *EDX) Set(val int32) {
	e.value = val
}

func (e *EDX) Get() int32 {
	return e.value
}

// ECX register.
type ESP struct {
	value int32
}

func (e *ESP) Set(val int32) {
	e.value = val
}

func (e *ESP) Get() int32 {
	return e.value
}

// PC register. Program register.
type PC struct {
	value int32
}

func (e PC) Set(val int32) {
	e.value = val
}

func (e PC) Get() int32 {
	return e.value
}

// Flag registers.

// Zero Flag.
type ZF struct {
	is_set int32
}

func (x *ZF) Get() int32 {
	return x.is_set
}

func (x *ZF) Set(val int32) {
	x.is_set = val
}

// Sign Flag.
type SF struct {
	is_set int32
}

func (x *SF) Get() int32 {
	return x.is_set
}

func (x *SF) Set(val int32) {
	x.is_set = val
}
