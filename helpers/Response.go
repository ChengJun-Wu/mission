package helpers

const CodeOk = 0
const CodeFail = 1
const CodeNeedLogin = 2
const CodeNeedAuth = 3

const MessageOk = "ok"
const MessageFail = "fail"
const MessageNeedLogin = "need login"
const MessageNeedAuth = "need auth"

func ResponseSuccess(args ...interface{}) map[string]interface{} {

	var (
		code interface{}
		message interface{}
		data interface{}
	)

	for idx, arg := range args{
		if idx == 0 {
			data = arg
		}
		if idx == 1 {
			message = arg
		}
		if idx == 2 {
			code = arg
		}
	}

	if code == nil {
		code = CodeOk
	}

	if message == nil {
		message = MessageOk
	}

	resp := map[string]interface{}{
		"code": code,
		"message": message,
		"data": data,
	}
	return resp
}

func ResponseFail(args ...interface{}) map[string]interface{} {

	var (
		code interface{}
		message interface{}
		data interface{}
	)

	for idx, arg := range args{
		if idx == 0 {
			message = arg
		}
		if idx == 1 {
			code = arg
		}
		if idx == 2 {
			data = arg
		}
	}

	if code == nil {
		code = CodeFail
	}

	if message == nil {
		message = MessageFail
	}

	resp := map[string]interface{}{
		"code": code,
		"message": message,
		"data": data,
	}
	return resp
}

func ResponseNeedLogin() map[string]interface{} {
	return ResponseFail(MessageNeedLogin, CodeNeedLogin)
}

func ResponseNeedAuth() map[string]interface{} {
	return ResponseFail(MessageNeedAuth, CodeNeedAuth)
}