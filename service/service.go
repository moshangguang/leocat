package service

import (
	"github.com/moshangguang/leocat/lcontext"
	"github.com/moshangguang/leocat/method"
)

type Service interface {
	Name() string
	Methods() []method.Method
	Middleware() []lcontext.Middleware
	ErrHandler() lcontext.ErrHandler
}

type service struct {
	name       string
	methods    []method.Method
	middleware []lcontext.Middleware
	errHandler lcontext.ErrHandler
}

func (svc service) Name() string {
	return svc.name
}

func (svc service) Methods() []method.Method {
	return svc.methods
}

func (svc service) Middleware() []lcontext.Middleware {
	return svc.middleware
}

func (svc service) ErrHandler() lcontext.ErrHandler {
	return svc.errHandler
}

func New(options Options) Service {
	ser := service{
		name:       options.Name,
		methods:    options.ExMethod,
		middleware: options.Middleware,
		errHandler: options.ErrHandler,
	}
	return ser
}
