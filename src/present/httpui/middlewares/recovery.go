package middlewares

import (
	"base/src/common"
	"base/src/common/log"
	"base/src/present/httpui/resources"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"runtime/debug"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				log.Error(c, "[Recovery from panic]\ntime: [%v]\nerror: [%v] \nrequest: [%v]\nstack: [%v]\n",
					time.Now(), err, string(httpRequest), string(debug.Stack()))
				e := common.ErrSystemError(c, fmt.Sprintf("recovery, err:[%s]", err))
				c.JSON(e.GetHttpStatus(), resources.ConvertErrorToResponse(e))
				c.Abort()
			}
		}()
		c.Next()
	}
}
