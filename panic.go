package kkpanic

import (
	kklogger "github.com/kklab-com/goth-kklogger"
)

func Call(f func(v interface{})) {
	if v := recover(); v != nil {
		f(v)
	}
}

func Log() {
	if v := recover(); v != nil {
		switch t := v.(type) {
		case error:
			kklogger.ErrorJ("panic.Log", t.Error())
		}
	}
}

func Catch(main func(), panic func(v interface{})) {
	defer Call(panic)
	main()
}
