package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/go-util/errors"
	"net/http"
)

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    -1,
		"message": err.Error(),
		"data":    nil,
	})
}

func ErrorAuth(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    -1,
		"message": errors.ErrorAuth.Error(),
		"data":    nil,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    nil,
	})
}

func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    data,
	})
}
