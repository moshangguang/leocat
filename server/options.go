package server

import "github.com/moshangguang/leocat/lcontext"

type Options struct {
	ErrHandler lcontext.ErrHandler
	Handlers   []lcontext.Middleware
}
