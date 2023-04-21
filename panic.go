package kkpanic

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"runtime/pprof"
	"strings"
)

func CallStack() string {
	return string(debug.Stack())
}

func GoroutineStacks() string {
	buffer := &bytes.Buffer{}
	pprof.Lookup("goroutine").WriteTo(buffer, 1)
	return buffer.String()
}

func Convert(v interface{}) *CaughtImpl {
	if v == nil {
		return nil
	}

	buffer := &bytes.Buffer{}
	pprof.Lookup("goroutine").WriteTo(buffer, 1)
	c := &CaughtImpl{
		obj:             v,
		CallStack:       string(debug.Stack()),
		GoroutineStacks: buffer.String(),
	}

	switch cast := v.(type) {
	case []byte:
		c.Message = strings.ToUpper(hex.EncodeToString(cast))
	case string:
		c.Message = cast
	case error:
		c.Message = cast.Error()
	case fmt.Stringer:
		c.Message = cast.String()
	default:
		bs, _ := json.Marshal(v)
		c.Message = string(bs)
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

func Catch(main func(), panic func(r Caught)) {
	defer Call(panic)
	main()
}

func CatchExcept(main func(), except interface{}, panic func(r Caught)) {
	defer CallExcept(except, panic)
	main()
}

func PanicNonNil(obj interface{}) {
	if obj != nil {
		panic(obj)
	}
}

type Caught interface {
	String() string
	Error() string
	Data() interface{}
	Trace() CaughtTrace
}

type CaughtTrace interface {
	CallStackString() string
	GoroutineStacksString() string
}

type CaughtImpl struct {
	obj             interface{}
	Message         string `json:"message,omitempty"`
	CallStack       string `json:"call_stack,omitempty"`
	GoroutineStacks string `json:"goroutine_stacks,omitempty"`
}

func (e *CaughtImpl) String() string {
	bs, _ := json.Marshal(e)
	return string(bs)
}

func (e *CaughtImpl) Error() string {
	return e.Message
}

func (e *CaughtImpl) Data() interface{} {
	return e.obj
}

func (e *CaughtImpl) Trace() CaughtTrace {
	return e
}

func (e *CaughtImpl) CallStackString() string {
	return e.CallStack
}

func (e *CaughtImpl) GoroutineStacksString() string {
	return e.GoroutineStacks
}
