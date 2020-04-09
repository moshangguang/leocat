package lcontext

import (
	"fmt"
	"runtime"
)

type Middleware func(ctx *Context)

func DefaultMiddleware() []Middleware {
	ms := make([]Middleware, 0, 1)
	ms = append(ms, func(ctx *Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			ctx.err = fmt.Errorf("panic Error: %v;stack: %s", err, buf[:n])
		}()
		ctx.Next()
	})
	return ms
}
