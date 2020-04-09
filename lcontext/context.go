package lcontext

import (
	"context"
	"math"
	"time"
)

const abortIndex int8 = math.MaxInt8 / 2

type Business func(ctx context.Context, req interface{}) (resp interface{}, err error)
type ErrHandler func(ctx *Context, e error) (code int64, info string, err error)
type Context struct {
	parent     context.Context
	data       map[interface{}]interface{}
	index      int8
	handlers   []Middleware
	errHandler ErrHandler
	business   Business
	req        interface{}
	resp       interface{}
	err        error
	code       int64
	info       string
	clientIP   string
}

func (ctx *Context) Next() {
	ctx.index++
	if ctx.index <= int8(len(ctx.handlers)) {
		ctx.handlers[ctx.index-1](ctx)
		return
	}
	ctx.resp, ctx.err = ctx.business(ctx, ctx.req)
	if ctx.err == nil || ctx.errHandler == nil {
		return
	}
	ctx.code, ctx.info, ctx.err = ctx.errHandler(ctx, ctx.err)
}
func (ctx *Context) Abort() {
	ctx.index = abortIndex
}
func (ctx *Context) SetValue(key, value interface{}) {
	if ctx.data == nil {
		ctx.data = make(map[interface{}]interface{})
	}
	ctx.data[key] = value
}
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.parent.Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.parent.Done()
}

func (ctx *Context) Err() error {
	return ctx.parent.Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	if v, ok := ctx.data[key]; ok {
		return v
	}
	return ctx.parent.Value(key)
}
func Reset(ctx *Context, goCtx context.Context, req interface{}, ip string, _ map[string][]string) {
	ctx.parent = goCtx
	ctx.index = 0
	ctx.req = req
	ctx.resp = nil
	ctx.data = nil
	ctx.err = nil
	ctx.info = ""
	ctx.code = 0
	ctx.clientIP = ip
}
