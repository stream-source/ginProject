package result

import (
	"ginProject/exception"
)

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
//var (
//	OK   = BusinessResponse(http.StatusOK, SUCCESS_CODE, "ok")
//	FAIL = BusinessResponse(http.StatusInternalServerError, FAIL_CODE, "fail")
//)

// 使用自定义错误异常处理
var (
	OK           = exception.DefineBaseError(SUCCESS_CODE, "", "ok")
	FAIL         = exception.DefineBaseError(FAIL_CODE, "", "fail")
	SEX_IS_NULL  = exception.DefineBaseError("20101", "", "sex不能为空")
	UNAUTHORIZED = exception.DefineBaseError("401", "", "未授权")
	FORBIDDEN    = exception.DefineBaseError("403", "", "无权限禁止访问")
)

// 定义全局常量
const (
	SUCCESS_CODE = "1" //请求成功，错误码
	FAIL_CODE    = "0" //请求失败
)
