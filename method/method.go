package method

import (
	"context"
	"github.com/moshangguang/leocat/lcontext"
	"reflect"
	"sync"
	"sync/atomic"
)

type Method interface {
	Name() string
	Service() string
	InType() reflect.Type
	OutType() reflect.Type
	SetErrHandle(errHandler lcontext.ErrHandler)
	SetHandlers(handlers []lcontext.Middleware)
	Lock()
	Exec(ctx context.Context, req interface{}, ip string) (resp interface{}, code int64, info string, err error)
}

type method struct {
	lock        bool
	name        string
	service     string
	business    lcontext.Business
	handlers    []lcontext.Middleware
	pool        *sync.Pool
	errHandler  lcontext.ErrHandler
	inType      reflect.Type
	outType     reflect.Type
	invokeTimes uint64
	errTimes    uint64
}

func (m *method) Lock() {
	m.lock = true
}

func (m *method) InType() reflect.Type {
	return m.inType
}

func (m *method) OutType() reflect.Type {
	return m.outType
}

func (m *method) SetErrHandle(errHandler lcontext.ErrHandler) {
	if m.lock {
		return
	}
	if errHandler == nil {
		return
	}
	m.errHandler = errHandler
}

func (m *method) SetHandlers(handlers []lcontext.Middleware) {
	if m.lock {
		return
	}
	if len(handlers) == 0 {
		return
	}
	m.handlers = handlers
}

func (m *method) Name() string {
	return m.name
}

func (m *method) Service() string {
	return m.service
}

func (m *method) Exec(goCtx context.Context, req interface{}, ip string) (resp interface{}, code int64, info string, err error) {
	atomic.AddUint64(&m.invokeTimes, 1)
	ctx := m.pool.Get().(*lcontext.Context)
	lcontext.Reset(ctx, goCtx, req, ip)
	ctx.Next()
	resp = ctx.Response()
	code = ctx.Code()
	info = ctx.Info()
	err = ctx.BusinessErr()
	m.pool.Put(ctx)
	if code != 0 || err != nil {
		atomic.AddUint64(&m.errTimes, 1)
	}
	return
}

func New(
	name, service string,
	business lcontext.Business,
) Method {
	m := &method{
		name:       name,
		service:    service,
		business:   business,
		handlers:   nil,
		pool:       nil,
		errHandler: nil,
	}
	m.pool = &sync.Pool{New: func() interface{} {
		return lcontext.New(
			m.business,
			m.handlers,
			m.service,
			m.name,
			m.errHandler,
		)
	},
	}
	return m
}
