package gate

import (
	"server/access"
	"server/msg"
	"server/msg/innermsg"
)

func init() {
	msg.Processor.SetRouter(&innermsg.ClientBody{}, access.ChanRPC)
}
