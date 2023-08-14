package result

/**
统一定义返回格式
{
	"code":200,//http状态码
	"errorCode":""//错误状态码
	"msg":"",//错误信息
	"request":"",//请求路径
	"data":{}//数据
}
*/

type Response struct {
	Code      int         `json:"code"`
	ErrorCode string      `json:"errorCode"`
	Msg       string      `json:"msg"`
	Request   string      `json:"request"`
	Data      interface{} `json:"data"`
}

// 定义构造函数
func response(code int, errorCode string, msg string, request string) *Response {
	return &Response{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
		Request:   request,
		Data:      nil,
	}
}

// BusinessResponse 允许自定义相应类型，则无Data数据
func BusinessResponse(code int, errorCode string, msg string) *Response {
	return response(code, errorCode, msg, "")
}

func (response *Response) BuildMsg(msg string, request string) Response {
	return Response{
		Code:      response.Code,
		ErrorCode: response.ErrorCode,
		Msg:       msg,
		Request:   request,
		Data:      response.Data,
	}
}

func (response *Response) BuildData(data interface{}, request string) Response {
	return Response{
		Code:      response.Code,
		ErrorCode: response.ErrorCode,
		Msg:       response.Msg,
		Request:   request,
		Data:      data,
	}
}
