package internal

import (
	"dataserver/base"
	"dataserver/normal"

	"github.com/name5566/leaf/module"
)

var (
	// ChanRPC 模块间通信RPC
	ChanRPC = skeleton.ChanRPCServer
	//
	skeleton = base.NewSkeleton()
	// agent管理器
	amgr = normal.NewAgentMap()
)

// Module leaf-module
type Module struct {
	*module.Skeleton
}

// OnInit leaf-module 要求实现接口
func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

// OnDestroy leaf-module 要求实现接口
func (m *Module) OnDestroy() {

}
