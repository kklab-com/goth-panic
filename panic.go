package kkpanic

import (
	"bytes"
	"encoding/json"
	"runtime/debug"
	"runtime/pprof"

	kklogger "github.com/kklab-com/goth-kklogger"
)

func Call(f func(v interface{})) {
	if v := recover(); v != nil {
		f(v)
	}
}

func Log() {
	if v := recover(); v != nil {
		buffer := &bytes.Buffer{}
		pprof.Lookup("goroutine").WriteTo(buffer, 1)
		switch t := v.(type) {
		case error:
			kklogger.ErrorJ("panic.Log", CaughtError{
				Err:             t.Error(),
				PanicCallStack:  string(debug.Stack()),
				GoRoutineStacks: buffer.String(),
			})
		}
	}
}

func Catch(main func(), panic func(v interface{})) {
	defer Call(panic)
	main()
}

type CaughtError struct {
	Err             string `json:"error,omitempty"`
	PanicCallStack  string `json:"panic_call_stack,omitempty"`
	GoRoutineStacks string `json:"go_routine_stacks,omitempty"`
}

func (e *CaughtError) Error() string {
	bs, _ := json.Marshal(e)
	return string(bs)
}
