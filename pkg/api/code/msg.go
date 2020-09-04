package code

var Message = map[int]string{
	Success: "OK",
}

func GetMessage(code int) string {
	if message, ok := Message[code]; ok {
		return message
	}

	return "未知错误"
}
