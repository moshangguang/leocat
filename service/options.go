package service

import (
	"github.com/moshangguang/leocat/lcontext"
	"github.com/moshangguang/leocat/method"
)

type Options struct {
	Name       string
	Middleware []lcontext.Middleware
	ErrHandler lcontext.ErrHandler
	ExMethod   []method.Method
}

type Option func(options *Options)

func Name(name string) Option {
	return func(options *Options) {
		if name != "" {
			options.Name = name
		}
	}
}

func AddMethods(methods ...method.Method) Option {
	return func(options *Options) {
		options.ExMethod = append(options.ExMethod, methods...)
	}
}

func WrapHandler(handler lcontext.Middleware) Option {
	return func(options *Options) {
		if options.Middleware != nil {
			options.Middleware = append(options.Middleware, handler)
		}
	}
}
func ErrorHandler(errHandler lcontext.ErrHandler) Option {
	return func(options *Options) {
		options.ErrHandler = errHandler
	}
}
