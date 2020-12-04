package Middlewares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"foodorder/Service"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func SignRequest(context *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: context.Writer}
		context.Writer = w

		hash := Service.MD5(w.body.String())
		context.Header("Signature", hash)

		context.Next()
		fmt.Println("Response body: \"" + w.body.String()+"\"")

	}

