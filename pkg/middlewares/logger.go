package middlewares

import (
	"bytes"
	
	"io"
	"time"

	"github.com/Cynthia/goapi/pkg"
	"github.com/Cynthia/goapi/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responseBodyWriter struct{
	gin.ResponseWriter
	body *bytes.Buffer
}

func(r responseBodyWriter)Write(b []byte)(int,error){
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger()gin.HandlerFunc{
	return func(c *gin.Context){
		w:=&responseBodyWriter{body: &bytes.Buffer{},ResponseWriter: c.Writer}
		c.Writer=w
		var requestBody []byte
		if c.Request.Body!=nil{
			requestBody,_=io.ReadAll(c.Request.Body)
			c.Request.Body=io.NopCloser(bytes.NewBuffer(requestBody))
		}

		start:=time.Now()
		c.Next()

		cost:=time.Since(start)
		responStatus:=c.Writer.Status()

		logFields:=[]zap.Field{
			zap.Int("status",responStatus),
			zap.String("request",c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query",c.Request.URL.RawQuery),
			zap.String("ip",c.ClientIP()),
			zap.String("user-agent",c.Request.UserAgent()),
			zap.String("errors",c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time",pkg.MicrosecondsStr(cost)),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
            // 请求的内容
            logFields = append(logFields, zap.String("Request Body", string(requestBody)))

            // 响应的内容
            logFields = append(logFields, zap.String("Response Body", w.body.String()))
        }

        if responStatus > 400 && responStatus <= 499 {
            // 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
            logger.Warn("HTTP Warning "+cast.ToString(responStatus), logFields...)
        } else if responStatus >= 500 && responStatus <= 599 {
            // 除了内部错误，记录 error
            logger.Error("HTTP Error "+cast.ToString(responStatus), logFields...)
        } else {
            logger.Debug("HTTP Access Log", logFields...)
        }
    
	}
}

