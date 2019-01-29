package main

import (
	"bytes"
	"client/globby"
	"client/innermsg"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"os/signal"
)

func IntToByte(num int32) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(err)
	return buffer.Bytes()
}

func shortToByte(num uint16) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(err)
	return buffer.Bytes()
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("err occur: ", err)
		os.Exit(1)
	}
}

func sendMsg(cmd uint32, uid int32, gameId int32, body []byte, conn net.Conn) {
	head := &innermsg.NetHead{
		Uid:    uid,
		Cmd:    cmd,
		GameId: uint32(gameId),
	}

	msg := &innermsg.ClientBody{
		Head: head,
		Body: body,
	}

	pData, err := proto.Marshal(msg)
	if err != nil {
		log.Println("proto error")
		return
	}
	size := len(pData) + 2 //(id len) + (msgdata len)
	data := make([]byte, 0)

	blen := IntToByte(int32(size))
	data = append(data, blen...)
	idbuf := shortToByte(uint16(0))
	data = append(data, idbuf...)
	data = append(data, pData...)
	n, err := conn.Write(data)
	if n > 0 {
		log.Println("send ok size:", n)
		return
	} else {
		log.Println("send data err")
	}
}

func send2Access(conn net.Conn) {
	body := &globby.CmdLoginRequest{
		Account:   "send2Access00000000000000000000000",
		Passwd:    "send2Access1111111111111",
		Nickname:  "send2Access",
		HeadUrl:   "https://www.baidu.com/send2Access",
		HeadId:    1,
		Sex:       1,
		OsType:    5,
		LoginType: 6,
	}

	bodyData, err := proto.Marshal(body)
	if err != nil {
		log.Println("proto error")
		return
	}
	sendMsg(0x2001, 0, 0, bodyData, conn)
}

func send2Access1(conn net.Conn) {
	body := &globby.CmdLoginRequest{
		Account:   "send2Access11111111111111111111111111",
		Passwd:    "send2Access1wwwwwww",
		Nickname:  "send2Access1",
		HeadUrl:   "https://www.baidu.com/send2Access1",
		HeadId:    1,
		Sex:       1,
		OsType:    5,
		LoginType: 6,
	}

	bodyData, err := proto.Marshal(body)
	if err != nil {
		log.Println("proto error")
		return
	}
	sendMsg(0x2001, 0, 0, bodyData, conn)
}

func sendHearbeat(conn net.Conn) {
	var body []byte
	sendMsg(0x2000, 0, 0, body, conn)
}

func sendCreateCmd(uid int32, conn net.Conn) {
	body := &globby.CmdRequestCreateRoom{
		GameId: 1,
		Param:  "balabalalbalblablalbla",
		Rule:   "RuleRuleRuleRuleRuleRuleRuleRuleRuleRuleRule",
	}
	bodyData, err := proto.Marshal(body)
	if err != nil {
		log.Println("proto error")
		return
	}
	sendMsg(0x2010, uid, 1, bodyData, conn)
}

func procLoginResponse(data []byte, conn net.Conn) {
	res := &globby.CmdLoginResponse{}
	err := proto.Unmarshal(data, res)
	if err != nil {
		log.Println("unmarshal err")
	}
	log.Printf("uid=%d|nickname=%s|headurl=%s|diamond=%d|coin=%d|card=%d|sex=%d", res.Uid, res.Nickname, res.HeadUrl, res.Diamond, res.Coin, res.Card, res.Sex)
	sendCreateCmd(res.Uid, conn)
}

func procCJRoomResponse(data []byte, conn net.Conn) {
	res := &globby.CmdCrtJonRoomResponse{}
	err := proto.Unmarshal(data, res)
	if err != nil {
		log.Println("unmarshal err")
	}
	log.Printf("Create Join errCode:%d", res.ErrCode)
}

func main() {
	host := "127.0.0.1:3563"
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	check := true

	send2Access(conn)

	buf := make([]byte, 2048)
	for check {
		select {
		case <-signalChan:
			check = false
		default:
			size, err := conn.Read(buf)
			if err != nil {
				log.Println("read err")
				check = false
			}
			if size > 0 {
				log.Println("recv size : ", size)
				buffsize := bytes.NewBuffer(buf[:4])
				var isze int32
				binary.Read(buffsize, binary.BigEndian, &isze)
				log.Println("recv size : ", isze)

				msg := &innermsg.ClientBody{}
				data := make([]byte, 0)
				data = append(data, buf[4:size]...)
				err = proto.Unmarshal(data, msg)
				if err != nil {
					log.Println("Unmarshal err:", err.Error())
				}
				log.Printf("cmd=0x%04x\n", msg.Head.Cmd)
				log.Printf("uid=%d\n", msg.Head.Uid)
				switch msg.Head.Cmd {
				case 0x2002:
					procLoginResponse(msg.Body, conn)
				case 0x2000:
					log.Println("recv hearbeat msg")
					sendHearbeat(conn)
				case 0x2012:
					procCJRoomResponse(msg.Body, conn)
				default:
					log.Printf("unknown cmd=0x%04x", msg.Head.Cmd)
					log.Println()
				}
			}
			buf = make([]byte, 2048)
		}
	}

}
