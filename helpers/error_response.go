package helpers

import (
	"github.com/gofiber/fiber/v2"
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

func ErrorResponse(c *fiber.Ctx, err error) error {
	if errorx.IsOfType(err, domain.ErrorRoot) {
		errorType := errorx.GetTypeName(err)

		httpStatusCode, ok := errorx.ExtractProperty(err, domain.ErrorHttpStatusProperty)

		if ok {
			return c.Status(httpStatusCode.(int)).JSON(fiber.Map{
				"error":   errorType,
				"message": err.Error(),
			})
		}

		ctx := c

		configStatus := false
		for errorType, statusCode := range ErrorStatusMap {
			if errorx.IsOfType(err, errorType) {
				ctx = ctx.Status(statusCode)
				configStatus = true
				break
			}
		}

		if !configStatus {
			for errorType, statusCode := range ErrorRootStatusMap {
				if errorx.IsOfType(err, errorType) {
					ctx = ctx.Status(statusCode)
					break
				}
			}
		}

		return ctx.JSON(fiber.Map{
			"error":   errorType,
			"message": err.Error(),
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"error":   "internal_server_error",
		"message": err.Error(),
	})
}
