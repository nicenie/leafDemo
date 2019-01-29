package internal

import (
	"github.com/name5566/leaf/gate"
	llog "github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	_ = agent
}

func rpcCloseAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	_ = agent
}
