package response

import (
	"net/http"
	"wmw-user-api/utility/wmwerrors"

	"github.com/gin-gonic/gin"
)

func ErrorExit(c *gin.Context, err wmwerrors.WMWError) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": err.Code(), "message": err.Error()})
}

func DataExit(c *gin.Context, data any) {
	err := wmwerrors.Nil()
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": err.Code(), "message": err.Error(), "data": data})
}
