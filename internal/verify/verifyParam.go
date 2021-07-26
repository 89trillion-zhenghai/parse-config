package verify

import (
	"parse-config/internal/globalError"
	"strconv"
)

//ParamIsEmpty 校验参数是否为空
func ParamIsEmpty(params ...string) error {
	for _, v := range params {
		if len(v) == 0 {
			return globalError.ParameterExpression("参数不能为空", globalError.ParameterIsEmpty)
		}
	}
	return nil
}

//ParamIsDigit 校验参数是否为数字
func ParamIsDigit(params ...string) error {
	for _, v := range params {
		if !isDigit(v) {
			return globalError.ParameterExpression("参数不能含有除数字外的其他字符", globalError.ParameterIsIllegal)
		}
	}
	return nil
}

func isDigit(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}
