package internal

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	llog "github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
	"server/msg/globby"
	"server/msg/innermsg"
)

func init() {
	handler(&innermsg.ClientBody{}, handleClientMSG)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleClientMSG(args []interface{}) {
	data := args[0].(*innermsg.ClientBody)
	agent := args[1].(gate.Agent)
	procCmd(agent, data)
}

func procCmd(agent gate.Agent, data *innermsg.ClientBody) {
	Cmd := data.Head.Cmd
	llog.Debug("Client Cmd=0x%04x", Cmd)
	switch Cmd {
	case msg.CMD_CLIENT_LOGIN_REQUEST:
		procLoginRequest(agent, data.Head, data.Body)
	}
}

/*处理客户端登录消息*/
func procLoginRequest(agent gate.Agent, head *innermsg.NetHead, data []byte) {
	req := &globby.CmdLoginRequest{}
	err := proto.Unmarshal(data, req)
	if err != nil {
		llog.Error("Unmarshal err:%s", err.Error())
		return
	}
	llog.Debug("login req: Account=%s|Nickname=%s|Sex=%d|HeadUrl=%s", req.Account, req.Nickname, req.Sex, req.HeadUrl)

}
