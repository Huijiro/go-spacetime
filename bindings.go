package main

import (
	"unsafe"
)

type LogLevel uint32

const (
	LOG_LEVEL_ERROR LogLevel = 0
	LOG_LEVEL_WARN  LogLevel = 1
	LOG_LEVEL_INFO  LogLevel = 2
	LOG_LEVEL_DEBUG LogLevel = 3
	LOG_LEVEL_TRACE LogLevel = 4
	LOG_LEVEL_PANIC LogLevel = 101
)

//go:wasmimport spacetime_10.0 console_log
func console_log(
	level LogLevel,

	targetPtr *byte,
	targetLen uintptr,

	filenamePtr *byte,
	filenameLen uintptr,

	lineNumber uint32,

	messagePtr *byte,
	messageLen uintptr,
)

//go:wasmimport spacetime_10.0 BytesSource
type BytesSource = uint32
type BytesSink = uint32

//export __describe_module__
func __describe_module__(description BytesSink) {
	t := "my_wasm_module"
	f := "main.go"
	m := "describing module without import"

	console_log(
		LOG_LEVEL_INFO,
		unsafe.StringData(t), unsafe.Sizeof(t),
		unsafe.StringData(f), unsafe.Sizeof(f),
		123,
		unsafe.StringData(m), unsafe.Sizeof(m),
	)
}

//export __call_reducer__
func __call_reducer__(
	id uintptr,
	sender_0 uint64,
	sender_1 uint64,
	sender_2 uint64,
	sender_3 uint64,
	conn_id_0 uint64,
	conn_id_1 uint64,
	timestamp uint64,
	args BytesSource,
	err BytesSink,
) uint32 {
	return 1
}

func main() {}
