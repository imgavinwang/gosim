package main

import (
    "fmt"
    "net"
    "flag"
    "io/ioutil"
    
    "msgreflect"
)

var help = flag.Bool("h", false, "help")
var msg = flag.String("msg", "Req", "proto message name")
var file = flag.String("file", "./msgjson/Req.json", "json file")
var ipport = flag.String("ipport", "172.18.18.117:5322", "ip:port")

func LoadInputFile(filename string) (content string, err error) {
    fileContentBytes, err := ioutil.ReadFile(filename)
    if err != nil {
        return
    }
    content = string(fileContentBytes[:])
    return
}

func main() {
    flag.Parse()

    if *help {
        flag.Usage()
        return
    }

    msgJson, err := LoadInputFile(*file)
    if err != nil {
        fmt.Printf("Load json file error: %s\n", err)
        return
    }

    //dynamic parse Json to protobuf message.
    msgreflect.RegistMsg()
    head_data, req_data, _ := msgreflect.PackageMessage(uint32(msgreflect.MsgMap[*msg]), msgJson)
    
    //send proto message to socket.
    conn, err := net.Dial("tcp", *ipport)
    if err != nil {
    	fmt.Println("dial error:", err.Error())
    	return
    }
    _, err = conn.Write(head_data)
    if err != nil {
    	fmt.Println("write head error:", err.Error())
    	return
    }
    _, err = conn.Write(req_data)
    if err != nil {
    	fmt.Println("write req error:", err.Error())
    	return
    }

    fmt.Println("send msg ok.", *ipport)
    conn.Close()
}
