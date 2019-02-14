package gate

import (
	"dataserver/access"
	"dataserver/msg"
	"dataserver/msg/innermsg"
)

func init() {
	msg.Processor.SetRouter(&innermsg.ClientBody{}, access.ChanRPC)
}
