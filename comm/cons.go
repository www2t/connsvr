//go:generate stringer -type=CMD,GET_MSG_KIND,PROTO -output=cons_string.go

package comm

import "time"

type CMD byte

const (
	PING CMD = iota + 1 // 1, 现在的技术方案用不到心跳
	ENTER
	LEAVE
	PUB
	MSGS
	PUSH
	ERR = 0xff
)

type GET_MSG_KIND byte

const (
	NOTIFY   GET_MSG_KIND = iota + 1 // 推送通知，然后客户端主动拉后端服务
	DISPLAY                          // 推送整条消息，客户端不用拉
	CONNDATA                         // 推送通知，然后客户端来connsvr拉消息
)

type PROTO int

const (
	TCP PROTO = iota + 1 //1
	HTTP
	UDP
)

const (
	BUSI_REPORT = "report"
	BUSI_STAT   = "stat"
	BUSI_PUSH   = "push"
)

type Stat struct {
	Ip    string
	N     int
	Rid   string
	Msg   string
	Num   int
	Btime time.Time
	Etime time.Time
}

// Msgs is from logic svr
type Msgs []*struct {
	MsgId string
	Uid   string
	Sid   string
	Body  string
}

// ServExt will be transfered to client
type ServExt struct {
	GetMsgKind GET_MSG_KIND
}

// PushExt is from backend
type PushExt struct {
	MsgId      string
	GetMsgKind GET_MSG_KIND
}

// CliExt is from client
type CliExt struct {
	Cookie string
}
