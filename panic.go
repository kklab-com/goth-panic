package kkpanic

import (
	"bytes"
	"encoding/json"
	"runtime/debug"
	"runtime/pprof"

	kklogger "github.com/kklab-com/goth-kklogger"
)

func Convert(v interface{}) *Caught {
	if v == nil {
		return nil
	}

	buffer := &bytes.Buffer{}
	pprof.Lookup("goroutine").WriteTo(buffer, 1)
	c := &Caught{
		Message:         v,
		PanicCallStack:  string(debug.Stack()),
		GoRoutineStacks: buffer.String(),
	}

	buffer.Reset()
	return c
}

func Call(f func(r *Caught)) {
	if v := recover(); v != nil {
		f(Convert(v))
	}
}

func Log() {
	Call(func(r *Caught) {
		kklogger.ErrorJ("panic.Log", r)
	})
}

func Catch(main func(), panic func(r *Caught)) {
	defer Call(panic)
	main()
}

type Caught struct {
	Message         interface{} `json:"message,omitempty"`
	PanicCallStack  string      `json:"panic_call_stack,omitempty"`
	GoRoutineStacks string      `json:"go_routine_stacks,omitempty"`
}

func (e *Caught) String() string {
	bs, _ := json.Marshal(e)
	return string(bs)
}

func (e *Caught) Error() string {
	bs, _ := json.Marshal(e.Message)
	return string(bs)
}
