package server

import (
	"github.com/moshangguang/leocat/lcontext"
	"github.com/moshangguang/leocat/server/codec"
	"github.com/moshangguang/leocat/service"
)

var (
	// DefaultAddress 默认服务地址
	DefaultAddress = ":0"
	// DefaultName 默认服务名
	DefaultName = "catServer"
)

type Options struct {
	Name       string
	Address    string
	ErrHandler lcontext.ErrHandler
	Chains     []lcontext.Middleware
	Services   []service.Service
	Cc         codec.Codec
}

type Option func(options *Options)

func DefaultOptions() Options {
	return Options{
		Name:       DefaultName,
		Address:    DefaultAddress,
		ErrHandler: nil,
		Chains:     nil,
		Services:   nil,
	}
}
