package internal

import (
	"server/normal"

	"github.com/name5566/leaf/gate"
	llog "github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	llog.Debug("new connect.")
	_ = agent
}

func rpcCloseAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	udata := agent.UserData()
	if udata == nil {
		return
	}
	meta := udata.(*normal.AgentMeta)
	if _, ok := amgr.Get(meta.UID); ok {
		llog.Debug("key=%d was removed in amgr.", meta.UID)
		amgr.Remove(meta.UID)
	}
}
