package globalError

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ParameterIsEmpty   = 1001 //参数为空
	ParameterIsIllegal = 1002 //参数不合法
	ServerError        = 1003 //服务器错误
	ResultIsEmpty      = 1004 //士兵不存在
)

type GlobalHandler func(c *gin.Context) (interface{}, error)

func ErrorHandler(handler GlobalHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := handler(c)
		if err != nil {
			globalError := err.(GlobalError)
			c.JSON(globalError.Status, globalError)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

//ParameterExpression 参数异常
func ParameterExpression(message string, code int) GlobalError {
	if code == ParameterIsEmpty {
		return GlobalError{
			Status:  http.StatusBadRequest,
			Code:    ParameterIsEmpty,
			Message: message,
		}
	} else {
		return GlobalError{
			Status:  http.StatusBadRequest,
			Code:    ParameterIsIllegal,
			Message: message,
		}
	}
}

//ServerExpression 服务器异常
func ServerExpression(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusInternalServerError,
		Code:    ServerError,
		Message: message,
	}
}

func ResultExpression(message string) GlobalError {
	return GlobalError{
		Status:  http.StatusNotFound,
		Code:    ResultIsEmpty,
		Message: message,
	}
}
