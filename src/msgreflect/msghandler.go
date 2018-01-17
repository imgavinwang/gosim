package msgreflect

import (
	"fmt"
	"bytes"
	"encoding/binary"

	"trademsg"
	"github.com/golang/protobuf/proto"
)

type PacketHead struct {
    Version uint32
    Length uint32
    Command uint32
    CheckSum uint32
}

/*******************************************************
    --修改处开始--hardcode
*******************************************************/
//1.新增消息名称映射及命令码
var MsgMap = map[string]uint32 {
        "Req": 0X01,
}

//2.注册新增proto消息映射，包括命令码和对应protobuf消息
func RegistMsg() {
	RegisterMessage(0X01, &trademsg.Req{})
}

/*******************************************************
    --修改处结束--
*******************************************************/

func PackageMessage(msgid uint32, msgJson string) (head_data []byte, req_data []byte, err error) {
    req, err := ParseMessage(msgid, msgJson)
    if err != nil {
        fmt.Println("ParseMessage error:", err.Error())
        return
    }    
    req_data, err = proto.Marshal(req)
    if err != nil {
        fmt.Println("proto Marshal error:", err.Error())
        return
    }

    head := PacketHead {
    	Version: 1,
    	Length: uint32(32 + len(req_data)),
    	Command: msgid,
    	VenderID: 1,
    	Market: 10,
    	IsCheckSum: 1,
    	CheckSum: 0,
    	Extend: 1,
    }
    head_buf := new(bytes.Buffer)
	binary.Write(head_buf, binary.BigEndian, head)
	head_data = head_buf.Bytes()
    
    fmt.Printf("msg_head:\n%+v\n\n", head)
    fmt.Printf("msg_body:\n{ %+v}\n\n", req)

	return
}