package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestZipHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"ZipHandler test 1", args{mockCtx(map[string]interface{}{"body": []string{"../sampleDirectory/user/log.txt", "../sampleDirectory/user/messages.txt", "./sampleDirectory/someFile.txt"}})}, http.StatusOK},
		{"ZipHandler test 2", args{mockCtx(map[string]interface{}{"body": []string{}})}, http.StatusBadRequest},
		{"ZipHandler test 3", args{mockCtx(map[string]interface{}{"body": `nil`})}, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ZipHandler(tt.args.c)

		})

		if tt.args.c.Writer.Status() != tt.want {
			t.Errorf("ZipFiles() %v", tt.args.c.Writer.Status())
		}
	}
}

func mockCtx(content map[string]interface{}) *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{
		Header: make(http.Header),
	}

	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	return c
}
