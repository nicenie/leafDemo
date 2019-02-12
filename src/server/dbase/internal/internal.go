package internal

import (
	"server/base"

	"github.com/name5566/leaf/module"
)

var (
	// ChanRPC db chanrpc
	ChanRPC  = skeleton.ChanRPCServer
	skeleton = base.NewSkeleton()
)

// Module leaf module
type Module struct {
	*module.Skeleton
}

// OnInit leaf module require init
func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

// OnDestroy leaf module require destory
func (m *Module) OnDestroy() {

}
