package sock_api

type Request uint8

const (
	UnknownRequest Request = iota
	HeartbeatInfo
	TerminatedInfo
	RegisterRequest
)
