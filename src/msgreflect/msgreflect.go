package msgreflect

import (
	"fmt"
	"errors"
	"reflect"

    "utils"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto" 
)

//type MessageHandler func(msgid uint16, msg interface{}) (nmsg proto.Message, command uint32)

type MessageInfo struct {
    msgType    reflect.Type
    //msgHandler MessageHandler
}

var (
    msg_map = make(map[uint32]MessageInfo)
)

func RegisterMessage(msgid uint32, msg interface{}) {
    var info MessageInfo
    info.msgType = reflect.TypeOf(msg.(proto.Message))
    //info.msgHandler = handler
    msg_map[msgid] = info
}

func ParseMessage(msgid uint32, msgJson string) (nmsg proto.Message, err error) {
    if info, ok := msg_map[msgid]; ok {
        nmsg = reflect.New(info.msgType.Elem()).Interface().(proto.Message)
        //msg := reflect.New(info.msgType.Elem()).Interface()
        //nmsg = info.msgHandler(msgid, msg)
        err = jsonpb.UnmarshalString(msgJson, nmsg)
        if err != nil {
            fmt.Printf("Unmarshal msgJson to pb request error:%s\n", err)
            return
        }

        //SID for GUID
        val := reflect.ValueOf(nmsg).Elem()
        val.FieldByName("SID").Elem().SetString(utils.GetGuid())
        return
    } else {
        err = errors.New("not found msgid")
        return 
    }    
}
