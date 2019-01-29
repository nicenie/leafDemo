package msg

import (
	"github.com/name5566/leaf/network/protobuf"
	"server/msg/innermsg"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&innermsg.ClientBody{})
}
