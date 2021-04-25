package kkpanic

import (
	"bytes"
	"encoding/json"
	"runtime/debug"
	"runtime/pprof"

	kklogger "github.com/kklab-com/goth-kklogger"
)

func Convert(v interface{}) *CaughtImpl {
	if v == nil {
		return nil
	}

	buffer := &bytes.Buffer{}
	pprof.Lookup("goroutine").WriteTo(buffer, 1)
	c := &CaughtImpl{
		Message:         v,
		CallStack:       string(debug.Stack()),
		GoroutineStacks: buffer.String(),
	}

	buffer.Reset()
	return c
}

func Call(f func(r Caught)) {
	if v := recover(); v != nil {
		f(Convert(v))
	}
}

func CallExcept(except interface{}, f func(r Caught)) {
	if v := recover(); v != nil && except != v {
		f(Convert(v))
	}
}

func Log() {
	Call(func(r Caught) {
		kklogger.ErrorJ("panic.Log", r)
	})
}

func LogExcept(except interface{}) {
	CallExcept(except, func(r Caught) {
		kklogger.ErrorJ("panic.Log", r)
	})
}

func Catch(main func(), panic func(r Caught)) {
	defer Call(panic)
	main()
}

func CatchExcept(main func(), except interface{}, panic func(r Caught)) {
	defer CallExcept(except, panic)
	main()
}

type Caught interface {
	String() string
	Error() string
	Data() interface{}
}

type CaughtImpl struct {
	Message         interface{} `json:"message,omitempty"`
	CallStack       string      `json:"call_stack,omitempty"`
	GoroutineStacks string      `json:"goroutine_stacks,omitempty"`
}

func (e *CaughtImpl) String() string {
	bs, _ := json.Marshal(e)
	return string(bs)
}

func (e *CaughtImpl) Error() string {
	bs, _ := json.Marshal(e.Message)
	return string(bs)
}

func (e *CaughtImpl) Data() interface{} {
	return e.Message
}
