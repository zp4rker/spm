package sock_api

import "strings"

func ParseInput(input string) Request {
	inputSplit := strings.Split(input, " ")
	if input[0] == '/' {
		cmd := strings.ToLower(inputSplit[0])
		switch cmd {
		case "/register":
			return RegisterRequest
		default:
			return UnknownRequest
		}
	} else {
		if inputSplit[0] == "HEARTBEAT" {
			return HeartbeatInfo
		} else if inputSplit[0] == "TERMINATED" {
			return TerminatedInfo
		} else {
			return UnknownRequest
		}
	}
}