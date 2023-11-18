package middleware

import (
	"github.com/fxh111111/utility/jwt"
	"github.com/fxh111111/utility/response"
	"github.com/fxh111111/utility/wmwerrors"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("wmw-token")
		if err != nil {
			response.ErrorExit(c, wmwerrors.Internal(err))
			return
		}
		var uid string
		uid, err = jwt.ParseToken(token)
		if err != nil {
			response.ErrorExit(c, wmwerrors.NoAuth(err))
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}
