package result

import "net/http"

// 错误码规则:
// (1) 错误码需为 > 0 的数;
//
// (2) 错误码为 5 位数:
//
//	----------------------------------------------------------
//	    第1位               2、3位                  4、5位
//	----------------------------------------------------------
//	  服务级错误码          模块级错误码	         具体错误码
//	----------------------------------------------------------
var (
	OK   = BusinessResponse(http.StatusOK, SUCCESS_CODE, "ok")
	FAIL = BusinessResponse(http.StatusInternalServerError, FAIL_CODE, "fail")
)

// 定义全局常量
const (
	SUCCESS_CODE = "1" //请求成功，错误码
	FAIL_CODE    = "0" //请求失败
)
