
## 模拟器功能

根据protobuf消息，动态的使用json格式内容生成protobuf报文，通过socket发送到指定IP和端口。
protobuf消息配置后，只需要修改json里的字段值，启动程序即可动态生成报文内容。


## 使用步骤说明

1.proto文件配置:  
  1）trademsg/trade.proto里配置需要使用的proto消息。
  2）protoc --g_out=. *.proto生成.go文件。

2.动态创建消息映射：  
  msgreflect/msghander.go里的RegistMsg()函数配置proto消息映射，包括命令码和对应protobuf消息。

3.根据protobuf字段配置json格式消息体：  
  msgjson目录下新增json格式文件，可参考CustomerOpenReq.json例子。

4.编译  
  在gosim目录下，export GOPAHT=`pwd`后，go build即可。

3.启动命令  

```
gosim -h  
Usage of gosim:  
  -file string  
        json file (default "./msgjson/Req.json")  
  -h    help
  -ipport string
        ip:port (default "172.18.18.117:5322")
  -msg string
        proto message name (default "Req")

eg:
gosim -msg Req -file msgjson/Req.json -ipport 172.18.18.117:5321

cat  msgjson/Req.json
{
  "SID": "sid",
  "UserID": 2

```