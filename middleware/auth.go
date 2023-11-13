package middleware

import (
	"wmw-user-api/utility"
	"wmw-user-api/utility/response"
	"wmw-user-api/utility/wmwerrors"

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
		uid, err = utility.ParseToken(token)
		if err != nil {
			response.ErrorExit(c, wmwerrors.NoAuth(err))
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}
