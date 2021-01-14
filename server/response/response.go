package response

type Response struct {
	Code   string      `json:"code"`
	Info   string      `json:"info"`
	Result interface{} `json:"result,omitempty"`
}

type ResponseCode struct {
	Code string
	Info string
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

const (
	SUCCESS_CODE                = "RC00000"
	GENERAL_ERROR_CODE          = "RC60000"
	PARAM_ERROR_CODE            = "RC80000"
	SYS_ERROR_CODE              = "RC90000"
	SAVE_ERROR_CODE             = "RC90001"
	DELETE_ERROR_CODE           = "RC90002"
	UNSUPPORTED_GRANT_TYPE_CODE = "RC100001"
	INVALID_CLIENT_CODE         = "RC100002"
	INVALID_GRANT_CODE          = "RC100003"
)

var SUCCESS = &ResponseCode{Code: SUCCESS_CODE, Info: "请求成功"}
var GENERAL_ERROR = &ResponseCode{Code: GENERAL_ERROR_CODE, Info: "出错提示"}
var PARAM_ERROR = &ResponseCode{Code: PARAM_ERROR_CODE, Info: "参数错误"}
var SYS_ERROR = &ResponseCode{Code: SYS_ERROR_CODE, Info: "系统错误"}
var SAVE_ERROR = &ResponseCode{Code: SAVE_ERROR_CODE, Info: "插入数据错误"}
var DELETE_ERROR = &ResponseCode{Code: DELETE_ERROR_CODE, Info: "删除数据错误"}
var UNSUPPORTED_GRANT_TYPE = &ResponseCode{Code: UNSUPPORTED_GRANT_TYPE_CODE, Info: "不支持的grant_type。"}
var INVALID_CLIENT = &ResponseCode{Code: INVALID_CLIENT_CODE, Info: "请求的appid或secret参数无效。"}
var INVALID_GRANT = &ResponseCode{Code: INVALID_GRANT_CODE, Info: "请求的Authorization Code、Access Token、Refresh Token等信息是无效的。"}

func GenSuccess(result interface{}) (resp *Response) {
	resp = Gen(SUCCESS, "", result)
	return
}

func GenError(code *ResponseCode, info string) (resp *Response) {
	resp = Gen(code, info, nil)
	return
}

func Gen(code *ResponseCode, info string, result interface{}) (resp *Response) {
	resp = &Response{Code: code.Code}
	if len(info) > 0 {
		resp.Info = info
	} else {
		resp.Info = code.Info
	}
	resp.Result = result
	return
}
