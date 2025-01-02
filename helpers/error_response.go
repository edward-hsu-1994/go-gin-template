package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
	"go-fiber-template/domain"
)

var ErrorRootStatusMap = map[*errorx.Type]int{
	domain.ErrorSystemRoot:        500,
	domain.ErrorServerSideRoot:    500,
	domain.ErrorClientSideRoot:    400,
	domain.ErrorBusinessLogicRoot: 400,
	domain.ErrorDependencyRoot:    500,
}

var ErrorStatusMap = map[*errorx.Type]int{
	domain.ErrorPostNotFound:      404,
	domain.ErrorInvalidPostStatus: 400,
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ErrorResponse(c, r.(error))
				c.Abort()
			}
		}()
		c.Next()
	}
}

func ErrorResponse(c *gin.Context, err error) error {
	if errorx.IsOfType(err, domain.ErrorRoot) {
		errorType := errorx.GetTypeName(err)

		httpStatusCode, ok := errorx.ExtractProperty(err, domain.ErrorHttpStatusProperty)

		if ok {
			c.JSON(httpStatusCode.(int), gin.H{
				"error":   errorType,
				"message": err.Error(),
			})
			return nil
		}

		ctx := c
		responseStatusCode := 500

		configStatus := false
		for errorType, statusCode := range ErrorStatusMap {
			if errorx.IsOfType(err, errorType) {
				responseStatusCode = statusCode
				configStatus = true
				break
			}
		}

		if !configStatus {
			for errorType, statusCode := range ErrorRootStatusMap {
				if errorx.IsOfType(err, errorType) {
					responseStatusCode = statusCode
					break
				}
			}
		}

		ctx.JSON(responseStatusCode, gin.H{
			"error":   errorType,
			"message": err.Error(),
		})
	}

	c.JSON(500, gin.H{
		"error":   "internal_server_error",
		"message": err.Error(),
	})

	return nil
}
