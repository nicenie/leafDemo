package msg

import (
	"dataserver/msg/innermsg"

	"github.com/name5566/leaf/network/protobuf"
)

// Processor 消息解析器
var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&innermsg.ClientBody{})
}
