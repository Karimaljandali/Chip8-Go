package main

// import "fmt"

type Chip8 struct {
	Memory         [4096]uint8
	Registers      [16]uint8
	Sound          uint8
	Delay          uint8
	StackPointer   uint16
	ProgramCounter uint16
	Stack          [16]uint16
	Display        [64][32]uint16
	Keyboard       [16]uint8
	OpCode         uint16
}
